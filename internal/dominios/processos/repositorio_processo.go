package processos

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

type RepositorioProcesso struct {
	conn *sql.DB
}

// NovoRepositorioProcesso cria uma nova instância do repositório de processos e estabelece a conexão.
func NovoRepositorioProcesso(conn *sql.DB) *RepositorioProcesso {
	repo := RepositorioProcesso{
		conn: conn,
	}

	return &repo
}

// Fechar encerra a conexão ativa com o banco de dados.
func (r *RepositorioProcesso) Fechar() {
	r.conn.Close()
}

// Criar insere um novo registro de processo na tabela de processos.
func (r *RepositorioProcesso) Criar(Processo entidades.Processo) error {

	db := r.conn

	_, err := db.Exec("INSERT INTO processos (uuid, nome) VALUES (?, ?)",
		Processo.UUID,
		Processo.Nome,
	)

	if err != nil {

		return err
	}

	return nil

}

// Listar retorna todos os processos cadastrados no banco de dados.
func (r *RepositorioProcesso) Listar() ([]entidades.Processo, error) {

	db := r.conn

	rows, err := db.Query("SELECT uuid, nome FROM processos ")

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listaProcesso []entidades.Processo

	for rows.Next() {

		Processo := entidades.Processo{}

		rows.Scan(&Processo.UUID, &Processo.Nome)

		listaProcesso = append(listaProcesso, Processo)
	}

	return listaProcesso, nil

}

// Deletar remove um processo do banco de dados utilizando seu UUID.
func (r *RepositorioProcesso) Atualizar(Processo entidades.Processo) error {

	if Processo.UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("UPDATE processos SET nome = ? WHERE uuid = ?", Processo.Nome, Processo.UUID)

	if err != nil {
		return err
	}

	return nil
}

// Deletar remove um processo do banco de dados utilizando seu UUID.
func (r *RepositorioProcesso) Deletar(UUID uuid.UUID) error {

	if UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("DELETE FROM Processos WHERE uuid=?", UUID)

	if err != nil {
		return err
	}

	return nil
}

// BuscarPorUUID recupera os dados de um processo específico através do seu identificador único.
func (r *RepositorioProcesso) BuscarPorUUID(UUID uuid.UUID) (*entidades.Processo, error) {

	db := r.conn

	row := db.QueryRow("SELECT uuid, nome FROM processos WHERE uuid=?; ", UUID.String())

	Processo := &entidades.Processo{}
	err := row.Scan(&Processo.UUID, &Processo.Nome)

	if err != nil {
		return nil, err
	}

	return Processo, nil

}

// verifica se existe
func (r *RepositorioProcesso) AutenticarProcesso(UUID uuid.UUID) (string, error) {
	db := r.conn

	row := db.QueryRow("SELECT uuid, nome FROM processos WHERE uuid=?;", UUID.String())

	var str string = ""

	err := row.Scan(str)

	if err != nil {
		return "", fmt.Errorf("Erro no SELECT de autenticacao de Processo %w", err)
	}

	return UUID.String(), nil
}
