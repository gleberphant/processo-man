package autenticacao

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IRepositorioToken interface {
	Fechar()
	Criar(Token) (*Token, error)
	BuscarPorUUID(UUID uuid.UUID) (*Token, error)
	DeletarPorUsuarioUUID(uuid.UUID) error
}

type IRepositorioUsuario interface {
	Fechar()
	BuscarPorEmail(string) (*usuarios.Usuario, error)
	BuscarPorUUID(uuid.UUID) (*usuarios.Usuario, error)
}

type CDUAutenticacao struct {
	RepoTokens     IRepositorioToken
	RepoUsuarios   IRepositorioUsuario
	MapaPermissoes map[string][]string
}

func NovoCDUAutenticacao(tokensRepo IRepositorioToken, usuariosRepo IRepositorioUsuario) *CDUAutenticacao {

	permissoes := map[string][]string{
		"/":                        {"usuario", "colaborador"},
		"/usuarios/clientes/":      {"usuario", "colaborador"},
		"/usuarios/colaboradores/": {"usuario", "colaborador"},
		"/usuarios/":               {"usuario", "colaborador"},
		"/processos/":              {"usuario", "colaborador", "cliente"},
		"/tarefas/":                {"usuario", "colaborador"},
	}

	return &CDUAutenticacao{
		RepoTokens:     tokensRepo,
		RepoUsuarios:   usuariosRepo,
		MapaPermissoes: permissoes,
	}
}

// verificar se token é valido. retorna error se token não encontrado
func (a *CDUAutenticacao) ValidarToken(tokenUUID uuid.UUID) error {

	_, err := a.RepoTokens.BuscarPorUUID(tokenUUID)

	if err != nil {
		return fmt.Errorf("token invalido: %w ", err)
	}

	return nil
}

// verificar se token é valido. retorna error se token não encontrado
func (a *CDUAutenticacao) VerificarPermissao(tokenUUID uuid.UUID, rota string) error {

	var perfisPermitidos []string
	var matchEncontrado bool
	tamanhoMatch := 0

	// Busca o prefixo de rota mais longo que corresponde à URL requisitada
	for chave, perfis := range a.MapaPermissoes {
		if strings.HasPrefix(rota, chave) {
			if len(chave) > tamanhoMatch {
				tamanhoMatch = len(chave)
				perfisPermitidos = perfis
				matchEncontrado = true
			}
		}
	}

	if !matchEncontrado {
		return errors.New("rota não mapeada no sistema de permissões")
	}

	token, err := a.RepoTokens.BuscarPorUUID(tokenUUID)

	if err != nil {
		return fmt.Errorf("token invalido: %w ", err)
	}

	for _, perfil := range token.Perfis {

		for _, perfilPermitido := range perfisPermitidos {

			if strings.EqualFold(perfil, perfilPermitido) {
				return nil
			}

		}

	}

	return fmt.Errorf("perfil [%v] não autorizado a rota [%s]", token.Perfis, rota)

}

func (a *CDUAutenticacao) AutenticarUsuario(email string, senha string) (string, error) {

	//valida usuario
	// 1 Verifica se o usuario/email existe
	if email == "" {
		return "", errors.New("Usuario nulo")
	}

	usuario, err := a.RepoUsuarios.BuscarPorEmail(email)
	if err != nil {
		return "", fmt.Errorf("usuario não encontrado: %w ", err)
	}

	// 2 verifica senha
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))

	if err != nil {
		if senha != usuario.Senha {
			return "", fmt.Errorf("senha inválida : %w ", err)
		}
	}

	// Gerar token
	token, err := a.GerarToken(usuario)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token: %w ", err)
	}

	//retornar o uuid do token gerado
	return token.UUID.String(), nil
}

// gerar token
func (a *CDUAutenticacao) GerarToken(usuario *usuarios.Usuario) (*Token, error) {

	// limpar tokens antigos no repositorio
	err := a.RepoTokens.DeletarPorUsuarioUUID(usuario.UUID)

	// gerar novo uuid do token
	tokenUUID, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	// inserir o novo token no repositorio
	token, err := a.RepoTokens.Criar(Token{
		UUID:        tokenUUID,
		UsuarioUUID: usuario.UUID,
		Validade:    "temporario",
		Perfis:      usuario.Perfis,
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}
