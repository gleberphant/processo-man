package servicos

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/aplicacao/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ServicoUsuario struct {
	RepoUsuarios *repositorios.RepositorioUsuario
}

func NovoCDUUsuario(usuariosRepo *repositorios.RepositorioUsuario) *ServicoUsuario {

	return &ServicoUsuario{
		RepoUsuarios: usuariosRepo,
	}
}

func (a *ServicoUsuario) Fechar() error {
	a.RepoUsuarios.Fechar()

	return nil
}

func (u *ServicoUsuario) CriaUsuario(usuario *entidades.Usuario) error {

	usuario.UUID, _ = uuid.NewV7()

	senhaForte, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	usuario.Senha = string(senhaForte)

	return u.RepoUsuarios.Criar(*usuario)
}

func (u *ServicoUsuario) CriarCliente(cliente *entidades.Cliente) error {
	// 1. Cria a entidade base usando a regra já existente (UUID é gerado aqui)
	err := u.CriaUsuario(&cliente.Usuario)
	if err != nil {
		return err
	}

	// 2. Vincula os dados na tabela específica
	return u.RepoUsuarios.AdicionarPerfilCliente(*cliente)
}

func (u *ServicoUsuario) CriarColaborador(colaborador *entidades.Colaborador) error {
	err := u.CriaUsuario(&colaborador.Usuario)
	if err != nil {
		return err
	}

	return u.RepoUsuarios.AdicionarPerfilColaborador(*colaborador)
}

func (u *ServicoUsuario) ListarUsuarios() ([]entidades.Usuario, error) {
	lista, err := u.RepoUsuarios.ListarUsuarios()
	if err != nil {
		return nil, err
	}
	return lista, nil
}

func (u *ServicoUsuario) ListarClientes() ([]entidades.Cliente, error) {
	return u.RepoUsuarios.ListarClientes()
}

func (u *ServicoUsuario) ListarColaboradores() ([]entidades.Colaborador, error) {
	return u.RepoUsuarios.ListarColaboradores()
}

func (u *ServicoUsuario) EditarUsuario(usuario entidades.Usuario) error {

	err := u.RepoUsuarios.Editar(usuario) // edição não altera senha

	if err != nil {
		return err
	}

	if usuario.Senha != "" {
		senhaForte, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		u.RepoUsuarios.MudarSenha(string(senhaForte), usuario.UUID)
	}

	return nil
}

func (u *ServicoUsuario) DeletarUsuario(usuarioUUID uuid.UUID) error {

	err1 := u.RepoUsuarios.Deletar(usuarioUUID)

	err2 := u.RepoUsuarios.DeletarCliente(usuarioUUID)

	err3 := u.RepoUsuarios.DeletarColaborador(usuarioUUID)

	if err := errors.Join(err1, err2, err3); err != nil {

		return err
	}
	return nil
}

func (u *ServicoUsuario) BuscarUsuarioPorUUID(usuarioUUID uuid.UUID) (*entidades.Usuario, error) {
	if usuarioUUID == uuid.Nil {
		return nil, errors.New("UUID nulo")
	}

	return u.RepoUsuarios.BuscarPorUUID(usuarioUUID)
}
