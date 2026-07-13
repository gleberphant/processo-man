package servicos

//temporariamente vou tirar as interfaces e acoplhar o repositorio ao service, porque dá muito trabalho ficar desenvolvendo com interfaces
/*
type IRepositorioUsuario interface {
	Fechar()
	Criar(entidades.Usuario) error
	ListarUsuarios() ([]entidades.Usuario, error)
	Editar(entidades.Usuario) error
	Deletar(uuid.UUID) error
	DeletarCliente(uuid.UUID) error
	DeletarColaborador(uuid.UUID) error
	BuscarPorUUID(uuid.UUID) (*entidades.Usuario, error)
	AdicionarPerfilCliente(entidades.Cliente) error
	AdicionarPerfilColaborador(entidades.Colaborador) error
	ListarClientes() ([]entidades.Cliente, error)
	ListarColaboradores() ([]entidades.Colaborador, error)
	MudarSenha(string, uuid.UUID) error
}

type IRepositorioUsuarioAutenticacao interface {
	Fechar()
	BuscarPorEmail(string) (*entidades.Usuario, error)
	BuscarPorUUID(uuid.UUID) (*entidades.Usuario, error)
}

type IRepositorioToken interface {
	Fechar()
	Criar(*entidades.Token) (*entidades.Token, error)
	BuscarPorUUID(UUID uuid.UUID) (*entidades.Token, error)
	DeletarPorUsuarioUUID(uuid.UUID) error
	VerificarPermissaoPerfil(string, string) bool
	//ObterMapaPermissoes(rota string) map[string]bool
}

type IRepositorioProcesso interface {
	Fechar()
	Criar(entidades.Processo) error
	Listar() ([]entidades.Processo, error)
	Editar(entidades.Processo) error
	Deletar(uuid.UUID) error
	BuscarPorUUID(uuid.UUID) (*entidades.Processo, error)
	ListarProcessosPorCliente(uuid.UUID) ([]entidades.Processo, error)
}
type IRepositorioProcesso interface {
	Fechar()
	ValidarProcesso(uuid.UUID) error
}

type IRepositorioTarefa interface {
	Fechar()
	CriarTarefa(entidades.Tarefa) error
	ListarTarefas() ([]entidades.Tarefa, error)
	ListarTarefasPorProcesso(uuid.UUID) ([]entidades.Tarefa, error)
	EditarTarefa(entidades.Tarefa) error
	DeletarTarefa(uuid.UUID) error
	BuscarTarefaPorUUID(UUID uuid.UUID) (*entidades.Tarefa, error)
	DeletarTarefasPorProcesso(UUID uuid.UUID) error
}
type IRepositorioTarefa interface {
	Fechar()
	CriarTarefa(entidades.Tarefa) error
	ListarTarefas() ([]entidades.Tarefa, error)
	ListarTarefasPorProcesso(uuid.UUID) ([]entidades.Tarefa, error)
	ListarTarefasPorResponsavel(uuid.UUID) ([]entidades.Tarefa, error)
	EditarTarefa(entidades.Tarefa) error
	DeletarTarefa(uuid.UUID) error
	BuscarTarefaPorUUID(UUID uuid.UUID) (*entidades.Tarefa, error)
}
*/
