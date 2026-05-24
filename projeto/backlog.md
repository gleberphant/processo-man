# Backlog

 [x] **Configuração do projeto**      Iniciar `go mod`, instalar pacotes `uuid`, `chi` (ou router escolhido) e `sqlite3`.
 [x] **Router Base**                  Implementar o `Roteador` para gerenciar as rotas iniciais.
 [x] **Singleton do Banco de Dados**  Criar struct `BancoDeDados` e função `Conectar()` que retorna `*sql.DB`.
 [x] **Structs de Usuário/Perfil**    Implementar as structs `Usuario`, `Perfil` e `Permissao` com tags JSON.
 [x] **Handler de Login (GET)**       Método `PageLogin`: Carregar e renderizar o template HTML de login.
 [x] **Lógica de Auth**               No `ServicoUsuario`, criar método `Autenticar(email, senha)`.
 [x] **Structs de Negócio**           Implementar `Cliente`, `Processo` e `Tarefa` com Embedding de Usuario.
 [x] **CRUD de Clientes**             Implementar `CriarCliente` e `ListarClientes` no `ServicoUsuario`.
 [x] **Página de Cadastro Cliente**   Método `PageCriarCliente` e `CriarCliente` (POST) no manipulador.
 [x] **Middleware de Sessão**         Criar função que verifica se o usuário está logado antes de acessar rotas "Page".  
 [x] **Gestão de Processos**          Implementar `CriarProcesso` vinculando o `ID` do cliente.
 [x] **Fluxo de Novo Processo**       Handler `PageCriarProcesso` e lógica de persistência no `ControleProcesso`.
 [x] **Visualização do Projeto**      Handler `PageVisualizar` que busca Processo + Lista + Tarefas em um único DTO.
 [x] **Entidade Fraca: Lista**        Implementar a `Lista de Tarefas` como dependente de um `ID_PROCESSO`.
 [x] **Lógica de Tarefas**            Implementar `InserirTarefa` e `AtualizarStatusTarefa` (concluída/aberta).
 [x] **Processo**                     Implementar `Editar` Processo
 [x] **Processo**                     Implementar `Deletar` Processo
 [x] **Tarefas**                      Implementar `Editar` `tarefa`
 [x] **Tarefas**                      Implementar `Deletar` tarefa
 [x] **Tarefas**                      Designar um `responsável` pela tarefa
 [] **Tipos de Usuarios**             Implementar `Cliente` e `Colaborador`
 [] **Autenticacao com perfil**       Inserir o perfil (`Cliente` e `Colaborador`) no token
 [] **Área do Cliente (HTML)**        Criar template para o cliente `listar meus processo` e `ver meus processo`
 [] **Refatoração JSON**              Adicionar `w.Header().Set("Content-Type", "application/json")` nos métodos de ação.
