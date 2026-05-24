package usuarios

import (
	"database/sql"
	"errors"
	"fmt"

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
func (r *RepositorioUsuario) Criar(usuario Usuario) error {

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
func (r *RepositorioUsuario) AdicionarPerfilCliente(cliente Cliente) error {
	db := r.conn

	_, err := db.Exec("INSERT INTO clientes (usuario_uuid, cpf, endereco, tipo_pessoa) VALUES (?, ?, ?, ?)",
		cliente.UUID, cliente.CPF, cliente.Endereco, cliente.TipoPessoa)

	return err
}

// AdicionarPerfilColaborador vincula os dados específicos de colaborador a um usuário já existente.
// Exige que a struct Colaborador contenha o UUID válido do Usuario base.
func (r *RepositorioUsuario) AdicionarPerfilColaborador(colaborador Colaborador) error {
	db := r.conn

	_, err := db.Exec("INSERT INTO colaboradores (usuario_uuid, cargo) VALUES (?, ?)",
		colaborador.UUID, colaborador.Cargo)

	return err
}

// Listar retorna todos os usuários cadastrados no banco de dados.
func (r *RepositorioUsuario) Listar() ([]Usuario, error) {

	db := r.conn

	query := `
		SELECT 
			u.uuid, 
			u.nome, 
			u.email,
			CASE WHEN c.usuario_uuid IS NOT NULL THEN 'Cliente' ELSE '' END,
			CASE WHEN col.usuario_uuid IS NOT NULL THEN 'Colaborador' ELSE '' END
		FROM usuarios u
		LEFT JOIN clientes c ON u.uuid = c.usuario_uuid
		LEFT JOIN colaboradores col ON u.uuid = col.usuario_uuid
	`

	rows, err := db.Query(query)

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listaUsuario []Usuario

	for rows.Next() {

		usuario := Usuario{}
		var perfilCliente, perfilColab string

		err := rows.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email, &perfilCliente, &perfilColab)
		if err != nil {
			return nil, err
		}

		perfis := ""
		if perfilCliente != "" {
			perfis += perfilCliente
		}
		if perfilColab != "" {
			if perfis != "" {
				perfis += " / "
			}
			perfis += perfilColab
		}
		if perfis == "" {
			perfis = "Usuário"
		}
		usuario.Perfis = perfis

		listaUsuario = append(listaUsuario, usuario)
	}

	return listaUsuario, nil

}

// ListarClientes retorna todos os usuários que possuem o perfil de cliente.
func (r *RepositorioUsuario) ListarClientes() ([]Cliente, error) {
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

	var lista []Cliente
	for rows.Next() {
		var c Cliente
		if err := rows.Scan(&c.UUID, &c.Nome, &c.Email, &c.CPF); err != nil {
			return nil, err
		}
		lista = append(lista, c)
	}
	return lista, nil
}

// ListarColaboradores retorna todos os usuários que possuem o perfil de colaborador.
func (r *RepositorioUsuario) ListarColaboradores() ([]Colaborador, error) {
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

	var lista []Colaborador
	for rows.Next() {
		var c Colaborador
		if err := rows.Scan(&c.UUID, &c.Nome, &c.Email, &c.Cargo); err != nil {
			return nil, err
		}
		lista = append(lista, c)
	}
	return lista, nil
}

// Deletar remove um usuário do banco de dados utilizando seu UUID.
func (r *RepositorioUsuario) Editar(usuario Usuario) error {

	if usuario.UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("UPDATE usuarios SET nome = ?, email= ?, senha = ? WHERE uuid = ?", usuario.Nome, usuario.Email, usuario.Senha, usuario.UUID)

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

// BuscarPorUUID recupera os dados de um usuário específico através do seu identificador único.
func (r *RepositorioUsuario) BuscarPorUUID(UUID uuid.UUID) (*Usuario, error) {

	db := r.conn

	row := db.QueryRow("SELECT uuid, nome, email FROM usuarios WHERE uuid=? ", UUID.String())

	usuario := &Usuario{}
	row.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email)

	return usuario, nil

}

func (r *RepositorioUsuario) BuscarPorEmail(email string) (*Usuario, error) {

	db := r.conn

	row := db.QueryRow("SELECT uuid, nome, email, senha FROM usuarios WHERE email=?", email)

	usuario := &Usuario{}

	err := row.Scan(&usuario.UUID, &usuario.Nome, &usuario.Email, &usuario.Senha)

	if err != nil {
		return nil, err
	}

	return usuario, nil

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
