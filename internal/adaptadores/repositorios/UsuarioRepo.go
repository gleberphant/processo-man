package repositorios

import (
	"database/sql"
	"errors"
	"log"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/BancoDeDados"
	"github.com/google/uuid"
)

type UsuarioRepo struct {
	Conn *sql.DB
}

// NovoUsuarioRepo cria uma nova instância do repositório de usuários e estabelece a conexão.
func NovoUsuarioRepo() (*UsuarioRepo, error) {
	repo := UsuarioRepo{}
	if err := repo.Conectar(); err != nil {
		return nil, err
	}
	return &repo, nil

}

// Conectar inicializa a conexão com o banco de dados SQLite.
func (r *UsuarioRepo) Conectar() error {
	var err error
	r.Conn, err = BancoDeDados.ConectarSQLITE()

	if err != nil {
		log.Printf("Erro na conexao com banco de dados: %s", err)
		return err
	}
	return nil
}

// Fechar encerra a conexão ativa com o banco de dados.
func (r *UsuarioRepo) Fechar() {
	r.Conn.Close()
}

// Criar insere um novo registro de usuário na tabela de usuários.
func (r *UsuarioRepo) Criar(usuario *entidades.Usuario) error {

	db := r.Conn

	_, err := db.Exec("INSERT INTO usuarios (uuid, nome, email, senha) VALUES (?, ?, ?, ?)",
		usuario.UUID,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
	)

	if err != nil {
		log.Printf("Erro na criacao do usuario: %s", err)
		return err
	}

	return nil

}

// verifica se existe
func (r *UsuarioRepo) VerificaExistencia(usuario *entidades.Usuario) error {
	db := r.Conn

	row := db.QueryRow("SELECT uuid FROM usuarios WHERE uuid=?", usuario.UUID)

	err := row.Scan()

	if err != nil {
		return err
	}

	return nil
}

// verifica se existe
func (r *UsuarioRepo) AutenticarUsuario(usuario *entidades.Usuario) error {
	db := r.Conn

	row := db.QueryRow("SELECT uuid FROM usuarios WHERE email=? AND senha=?", usuario.Email, usuario.Senha)

	err := row.Scan(&usuario.UUID)

	if err != nil {
		return err
	}

	return nil
}

// BuscarPorUUID recupera os dados de um usuário específico através do seu identificador único.
func (r *UsuarioRepo) BuscarPorUUID(usuario *entidades.Usuario) (*entidades.Usuario, error) {

	db := r.Conn

	rows, err := db.Query("SELECT uuid, nome, email FROM usuarios WHERE uuid=?; ", usuario.UUID)

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("Usuario não encontrado")
	}

	rows.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email)

	return usuario, nil

}

// Listar retorna todos os usuários cadastrados no banco de dados.
func (r *UsuarioRepo) Listar() (*[]entidades.Usuario, error) {

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

	return &listaUsuario, nil

}

// Deletar remove um usuário do banco de dados utilizando seu UUID.
func (r *UsuarioRepo) Atualizar(usuario entidades.Usuario) error {

	return nil
}

// Deletar remove um usuário do banco de dados utilizando seu UUID.
func (r *UsuarioRepo) Deletar(usuario entidades.Usuario) error {

	if usuario.UUID == uuid.Nil {
		return errors.New("UUID do usuario não informado")
	}

	db := r.Conn

	_, err := db.Exec("DELETE FROM usuarios WHERE uuid=?", usuario.UUID)

	if err != nil {
		return err
	}

	return nil
}
