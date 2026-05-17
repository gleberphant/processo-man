package processos

import (
	"errors"
	"fmt"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

// Criar insere um novo registro de tarefa na tabela de tarefas.
func (r *RepositorioProcesso) CriarTarefa(tarefa entidades.Tarefa) error {

	db := r.conn

	_, err := db.Exec("INSERT INTO tarefas (uuid, processo_uuid, nome) VALUES (?, ?, ?)",
		tarefa.UUID,
		tarefa.ProcessoUUID,
		tarefa.Nome,
	)

	return err

}

// Listar retorna todos os tarefas cadastrados no banco de dados.
func (r *RepositorioProcesso) ListarTarefas() ([]entidades.Tarefa, error) {

	db := r.conn

	rows, err := db.Query("SELECT uuid, processo_uuid, nome FROM tarefas ")

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var lista []entidades.Tarefa

	for rows.Next() {

		tarefa := entidades.Tarefa{}

		rows.Scan(&tarefa.UUID, &tarefa.ProcessoUUID, &tarefa.Nome)

		lista = append(lista, tarefa)
	}

	return lista, nil

}

// Deletar remove um tarefa do banco de dados utilizando seu UUID.
func (r *RepositorioProcesso) AtualizarTarefa(tarefa entidades.Tarefa) error {

	if tarefa.UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("UPDATE tarefas SET nome = ? WHERE uuid = ?", tarefa.Nome, tarefa.UUID)

	return err
}

// Deletar remove um tarefa do banco de dados utilizando seu UUID.
func (r *RepositorioProcesso) DeletarTarefa(UUID uuid.UUID) error {

	if UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("DELETE FROM tarefas WHERE uuid=?", UUID)

	return err
}

// BuscarPorUUID recupera os dados de um tarefa específico através do seu identificador único.
func (r *RepositorioProcesso) BuscarTarefaPorUUID(UUID uuid.UUID) (*entidades.Tarefa, error) {

	db := r.conn

	row := db.QueryRow("SELECT uuid, processo_uuid, nome FROM tarefas WHERE uuid=?", UUID.String())

	tarefa := &entidades.Tarefa{}
	err := row.Scan(&tarefa.UUID, &tarefa.ProcessoUUID, &tarefa.Nome)

	if err != nil {
		return nil, err
	}

	return tarefa, nil

}

// verifica se existe
func (r *RepositorioProcesso) AutenticarTarefa(UUID uuid.UUID) (string, error) {
	db := r.conn

	row := db.QueryRow("SELECT uuid FROM tarefas WHERE uuid=?;", UUID.String())

	var strUUID string

	err := row.Scan(&strUUID)

	if err != nil {
		return "", fmt.Errorf("Erro no SELECT de autenticacao de tarefa %w", err)
	}

	return strUUID, nil
}
