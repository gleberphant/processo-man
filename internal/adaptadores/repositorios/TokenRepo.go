package repositorios

import (
	"database/sql"
	"errors"
	"log"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/BancoDeDados"
	"github.com/google/uuid"
)

type TokenRepo struct {
	Conn *sql.DB
}

// NovoTokenRepo cria uma nova instância do repositório de tokens e estabelece a conexão.
func NovoTokenRepo() (*TokenRepo, error) {
	repo := TokenRepo{}
	if err := repo.Conectar(); err != nil {
		return nil, err
	}
	return &repo, nil

}

// Conectar inicializa a conexão com o banco de dados SQLite.
func (r *TokenRepo) Conectar() error {
	var err error
	r.Conn, err = BancoDeDados.ConectarSQLITE()

	if err != nil {
		log.Printf("Erro na conexao com banco de dados: %s", err)
		return err
	}
	return nil
}

// Fechar encerra a conexão ativa com o banco de dados.
func (r *TokenRepo) Fechar() {
	r.Conn.Close()
}

// Criar insere um novo registro de token na tabela de tokens.
func (r *TokenRepo) Criar(token entidades.Token) error {

	db := r.Conn

	_, err := db.Exec("INSERT INTO tokens(uuid, usuario_uuid, validade) VALUES(?,?, ?)", token.UUID, token.UsuarioUUID, token.Validade)

	if err != nil {
		log.Printf("Erro na criacao do token: %s", err)
		return err
	}

	return nil

}

// Ver busca os detalhes de um token específico através do seu UUID.
func (r *TokenRepo) BuscarPorUUID(token entidades.Token) (*entidades.Token, error) {

	db := r.Conn

	rows, err := db.Query("SELECT usuario_uuid, data_criacao FROM tokens WHERE uuid=?; ", token.UUID)

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("Token não encontrado")
	}

	rows.Scan(&token.UsuarioUUID, &token.DataCriacao)
	return &token, nil

}

// Listar recupera todos os tokens registrados no banco de dados
func (r *TokenRepo) Listar() ([]entidades.Token, error) {
	db := r.Conn

	rows, err := db.Query("SELECT uuid, usuario_uuid, validade, data_criacao FROM tokens")
	if err != nil {
		log.Printf("Erro ao listar tokens: %s", err)
		return nil, err
	}
	defer rows.Close()

	var tokens []entidades.Token
	for rows.Next() {
		var t entidades.Token
		err := rows.Scan(&t.UUID, &t.UsuarioUUID, &t.Validade, &t.DataCriacao)
		if err != nil {
			log.Printf("Erro ao escanear token: %s", err)
			continue
		}
		tokens = append(tokens, t)
	}

	return tokens, nil
}

// Atualizar modifica os dados de validade ou comentários de um token existente
func (r *TokenRepo) Atualizar(token entidades.Token) error {
	db := r.Conn

	_, err := db.Exec("UPDATE tokens SET validade = ?, comentarios = ? WHERE uuid = ?",
		token.Validade,
		token.Comentarios,
		token.UUID,
	)

	if err != nil {
		log.Printf("Erro ao atualizar token: %s", err)
		return err
	}

	return nil
}

// Deletar remove todos os tokens associados a um UUID de usuário específico.
func (r *TokenRepo) Deletar(token entidades.Token) error {

	if token.UsuarioUUID == uuid.Nil {
		return errors.New("UUID do usuario não informado")
	}
	db := r.Conn

	_, err := db.Exec("DELETE FROM tokens WHERE usuario_uuid=?", token.UsuarioUUID)

	if err != nil {
		log.Printf("Erro ao deletar token: %s", err)
		return err
	}

	return nil
}
