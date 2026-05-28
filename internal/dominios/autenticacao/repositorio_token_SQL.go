package autenticacao

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
	"github.com/google/uuid"
)

/*
* possui uma conexao persistente para carregar o mapa de
* permissoes em memoria e atualizar as permissoes e possui tambem uma
* conexao EM memoria para os tokens
 */

type RepositorioToken struct {
	connPermissoes *sql.DB                    // conexao persisente para carregar permissoes
	connToken      *sql.DB                    // conexao em memoria
	mapaPermissoes map[string]map[string]bool // mapa carregado em memoria
}

// NovoRepositorioToken cria uma nova instância do repositório de tokens e estabelece a conexão.
func NovoRepositorioToken(connPermissoes *sql.DB) *RepositorioToken {

	repo := RepositorioToken{
		connPermissoes: connPermissoes,
	}

	err := repo.carregarBancoDeTokensEmMemoria()
	if err != nil {
		panic(fmt.Errorf("Erro ao carregar Banco de Tokens em Memoria %w", err))
	}

	err = repo.carregarMapaPermissoesEmMemoria()
	if err != nil {
		panic(fmt.Errorf("Erro ao carregar mapa de Permissoes %w", err))
	}

	return &repo
}

// carrega o banco de dados em memoria. acoplado ao objeto para evitar falhas
func (r *RepositorioToken) carregarBancoDeTokensEmMemoria() error {

	log.Printf("Carregando repositorio token: banco de tokens em memoria")
	// carregar schema do banco de dados de tokens
	schemaDatabase, err := os.ReadFile("../database/schema_tokens.sql")

	if err != nil {
		return err
	}

	// carregar banco em memoria
	connEmMemoria, err := bancodedados.ConectarSQLITE()

	if err != nil {
		return err
	}

	// executar configuração do banco de  em memoria
	_, err = connEmMemoria.Exec(string(schemaDatabase))

	if err != nil {
		return err
	}

	r.connToken = connEmMemoria

	return nil
}

// carrega carrega o Mapa de Permissoes na memoria
func (r *RepositorioToken) carregarMapaPermissoesEmMemoria() error {

	log.Printf("Carregando repositorio token: mapa de permissoes")
	rows, err := r.connPermissoes.Query("SELECT rota, perfil, metodo FROM permissoes")
	if err != nil {
		return err
	}
	defer rows.Close()

	permissoes := make(map[string]map[string]bool)

	for rows.Next() {
		var rota, perfis, metodos string
		if err := rows.Scan(&rota, &perfis, &metodos); err != nil {
			return err
		}

		listaPerfil := strings.Split(strings.ToLower(perfis), ";")
		listaMetodo := strings.Split(strings.ToLower(metodos), ";")

		for _, perfil := range listaPerfil {

			for _, metodo := range listaMetodo {
				chave := metodo + ":" + rota
				if permissoes[chave] == nil {
					permissoes[chave] = make(map[string]bool)
				}
				permissoes[chave][perfil] = true
			}
		}
	}

	r.mapaPermissoes = permissoes

	for key, value := range r.mapaPermissoes {
		for key2, value2 := range value {
			log.Printf("\n rota [%s] perfil [%s] valor [%v]", key, key2, value2)
		}
	}
	return nil
}

// Fechar encerra a conexão ativa com o banco de dados.
func (r *RepositorioToken) Fechar() {

	r.connToken.Close()
	r.connPermissoes.Close()

}

func (r *RepositorioToken) VerificarPermissaoPerfil(perfil string, rota string) bool {

	return r.mapaPermissoes[strings.ToLower(rota)][strings.ToLower(perfil)]

}

// Criar insere um novo registro de token na tabela de tokens.
func (r *RepositorioToken) Criar(token *entidades.Token) (*entidades.Token, error) {

	db := r.connToken

	listaPerfis := strings.Join(token.Perfis, ",")

	_, err := db.Exec("INSERT INTO tokens(uuid, usuario_uuid, validade, perfis) VALUES(?, ?, ?, ?)",
		token.UUID.String(),
		token.UsuarioUUID.String(),
		token.Validade,
		listaPerfis)

	if err != nil {
		return nil, fmt.Errorf("[Erro no INSERT de criacao  do token]: %w", err)
	}

	// confere se realmente criou e devolve o token
	row := db.QueryRow("SELECT data_criacao FROM tokens WHERE uuid=?", token.UUID.String())

	err = row.Scan(&token.DataCriacao)
	if err != nil {
		return nil, fmt.Errorf("[Erro no SELECT de confirmacao da criacao do token ]: %w", err)
	}

	return token, nil

}

// Deletar remove todos os tokens associados a um UUID de usuário específico.
func (r *RepositorioToken) DeletarPorUsuarioUUID(UsuarioUUID uuid.UUID) error {

	strUUID := UsuarioUUID.String()

	if strUUID == "" {
		return errors.New("UUID do usuario vazio")
	}
	db := r.connToken

	_, err := db.Exec("DELETE FROM tokens WHERE usuario_uuid=?", strUUID)

	if err != nil {
		return fmt.Errorf("[Erro no DELETE POR USUARIO UUID DO  token ]: %w", err)
	}

	return nil
}

// Ver busca os detalhes de um token específico através do seu UUID.
func (r *RepositorioToken) BuscarPorUUID(UUID uuid.UUID) (*entidades.Token, error) {

	db := r.connToken
	row := db.QueryRow("SELECT uuid, usuario_uuid, validade, perfis, data_criacao FROM tokens WHERE uuid=?; ", UUID.String())

	var token entidades.Token = entidades.Token{}
	var perfis string
	err := row.Scan(&token.UUID, &token.UsuarioUUID, &token.Validade, &perfis, &token.DataCriacao)

	token.Perfis = strings.Split(perfis, ",")

	if err != nil {
		return nil, err
	}

	return &token, nil

}

// ListarTodos busca todos os tokens ativos no banco. Recomendado apenas para ambiente de debug.
func (r *RepositorioToken) ListarTodos() ([]*entidades.Token, error) {
	rows, err := r.connToken.Query("SELECT uuid, usuario_uuid, validade, perfis, data_criacao FROM tokens")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []*entidades.Token
	for rows.Next() {
		var t entidades.Token
		var perfis string
		if err := rows.Scan(&t.UUID, &t.UsuarioUUID, &t.Validade, &perfis, &t.DataCriacao); err != nil {
			return nil, err
		}
		t.Perfis = strings.Split(perfis, ",")
		tokens = append(tokens, &t)
	}

	return tokens, nil
}
