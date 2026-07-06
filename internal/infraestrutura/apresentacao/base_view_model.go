package apresentacao

// UsuarioLogadoViewModel contém os dados do usuário que serão exibidos no template.
type UsuarioLogadoViewModel struct {
	Nome string
}

// BaseViewModel é uma estrutura que pode ser embutida em outros ViewModels para fornecer dados comuns.
type BaseViewModel struct {
	UsuarioLogado UsuarioLogadoViewModel
}
