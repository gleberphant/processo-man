package autenticacao

import (
	"database/sql"
	"errors"
	"fmt"

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

	_, err := db.Exec("INSERT INTO tokens(uuid, usuario_uuid, validade) VALUES(?,?, ?)", token.UUID.String(), token.UsuarioUUID.String(), token.Validade)

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

// Listar recupera todos os tokens registrados no banco de dados
func (r *RepositorioToken) Listar() ([]Token, error) {
	db := r.Conn

	rows, err := db.Query("SELECT uuid, usuario_uuid, validade, data_criacao FROM tokens")
	if err != nil {
		return nil, fmt.Errorf("[Erro no SELECT de LISTAR  token ]: %w", err)
	}
	defer rows.Close()

	var tokens []Token
	for rows.Next() {
		var t Token
		err := rows.Scan(&t.UUID, &t.UsuarioUUID, &t.Validade, &t.DataCriacao)
		if err != nil {
			return nil, fmt.Errorf("[Erro no SCAN ROWS DE LISTAR  token ]: %w", err)
		}
		tokens = append(tokens, t)
	}

	return tokens, nil
}

// Atualizar modifica os dados de validade ou comentários de um token existente
func (r *RepositorioToken) Editar(token Token) error {
	db := r.Conn

	_, err := db.Exec("UPDATE tokens SET validade = ?, comentarios = ? WHERE uuid = ?",
		token.Validade,
		token.Comentarios,
		token.UUID,
	)

	if err != nil {
		return fmt.Errorf("[Erro no UPDATE DO  token ]: %w", err)

	}

	return nil
}

// Deletar remove todos os tokens associados a um UUID de usuário específico.
func (r *RepositorioToken) Deletar(UUID string) error {

	if UUID == "" {
		return errors.New("UUID vazio")
	}
	db := r.Conn

	_, err := db.Exec("DELETE FROM tokens WHERE uuid=?", UUID)

	if err != nil {
		return fmt.Errorf("[Erro no DELETE POR UUID DO  token ]: %w", err)
	}

	return nil
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
	row := db.QueryRow("SELECT uuid, usuario_uuid, data_criacao, validade FROM tokens WHERE uuid=?; ", UUID.String())

	var token Token

	err := row.Scan(&token.UUID, &token.UsuarioUUID, &token.DataCriacao, &token.Validade)

	if err != nil {
		return nil, err
	}

	return &token, nil

}
