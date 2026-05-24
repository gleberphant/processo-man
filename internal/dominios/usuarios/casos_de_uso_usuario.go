package usuarios

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IRepositorioUsuario interface {
	Criar(Usuario) error
	Listar() ([]Usuario, error)
	Editar(Usuario) error
	Deletar(uuid.UUID) error
	BuscarPorUUID(uuid.UUID) (*Usuario, error)
	AdicionarPerfilCliente(Cliente) error
	AdicionarPerfilColaborador(Colaborador) error
	ListarClientes() ([]Cliente, error)
	ListarColaboradores() ([]Colaborador, error)
}

type CDUUsuario struct {
	RepoUsuarios IRepositorioUsuario
}

func NovoCDUUsuario(usuariosRepo IRepositorioUsuario) *CDUUsuario {

	return &CDUUsuario{
		RepoUsuarios: usuariosRepo,
	}
}

func (u *CDUUsuario) CriaUsuario(usuario *Usuario) error {

	usuario.UUID = uuid.New()

	senhaForte, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	usuario.Senha = string(senhaForte)

	return u.RepoUsuarios.Criar(*usuario)
}

func (u *CDUUsuario) CriarCliente(cliente *Cliente) error {
	// 1. Cria a entidade base usando a regra já existente (UUID é gerado aqui)
	err := u.CriaUsuario(&cliente.Usuario)
	if err != nil {
		return err
	}

	// 2. Vincula os dados na tabela específica
	return u.RepoUsuarios.AdicionarPerfilCliente(*cliente)
}

func (u *CDUUsuario) CriarColaborador(colaborador *Colaborador) error {
	err := u.CriaUsuario(&colaborador.Usuario)
	if err != nil {
		return err
	}

	return u.RepoUsuarios.AdicionarPerfilColaborador(*colaborador)
}

func (u *CDUUsuario) ListarUsuarios() ([]Usuario, error) {
	lista, err := u.RepoUsuarios.Listar()
	if err != nil {
		return nil, err
	}
	return lista, nil
}

func (u *CDUUsuario) ListarClientes() ([]Cliente, error) {
	return u.RepoUsuarios.ListarClientes()
}

func (u *CDUUsuario) ListarColaboradores() ([]Colaborador, error) {
	return u.RepoUsuarios.ListarColaboradores()
}

func (u *CDUUsuario) EditarUsuario(usuario Usuario) error {

	senhaForte, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	usuario.Senha = string(senhaForte)

	u.RepoUsuarios.Editar(usuario)

	return err
}

func (u *CDUUsuario) DeletarUsuario(usuarioUUID uuid.UUID) error {
	return u.RepoUsuarios.Deletar(usuarioUUID)
}

func (u *CDUUsuario) BuscarUsuarioPorUUID(usuarioUUID uuid.UUID) (*Usuario, error) {
	if usuarioUUID == uuid.Nil {
		return nil, errors.New("UUID nulo")
	}

	return u.RepoUsuarios.BuscarPorUUID(usuarioUUID)
}
