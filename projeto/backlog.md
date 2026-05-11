| check | **Sprint** | **Módulo**   | **ID** | **Tarefa Técnica**              | **Detalhamento / Critério de Aceite**                                               |
| --------------------------------- | ---------- | ------------ | ------ | ------------------------------- | ----------------------------------------------------------------------------------- |
| [x]                               | **1**      | **Infra**    | 1.1    | **Configuração do projeto**     | Iniciar `go mod`, instalar pacotes `uuid`, `chi` (ou router escolhido) e `sqlite3`. |
| [x]                               | **1**      | **Infra**    | 1.2    | **Router Base**                 | Implementar o `Roteador` para gerenciar as rotas iniciais.                          |
| [x]                               | **1**      | **Infra**    | 1.2    | **Singleton do Banco de Dados** | Criar struct `BancoDeDados` e função `Conectar()` que retorna `*sql.DB`.            |
| [x]                               | **1**      | **Entidade** | 1.3    | **Structs de Usuário/Perfil**   | Implementar as structs `Usuario`, `Perfil` e `Permissao` com tags JSON.             |
| []                                | **1**      | **Manip**    | 1.4    | **Handler de Login (GET)**      | Método `PageLogin`: Carregar e renderizar o template HTML de login.                 |
| []                                | **1**      | **Serviço**  | 1.5    | **Lógica de Auth**              | No `ServicoUsuario`, criar método `Autenticar(email, senha)`.                       |
| []                                | **2**      | **Entidade** | 2.1    | **Structs de Negócio**          | Implementar `Cliente`, `Processo` e `Tarefa` com Embedding de Usuario.              |
| []                                | **2**      | **Serviço**  | 2.2    | **CRUD de Clientes**            | Implementar `CriarCliente` e `ListarClientes` no `ServicoUsuario`.                  |
| []                                | **2**      | **Manip**    | 2.3    | **Página de Cadastro Cliente**  | Método `PageCriarCliente` e `CriarCliente` (POST) no manipulador.                   |
| []                                | **2**      | **Infra**    | 2.4    | **Middleware de Sessão**        | Criar função que verifica se o usuário está logado antes de acessar rotas "Page".   |
| []                                | **3**      | **Serviço**  | 3.1    | **Gestão de Processos**         | Implementar `CriarProcesso` vinculando o `ID` do cliente.                           |
| []                                | **3**      | **Manip**    | 3.2    | **Fluxo de Novo Processo**      | Handler `PageCriarProcesso` e lógica de persistência no `ControleProcesso`.         |
| []                                | **3**      | **Entidade** | 3.3    | **Entidade Fraca: Lista**       | Implementar a `Lista de Tarefas` como dependente de um `ID_PROCESSO`.               |
| []                                | **4**      | **Serviço**  | 4.1    | **Lógica de Tarefas**           | Implementar `InserirTarefa` e `AtualizarStatusTarefa` (concluída/aberta).           |
| []                                | **4**      | **Manip**    | 4.2    | **Visualização do Projeto**     | Handler `PageVisualizar` que busca Processo + Lista + Tarefas em um único DTO.      |
| []                                | **4**      | **Front**    | 4.3    | **Área do Cliente (HTML)**      | Criar template para o caso de uso "Acompanhar Andamento".                           |
| []                                | **5**      | **API**      | 5.1    | **Refatoração JSON**            | Adicionar `w.Header().Set("Content-Type", "application/json")` nos métodos de ação. |

### **Instruções para o Desenvolvimento**

#### **1. O Ciclo do Manipulador (Handler)**

Para cada funcionalidade acima, siga este fluxo de implementação dentro do Go:

* **Recepção:** O Manipulador recebe `(w, r)`.

* **Parsing:** Lê o formulário ou JSON para a `struct` de entrada.

* **Serviço:** Chama a lógica no `Servico...`.

* **Resposta:** Se `r.Header` pedir HTML, renderiza template; se não, envia JSON.

#### **2. Simplificação do Banco de Dados**

Como você quer baixar a granularidade, não gaste tempo com migrações complexas agora. Crie um script `schema.sql` com os comandos `CREATE TABLE` baseados no seu Diagrama de Entidade-Relacionamento e execute-o manualmente no SQLite.
