package repositorios

import (
	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/google/uuid"
)

// repositorio para conexao com o google drive
type RepositorioArquivo struct {
	//receber a conexao com o google drive
	conn interface{}
}

// NovoRepositorioUsuario cria uma nova instância do repositório de usuários e estabelece a conexão.
func NovoRepositorioArquivo(conn interface{}) *RepositorioArquivo {
	repo := RepositorioArquivo{
		conn: conn,
	}

	return &repo
}

func (r *RepositorioArquivo) Fechar() {
	//conn.Close()
}

func (r *RepositorioArquivo) Listar() ([]entidades.Arquivo, error) {

	listaArquivo := []entidades.Arquivo{
		{UUID: uuid.New(), Nome: "arquivo1.txt", URL: "file://123"},
		{UUID: uuid.New(), Nome: "arquivo2.txt", URL: "file://123"},
		{UUID: uuid.New(), Nome: "arquivo3.txt", URL: "file://123"},
		{UUID: uuid.New(), Nome: "arquivo4.txt", URL: "file://123"},
	}

	return listaArquivo, nil

}

func (r *RepositorioArquivo) ListarArquivosPorUsuario(usuarioUUID uuid.UUID) ([]entidades.Arquivo, error) {

	listaArquivo := []entidades.Arquivo{
		{UUID: uuid.New(), Nome: "arquivo1.txt", URL: "file://123", Usuario: &entidades.Usuario{UUID: usuarioUUID}},
		{UUID: uuid.New(), Nome: "arquivo2.txt", URL: "file://123", Usuario: &entidades.Usuario{UUID: usuarioUUID}},
		{UUID: uuid.New(), Nome: "arquivo3.txt", URL: "file://123", Usuario: &entidades.Usuario{UUID: usuarioUUID}},
		{UUID: uuid.New(), Nome: "arquivo4.txt", URL: "file://123", Usuario: &entidades.Usuario{UUID: usuarioUUID}},
	}

	return listaArquivo, nil

}
