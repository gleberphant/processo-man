package repositorios

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

type RepositorioUsuario struct {
	Conn *sql.DB
}

// NovoRepositorioUsuario cria uma nova instância do repositório de usuários e estabelece a conexão.
func NovoRepositorioUsuario(conn *sql.DB) *RepositorioUsuario {
	repo := RepositorioUsuario{
		Conn: conn,
	}

	return &repo
}

// Fechar encerra a conexão ativa com o banco de dados.
func (r *RepositorioUsuario) Fechar() {
	r.Conn.Close()
}

// Criar insere um novo registro de usuário na tabela de usuários.
func (r *RepositorioUsuario) Criar(usuario entidades.Usuario) error {

	db := r.Conn

	_, err := db.Exec("INSERT INTO usuarios (uuid, nome, email, senha) VALUES (?, ?, ?, ?)",
		usuario.UUID,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
	)

	if err != nil {

		return err
	}

	return nil

}

// Listar retorna todos os usuários cadastrados no banco de dados.
func (r *RepositorioUsuario) Listar() ([]entidades.Usuario, error) {

	db := r.Conn

	rows, err := db.Query("SELECT uuid, nome, email FROM usuarios ")

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listaUsuario []entidades.Usuario

	for rows.Next() {

		usuario := entidades.Usuario{}

		rows.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email)

		listaUsuario = append(listaUsuario, usuario)
	}

	return listaUsuario, nil

}

// Deletar remove um usuário do banco de dados utilizando seu UUID.
func (r *RepositorioUsuario) Atualizar(usuario entidades.Usuario) error {

	return nil
}

// Deletar remove um usuário do banco de dados utilizando seu UUID.
func (r *RepositorioUsuario) Deletar(UUID uuid.UUID) error {

	if UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.Conn

	_, err := db.Exec("DELETE FROM usuarios WHERE uuid=?", UUID)

	if err != nil {
		return err
	}

	return nil
}

// BuscarPorUUID recupera os dados de um usuário específico através do seu identificador único.
func (r *RepositorioUsuario) BuscarPorUUID(uuid uuid.UUID) (*entidades.Usuario, error) {

	db := r.Conn

	rows, err := db.Query("SELECT uuid, nome, email FROM usuarios WHERE uuid=?; ", uuid.String())

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("Usuario não encontrado")
	}

	usuario := &entidades.Usuario{}
	rows.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email)

	return usuario, nil

}

func (r *RepositorioUsuario) BuscarPorEmail(email string) (*entidades.Usuario, error) {

	db := r.Conn

	row := db.QueryRow("SELECT uuid, nome, email FROM usuarios WHERE email=?", email)

	usuario := &entidades.Usuario{}

	err := row.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email, &usuario.Senha)

	if err != nil {
		return nil, err
	}

	return usuario, nil

}

// verifica se existe
func (r *RepositorioUsuario) AutenticarUsuario(email string, senha string) (string, error) {
	db := r.Conn

	row := db.QueryRow("SELECT uuid FROM usuarios WHERE email=? AND senha=?", email, senha)

	var UUID string
	err := row.Scan(&UUID)

	if err != nil {
		return "", fmt.Errorf("Erro no SELECT de autenticacao de USUARIO %w", err)
	}

	return UUID, nil
}
