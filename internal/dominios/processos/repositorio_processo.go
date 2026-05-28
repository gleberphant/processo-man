package processos

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

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

func (r *RepositorioProcesso) ListarProcessosPorCliente(cliente_uuid uuid.UUID) ([]entidades.Processo, error) {

	db := r.conn

	rows, err := db.Query("SELECT uuid, nome, cliente_uuid, colaborador_uuid  FROM processos WHERE cliente_uuid=?", cliente_uuid.String())

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listaProcesso []entidades.Processo

	for rows.Next() {

		Processo := entidades.Processo{}

		rows.Scan(&Processo.UUID, &Processo.Nome, &Processo.Dono.UUID, &Processo.Responsavel.UUID)

		listaProcesso = append(listaProcesso, Processo)
	}

	return listaProcesso, nil

}

// Deletar remove um processo do banco de dados utilizando seu UUID.
func (r *RepositorioProcesso) Editar(processo entidades.Processo) error {

	if processo.UUID == uuid.Nil {
		return errors.New("UUID NULO")
	}

	db := r.conn

	_, err := db.Exec("UPDATE processos SET nome = ? WHERE uuid = ?", processo.Nome, processo.UUID)

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

	row := db.QueryRow("SELECT nome, data_criacao, comentarios FROM processos WHERE uuid=?; ", UUID.String())

	var nome, dataCriacao string
	var comentarios sql.NullString

	err := row.Scan(&nome, &dataCriacao, &comentarios)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler dados do banco: %w", err)
	}

	processo := &entidades.Processo{}

	processo.UUID = UUID
	processo.Nome = nome
	if comentarios.Valid {
		processo.Comentarios = comentarios.String
	}

	// Utiliza o layout constante RFC3339 para interpretar a data e hora retornada com T e Z
	processo.DataCriacao, err = time.Parse(time.RFC3339, dataCriacao)

	if err != nil {
		return nil, fmt.Errorf("Erro ao fazer parse da data: %w", err)
	}

	return processo, nil

}

// verifica se existe
func (r *RepositorioProcesso) ValidarProcesso(UUID uuid.UUID) error {
	db := r.conn

	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM processos WHERE uuid = ?)", UUID.String())

	var valido bool

	err := row.Scan(&valido)

	if err != nil {
		return fmt.Errorf("Erro>Autenticar Processo> SELECT > %w", err)
	}

	if !valido {
		return errors.New("processo não encontrado")
	}

	return nil
}
