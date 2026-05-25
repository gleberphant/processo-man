package autenticacao

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type RepositorioToken struct {
	Conn *sql.DB
}

// NovoRepositorioToken cria uma nova instância do repositório de tokens e estabelece a conexão.
func NovoRepositorioToken(conn *sql.DB) *RepositorioToken {
	repo := RepositorioToken{
		Conn: conn,
	}

	return &repo

}

// Fechar encerra a conexão ativa com o banco de dados.
func (r *RepositorioToken) Fechar() {
	r.Conn.Close()
}

// Criar insere um novo registro de token na tabela de tokens.
func (r *RepositorioToken) Criar(token Token) (*Token, error) {

	db := r.Conn

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

	return &token, nil

}

// Deletar remove todos os tokens associados a um UUID de usuário específico.
func (r *RepositorioToken) DeletarPorUsuarioUUID(UsuarioUUID uuid.UUID) error {

	strUUID := UsuarioUUID.String()

	if strUUID == "" {
		return errors.New("UUID do usuario vazio")
	}
	db := r.Conn

	_, err := db.Exec("DELETE FROM tokens WHERE usuario_uuid=?", strUUID)

	if err != nil {
		return fmt.Errorf("[Erro no DELETE POR USUARIO UUID DO  token ]: %w", err)
	}

	return nil
}

// Ver busca os detalhes de um token específico através do seu UUID.
func (r *RepositorioToken) BuscarPorUUID(UUID uuid.UUID) (*Token, error) {

	db := r.Conn
	row := db.QueryRow("SELECT uuid, usuario_uuid, validade, perfis, data_criacao FROM tokens WHERE uuid=?; ", UUID.String())

	var token Token = Token{}
	var perfis string
	err := row.Scan(&token.UUID, &token.UsuarioUUID, &token.Validade, &perfis, &token.DataCriacao)

	token.Perfis = strings.Split(perfis, ",")

	if err != nil {
		return nil, err
	}

	return &token, nil

}
