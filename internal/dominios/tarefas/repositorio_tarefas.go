package tarefas

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type RepositorioTarefa struct {
	conn *sql.DB
}

// NovoRepositorioProcesso cria uma nova instância do repositório de processos e estabelece a conexão.
func NovoRepositorioTarefa(conn *sql.DB) *RepositorioTarefa {
	repo := RepositorioTarefa{
		conn: conn,
	}

	return &repo
}

// Criar insere um novo registro de tarefa na tabela de tarefas.
func (r *RepositorioTarefa) CriarTarefa(tarefa Tarefa) error {

	db := r.conn

	_, err := db.Exec("INSERT INTO tarefas (uuid, processo_uuid, nome) VALUES (?, ?, ?)",
		tarefa.UUID,
		tarefa.ProcessoUUID,
		tarefa.Nome,
	)

	return err

}

// Listar retorna todos os tarefas cadastrados no banco de dados.
func (r *RepositorioTarefa) ListarTarefas(processoUUID uuid.UUID) ([]Tarefa, error) {

	db := r.conn

	rows, err := db.Query("SELECT uuid, responsavel_uuid, nome FROM tarefas WHERE processo_uuid=?", processoUUID.String())

	if err != nil {
		return nil, fmt.Errorf("Error: SQLITE>ListarTarefa>SELECT: %w", err)
	}

	defer rows.Close()

	var lista []Tarefa

	for rows.Next() {

		tarefa := Tarefa{}

		rows.Scan(&tarefa.UUID, &tarefa.ResponsavelUUID, &tarefa.Nome)

		lista = append(lista, tarefa)
	}

	return lista, nil

}

// Atular uma tarefa do banco de dados
func (r *RepositorioTarefa) EditarTarefa(tarefa Tarefa) error {

	if tarefa.UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("UPDATE tarefas SET nome = ? WHERE uuid = ?", tarefa.Nome, tarefa.UUID)

	return err
}

// Deletar remove um tarefa do banco de dados utilizando seu UUID.
func (r *RepositorioTarefa) DeletarTarefa(UUID uuid.UUID) error {

	if UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("DELETE FROM tarefas WHERE uuid=?", UUID)

	return err
}

// BuscarPorUUID recupera os dados de um tarefa específico através do seu identificador único.
func (r *RepositorioTarefa) BuscarTarefaPorUUID(UUID uuid.UUID) (*Tarefa, error) {

	db := r.conn

	row := db.QueryRow("SELECT uuid, processo_uuid, nome FROM tarefas WHERE uuid=?", UUID.String())

	tarefa := &Tarefa{}
	err := row.Scan(&tarefa.UUID, &tarefa.ProcessoUUID, &tarefa.Nome)

	if err != nil {
		return nil, err
	}

	return tarefa, nil

}
