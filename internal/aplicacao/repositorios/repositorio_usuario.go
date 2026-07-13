package repositorios

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/google/uuid"
)

type RepositorioUsuario struct {
	conn *sql.DB
}

// NovoRepositorioUsuario cria uma nova instância do repositório de usuários e estabelece a conexão.
func NovoRepositorioUsuario(conn *sql.DB) *RepositorioUsuario {
	repo := RepositorioUsuario{
		conn: conn,
	}

	return &repo
}

// Fechar encerra a conexão ativa com o banco de dados.
func (r *RepositorioUsuario) Fechar() {
	r.conn.Close()
}

// Criar insere um novo registro de usuário na tabela de usuários.
func (r *RepositorioUsuario) Criar(usuario entidades.Usuario) error {

	db := r.conn

	_, err := db.Exec("INSERT INTO usuarios (uuid, nome, email, senha) VALUES (?, ?, ?, ?)",
		usuario.UUID,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
	)

	return err

}

// AdicionarPerfilCliente vincula os dados específicos de cliente a um usuário já existente.
// Exige que a struct Cliente contenha o UUID válido do Usuario base.
func (r *RepositorioUsuario) AdicionarPerfilCliente(cliente entidades.Cliente) error {
	db := r.conn

	_, err := db.Exec("INSERT INTO clientes (usuario_uuid, cpf, endereco, tipo_pessoa) VALUES (?, ?, ?, ?)",
		cliente.UUID, cliente.CPF, cliente.Endereco, cliente.TipoPessoa)

	return err
}

// AdicionarPerfilColaborador vincula os dados específicos de colaborador a um usuário já existente.
// Exige que a struct Colaborador contenha o UUID válido do Usuario base.
func (r *RepositorioUsuario) AdicionarPerfilColaborador(colaborador entidades.Colaborador) error {
	db := r.conn

	_, err := db.Exec("INSERT INTO colaboradores (usuario_uuid, cargo) VALUES (?, ?)",
		colaborador.UUID, colaborador.Cargo)

	return err
}

// Listar retorna todos os usuários cadastrados no banco de dados.
func (r *RepositorioUsuario) ListarUsuarios() ([]entidades.Usuario, error) {

	db := r.conn

	rows, err := db.Query(`
		SELECT u.uuid, u.nome, u.email,	cli.usuario_uuid, col.usuario_uuid
		FROM usuarios u
		LEFT JOIN clientes cli ON u.uuid = cli.usuario_uuid
		LEFT JOIN colaboradores col ON u.uuid = col.usuario_uuid
	`)

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listaUsuario []entidades.Usuario

	for rows.Next() {

		var uuidCliente, uuidColaborador sql.NullString

		usuario := entidades.Usuario{}

		err := rows.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email, &uuidCliente, &uuidColaborador)

		if err != nil {
			return nil, err
		}

		if uuidCliente.Valid {
			usuario.Perfis = append(usuario.Perfis, "Cliente")
		}

		if uuidColaborador.Valid {
			usuario.Perfis = append(usuario.Perfis, "Colaborador")
		}

		if len(usuario.Perfis) == 0 {
			usuario.Perfis = append(usuario.Perfis, "Usuario")
		}

		listaUsuario = append(listaUsuario, usuario)
	}

	return listaUsuario, nil

}

// ListarClientes retorna todos os usuários que possuem o perfil de cliente.
func (r *RepositorioUsuario) ListarClientes() ([]entidades.Cliente, error) {
	db := r.conn

	query := `
		SELECT u.uuid, u.nome, u.email, c.cpf 
		FROM usuarios u 
		INNER JOIN clientes c ON u.uuid = c.usuario_uuid
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []entidades.Cliente
	for rows.Next() {
		var c entidades.Cliente
		if err := rows.Scan(&c.UUID, &c.Nome, &c.Email, &c.CPF); err != nil {
			return nil, err
		}
		lista = append(lista, c)
	}
	return lista, nil
}

// ListarColaboradores retorna todos os usuários que possuem o perfil de colaborador.
func (r *RepositorioUsuario) ListarColaboradores() ([]entidades.Colaborador, error) {
	db := r.conn

	query := `
		SELECT u.uuid, u.nome, u.email, c.cargo 
		FROM usuarios u 
		INNER JOIN colaboradores c ON u.uuid = c.usuario_uuid
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []entidades.Colaborador
	for rows.Next() {
		var c entidades.Colaborador
		if err := rows.Scan(&c.UUID, &c.Nome, &c.Email, &c.Cargo); err != nil {
			return nil, err
		}
		lista = append(lista, c)
	}
	return lista, nil
}

// Deletar remove um usuário do banco de dados utilizando seu UUID.
func (r *RepositorioUsuario) Editar(usuario entidades.Usuario) error {

	if usuario.UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn
	var err error

	_, err = db.Exec("UPDATE usuarios SET nome = ?, email= ? WHERE uuid = ?", usuario.Nome, usuario.Email, usuario.UUID)

	if err != nil {
		return err
	}

	return nil
}

// altera senha
func (r *RepositorioUsuario) MudarSenha(novaSenha string, usuarioUUID uuid.UUID) error {

	if usuarioUUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("UPDATE usuarios SET senha = ? WHERE uuid = ?", novaSenha, usuarioUUID)

	if err != nil {
		return err
	}

	return nil
}

// Deletar remove um usuário do banco de dados utilizando seu UUID.
func (r *RepositorioUsuario) Deletar(UUID uuid.UUID) error {

	if UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("DELETE FROM usuarios WHERE uuid=?", UUID)

	if err != nil {
		return err
	}

	return nil
}

func (r *RepositorioUsuario) DeletarCliente(UUID uuid.UUID) error {

	if UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("DELETE FROM clientes WHERE usuario_uuid=?", UUID)

	if err != nil {
		return err
	}

	return nil
}

func (r *RepositorioUsuario) DeletarColaborador(UUID uuid.UUID) error {

	if UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("DELETE FROM colaboradores WHERE usuario_uuid=?", UUID)

	if err != nil {
		return err
	}

	return nil
}

// BuscarPorUUID recupera os dados de um usuário específico através do seu identificador único.
func (r *RepositorioUsuario) BuscarPorUUID(UUID uuid.UUID) (*entidades.Usuario, error) {

	db := r.conn

	row := db.QueryRow(`
		SELECT u.uuid, u.nome, u.email, cli.usuario_uuid, col.usuario_uuid
		FROM usuarios u
		LEFT JOIN clientes cli ON u.uuid = cli.usuario_uuid
		LEFT JOIN colaboradores col ON u.uuid = col.usuario_uuid
		WHERE u.uuid=?
	`, UUID)

	var uuidCliente, uuidColaborador sql.NullString

	usuario := entidades.Usuario{}

	err := row.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email, &uuidCliente, &uuidColaborador)

	if err != nil {
		return nil, err
	}
	// define a lista de perfis
	if uuidCliente.Valid {
		usuario.Perfis = append(usuario.Perfis, "Cliente")
	}

	if uuidColaborador.Valid {
		usuario.Perfis = append(usuario.Perfis, "Colaborador")
	}

	if len(usuario.Perfis) == 0 {
		usuario.Perfis = append(usuario.Perfis, "Usuario")
	}

	return &usuario, nil

}

func (r *RepositorioUsuario) BuscarPorEmail(email string) (*entidades.Usuario, error) {

	db := r.conn

	row := db.QueryRow(`
		SELECT u.uuid, u.nome, u.email,	u.senha, cli.usuario_uuid, col.usuario_uuid
		FROM usuarios u
		LEFT JOIN clientes cli ON u.uuid = cli.usuario_uuid
		LEFT JOIN colaboradores col ON u.uuid = col.usuario_uuid
		WHERE u.email=?
	`, email)

	var uuidCliente, uuidColaborador sql.NullString

	usuario := entidades.Usuario{}

	err := row.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email, &usuario.Senha, &uuidCliente, &uuidColaborador)

	if err != nil {
		return nil, err
	}
	// define a lista de perfis
	if uuidCliente.Valid {
		usuario.Perfis = append(usuario.Perfis, "cliente")
	}

	if uuidColaborador.Valid {
		usuario.Perfis = append(usuario.Perfis, "colaborador")
	}

	if len(usuario.Perfis) == 0 {
		usuario.Perfis = append(usuario.Perfis, "admin")
	}

	return &usuario, nil

}

// verifica se existe
func (r *RepositorioUsuario) AutenticarUsuario(email string, senha string) (string, error) {
	db := r.conn

	row := db.QueryRow("SELECT uuid FROM usuarios WHERE email=? AND senha=?", email, senha)

	var UUID string
	err := row.Scan(&UUID)

	if err != nil {
		return "", fmt.Errorf("Erro no SELECT de autenticacao de USUARIO %w", err)
	}

	return UUID, nil
}
