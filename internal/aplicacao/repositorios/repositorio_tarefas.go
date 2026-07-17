package repositorios

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
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

func (r *RepositorioTarefa) Fechar() {
	r.conn.Close()
}

// Criar insere um novo registro de tarefa na tabela de tarefas.
func (r *RepositorioTarefa) CriarTarefa(tarefa entidades.Tarefa) error {

	db := r.conn

	_, err := db.Exec("INSERT INTO tarefas (uuid, processo_uuid, responsavel_uuid, nome, comentarios) VALUES (?, ?, ?, ?, ?)",
		tarefa.UUID,
		tarefa.ProcessoUUID,
		tarefa.ResponsavelUUID,
		tarefa.Nome,
		tarefa.Comentarios,
	)

	return err

}

func (r *RepositorioTarefa) ListarTarefas() ([]entidades.Tarefa, error) {

	db := r.conn

	rows, err := db.Query("SELECT uuid, processo_uuid, responsavel_uuid, nome, comentarios FROM tarefas")

	if err != nil {
		return nil, fmt.Errorf("Error: SQLITE>ListarTodasTarefa>SELECT: %w", err)
	}

	defer rows.Close()

	var lista []entidades.Tarefa

	for rows.Next() {

		tarefa := entidades.Tarefa{}

		if err := rows.Scan(&tarefa.UUID, &tarefa.ProcessoUUID, &tarefa.ResponsavelUUID, &tarefa.Nome, &tarefa.Comentarios); err != nil {
			return nil, err
		}

		lista = append(lista, tarefa)
	}

	// Verifica se ocorreu algum erro durante a iteração das linhas
	return lista, rows.Err()
}

// Listar retorna todos os tarefas cadastrados no banco de dados.
func (r *RepositorioTarefa) ListarTarefasPorProcesso(processoUUID uuid.UUID) ([]entidades.Tarefa, error) {

	db := r.conn

	rows, err := db.Query("SELECT uuid, processo_uuid, responsavel_uuid, nome, comentarios FROM tarefas WHERE processo_uuid=?", processoUUID.String())

	if err != nil {
		return nil, fmt.Errorf("Error: SQLITE>ListarTarefa>SELECT: %w", err)
	}

	defer rows.Close()

	var lista []entidades.Tarefa

	for rows.Next() {

		tarefa := entidades.Tarefa{}

		if err := rows.Scan(&tarefa.UUID, &tarefa.ProcessoUUID, &tarefa.ResponsavelUUID, &tarefa.Nome, &tarefa.Comentarios); err != nil {
			return nil, err
		}

		lista = append(lista, tarefa)
	}

	// Verifica se ocorreu algum erro durante a iteração das linhas
	return lista, rows.Err()
}

// Listar retorna todos os tarefas cadastrados no banco de dados.
func (r *RepositorioTarefa) ListarTarefasPorResponsavel(responavelUUID uuid.UUID) ([]entidades.Tarefa, error) {

	db := r.conn

	rows, err := db.Query("SELECT uuid, processo_uuid, responsavel_uuid, nome, comentarios FROM tarefas WHERE responsavel_uuid=?", responavelUUID.String())

	if err != nil {
		return nil, fmt.Errorf("Error: SQLITE>ListarTarefaResponsavel>SELECT: %w", err)
	}

	defer rows.Close()

	var lista []entidades.Tarefa

	for rows.Next() {

		tarefa := entidades.Tarefa{}

		if err := rows.Scan(&tarefa.UUID, &tarefa.ProcessoUUID, &tarefa.ResponsavelUUID, &tarefa.Nome, &tarefa.Comentarios); err != nil {
			return nil, err
		}

		lista = append(lista, tarefa)
	}

	// Verifica se ocorreu algum erro durante a iteração das linhas
	return lista, rows.Err()
}

// Atular uma tarefa do banco de dados
func (r *RepositorioTarefa) EditarTarefa(tarefa entidades.Tarefa) error {

	if tarefa.UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("UPDATE tarefas SET nome = ?, responsavel_uuid = ?, comentarios = ? WHERE uuid = ?", tarefa.Nome, tarefa.ResponsavelUUID, tarefa.Comentarios, tarefa.UUID)

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

func (r *RepositorioTarefa) DeletarTarefasPorProcesso(ProcessoUUID uuid.UUID) error {

	if ProcessoUUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("DELETE FROM tarefas WHERE processo_uuid=?", ProcessoUUID)

	return err
}

// BuscarPorUUID recupera os dados de um tarefa específico através do seu identificador único.
func (r *RepositorioTarefa) BuscarTarefaPorUUID(UUID uuid.UUID) (*entidades.Tarefa, error) {

	db := r.conn

	row := db.QueryRow("SELECT uuid, processo_uuid,  responsavel_uuid,nome, comentarios, concluida, data_conclusao, data_criacao FROM tarefas WHERE uuid=?",
		UUID.String())

	tarefa := &entidades.Tarefa{}
	err := row.Scan(&tarefa.UUID,
		&tarefa.ProcessoUUID,
		&tarefa.ResponsavelUUID,
		&tarefa.Nome,
		&tarefa.Comentarios,
		&tarefa.Concluida,
		&tarefa.DataConclusao,
		&tarefa.DataCriacao)

	if err != nil {
		return nil, err
	}

	return tarefa, nil

}
