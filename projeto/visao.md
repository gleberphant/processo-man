# Visão

## descrição

Sistemas para controle de processos. 

## Requisitos

| ID Requisito | Módulo / Área | Requisito Funcional (Descrição) | Ator Principal | Dependências / Vínculos (`<<inclui>>`) |
| :--- | :--- | :--- | :--- | :--- |
| **RF-01** | Área de Administração | Manter usuário (Cadastrar, Alterar, Consultar, Excluir) | Administrador do Sistema | - |
| **RF-02** | Área de Administração | Atribuir perfil a um usuário | Administrador do Sistema | - |
| **RF-03** | Secretaria | Manter clientes | Secretaria | - |
| **RF-04** | Secretaria | Manter documentos de um cliente | Secretaria | - |
| **RF-05** | Secretaria | Manter processos | Secretaria | - |
| **RF-06** | Secretaria / Cliente | Atribuir cliente a um processo | Secretaria / Cliente | - |
| **RF-07** | Área do Colaborador | Ver detalhes de um processo | Colaborador | Inclui **RF-08** e **RF-15** |
| **RF-08** | Área do Colaborador | Listar tarefas no processo | Colaborador | *(Invocado por RF-07)* |
| **RF-09** | Área do Colaborador | Adicionar tarefa a um processo | Colaborador | Inclui **RF-10** |
| **RF-10** | Área do Colaborador | Atribuir um responsável para uma tarefa | Colaborador | *(Invocado por RF-09)* |
| **RF-11** | Área do Colaborador | Manter suas tarefas | Colaborador | - |
| **RF-12** | Área do Colaborador | Concluir uma tarefa (Marcar como concluída) | Colaborador | - |
| **RF-13** | Área do Colaborador | Adicionar documento a um processo | Colaborador | - |
| **RF-14** | Área do Colaborador | Manter seus documentos | Colaborador | - |
| **RF-15** | Área do Colaborador | Listar documento no processo | Colaborador | *(Invocado por RF-07)* |
| **RF-16** | Área do Cliente | Listar seus processos | Cliente | - |
| **RF-17** | Área do Cliente | Ver detalhes dos seus processos | Cliente | - |
| **RF-18** | Área do Cliente | Ver seus documentos | Cliente | - |
| **RF-19** | Área do Usuário (Geral) | Ver seus dados | Usuário (Todos) | - |
| **RF-20** | Área do Usuário (Geral) | Corrigir seus dados | Usuário (Todos) | - |
| **RF-21** | Autenticação | Efetuar Login (Validar credenciais e Token) | Usuário (Todos) | - |
| **RF-22** | Autenticação | Efetuar Logout | Usuário (Todos) | - |

## **Instruções para o Desenvolvimento**

### Fluxo de Dados

* O Roteador recebe a requisição e envia para o Manipulador

* O Manipulador recebe `(w, r)` e extrai os dados (do formulário HTML ou da URL) para a `struct` de entrada e Chama a lógica no `Servico...`..

* O Serviço valida as regras de negócio, executa a função e chama o repositório para persistência dos dados.

* O Repositório (dentro do Serviço) salva no banco.

* O Serviço devolver resposta para o manipulador

* O Manipulador decide se envia um ExecuteTemplate (HTML) ou um json.Encode (API).

### Banco de Dados**

  um script `schema.sql` com os comandos `CREATE TABLE` baseados no seu Diagrama de Entidade-Relacionamento 