package autenticacao

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IRepositorioToken interface {
	Fechar()
	Criar(*entidades.Token) (*entidades.Token, error)
	BuscarPorUUID(UUID uuid.UUID) (*entidades.Token, error)
	DeletarPorUsuarioUUID(uuid.UUID) error
	VerificarPermissaoPerfil(string, string) bool
	//ObterMapaPermissoes(rota string) map[string]bool
}

type IRepositorioUsuario interface {
	Fechar()
	BuscarPorEmail(string) (*entidades.Usuario, error)
	BuscarPorUUID(uuid.UUID) (*entidades.Usuario, error)
}

type CDUAutenticacao struct {
	RepoTokens   IRepositorioToken
	RepoUsuarios IRepositorioUsuario
}

func NovoCDUAutenticacao(tokensRepo IRepositorioToken, usuariosRepo IRepositorioUsuario) *CDUAutenticacao {

	return &CDUAutenticacao{
		RepoTokens:   tokensRepo,
		RepoUsuarios: usuariosRepo,
	}
}

func (a *CDUAutenticacao) Fechar() error {
	a.RepoTokens.Fechar()
	a.RepoUsuarios.Fechar()
	return nil
}

// verificar se token tem permissão para acessar uma rota
func (a *CDUAutenticacao) VerificarPermissao(tokenUUID uuid.UUID, rota string, metodo string) error {

	// busca o token completo
	log.Printf("buscando token por uuid")
	token, err := a.RepoTokens.BuscarPorUUID(tokenUUID)
	log.Printf("fim da busca")
	if err != nil {
		return fmt.Errorf("Token não encontrado: %w ", err)
	}

	chaveRota := strings.ToLower(metodo + ":" + rota)
	// verificar permissao
	for _, perfil := range token.Perfis {
		if a.RepoTokens.VerificarPermissaoPerfil(chaveRota, strings.ToLower(perfil)) {
			return nil
		}
	}

	return fmt.Errorf("perfil [%v] não autorizado a rota [%s]", token.Perfis, chaveRota)

}

func (a *CDUAutenticacao) AutenticarUsuario(email string, senha string) (string, error) {

	//valida usuario
	if email == "" {
		return "", errors.New("Usuario nulo")
	}
	// 1 Verifica se o usuario/email existe
	usuario, err := a.RepoUsuarios.BuscarPorEmail(email)
	if err != nil {
		return "", fmt.Errorf("usuario não encontrado: %w ", err)
	}

	// 2 verifica senha
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))

	if err != nil {
		//fallback para usuario teste
		if senha != usuario.Senha && usuario.UUID.String() == "00000000-0000-0000-0000-000000000000" {
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
func (a *CDUAutenticacao) GerarToken(usuario *entidades.Usuario) (*entidades.Token, error) {

	// limpar tokens antigos no repositorio
	err := a.RepoTokens.DeletarPorUsuarioUUID(usuario.UUID)

	if err != nil {
		return nil, err
	}

	// gerar novo uuid do token
	tokenUUID, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	// inserir o novo token no repositorio
	token, err := a.RepoTokens.Criar(
		&entidades.Token{
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
