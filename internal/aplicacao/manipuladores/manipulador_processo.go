package manipuladores

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/gleberphant/ProcessoMan/internal/dominio/servicos"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// servindo como interface entre a camada de apresentação e os casos de uso.
type ManipuladorProcesso struct {
	servicoProcesso *servicos.CDUProcesso
	servicoUsuario  *servicos.ServicoUsuario
}

// NovoManipuladorProcesso cria e retorna uma nova instância de ManipuladorProcesso.
func NovoManipuladorProcesso(CasosDeUsoProcesso *servicos.CDUProcesso, CasosDeUsoUsuario *servicos.ServicoUsuario) *ManipuladorProcesso {
	return &ManipuladorProcesso{
		servicoProcesso: CasosDeUsoProcesso,
		servicoUsuario:  CasosDeUsoUsuario,
	}
}

func (m *ManipuladorProcesso) Fechar() {
	m.servicoProcesso.Fechar()
}

// PageCriar renderiza o formulário para criação de um novo processo.
func (m *ManipuladorProcesso) PageCriar(w http.ResponseWriter, r *http.Request) {

	listaClientes, err := m.servicoUsuario.ListarClientes()
	if err != nil {
		//tratar error
		log.Printf("Erro %v", err)
		return
	}

	listaColaboradores, err := m.servicoUsuario.ListarColaboradores()
	if err != nil {
		//tratar error
		log.Printf("Erro %v", err)
		return
	}

	viewModel := ViewModelProcesso{
		Clientes:      listaClientes,
		Colaboradores: listaColaboradores,
		Processos:     []entidades.Processo{},
	}

	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, r, viewModel)
}

// PageListar renderiza a página contendo a listagem de todos os processos.
func (m *ManipuladorProcesso) PageListar(w http.ResponseWriter, r *http.Request) {

	listaProcessos, err := m.servicoProcesso.ListarProcessos()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro Page Listar Processo:%v", err))
		return
	}

	viewModel := ViewModelProcesso{
		Processos: listaProcessos,
	}

	apresentacao.ExibirPaginaHTML("processo/page-listar-processos.html", w, r, viewModel)

}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageVerProcesso(w http.ResponseWriter, r *http.Request) {
	strUUID := r.PathValue("UUID")

	processoUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do processo inválido: %v", err))
		return
	}

	processo, err := m.servicoProcesso.BuscarProcessoPorUUID(processoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro Page Ver Processo:%v", err))
		return
	}

	viewModel := ViewModelProcesso{
		UUID:     strUUID,
		Processo: *processo,
		Anexos:   []string{"arquivo1.doc", "arquivo2.doc"},
	}

	apresentacao.ExibirPaginaHTML("processo/page-ver-processo.html", w, r, viewModel)
}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageEditar(w http.ResponseWriter, r *http.Request) {
	strUUID := r.PathValue("UUID")

	processoUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do processo inválido: %v", err))
		return
	}

	processo, err := m.servicoProcesso.BuscarProcessoPorUUID(processoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageEditar:%v", err))
		return
	}

	viewModel := ViewModelProcesso{
		UUID:     strUUID,
		Processo: *processo,
		//Anexos:   []string{"arquivo1.doc", "arquivo2.doc"},
	}

	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, r, viewModel)
}

// CriarProcessoPost processa a submissão do formulário para persistir um novo processo.
func (m *ManipuladorProcesso) CriarProcessoPost(w http.ResponseWriter, r *http.Request) {

	donoUUID, _ := uuid.Parse(r.PostFormValue("ColaboradorUUID"))

	responsavelUUID, _ := uuid.Parse(r.PostFormValue("ClienteUUID"))

	dono := entidades.Cliente{}
	dono.UUID = donoUUID

	responsavel := entidades.Colaborador{}
	responsavel.UUID = responsavelUUID

	Processo := entidades.Processo{
		Nome:        r.PostFormValue("nome"),
		Dono:        dono,
		Responsavel: responsavel,
	}

	err := m.servicoProcesso.CriarProcesso(Processo)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro criar Processo: %v", err))
		return
	}

	http.Redirect(w, r, "/processos", http.StatusSeeOther)

}

// EditarProcessoPost processa a atualização de um processo existente.
func (m *ManipuladorProcesso) EditarProcessoPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	UUID, err := uuid.Parse(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
		return
	}

	processoDTO := entidades.Processo{
		UUID: UUID,
		Nome: r.PostFormValue("nome"),
	}

	err = m.servicoProcesso.EditarProcesso(processoDTO)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
		return
	}

	http.Redirect(w, r, "/processos", http.StatusSeeOther)
}

// DeletarProcessoPost remove um processo com base no identificador enviado via formulário.
func (m *ManipuladorProcesso) DeletarProcessoPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	processoUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do processo inválido: %v", err))
		return
	}

	err = m.servicoProcesso.DeletarProcesso(processoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro deletar Processo: %v", err))
		return
	}

	http.Redirect(w, r, "/processos", http.StatusSeeOther)
}
