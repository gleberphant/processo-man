-- CONFIGURACAO DA TABELA DE USUARIOS

-- resetar a tabela
DROP TABLE IF EXISTS usuarios;

-- criação da tabela
CREATE TABLE IF NOT EXISTS usuarios (
    uuid TEXT UNIQUE PRIMARY KEY,
    nome TEXT,
    email TEXT UNIQUE NOT NULL,
    senha TEXT NOT NULL,
    perfis TEXT NOT NULL DEFAULT 'colaborador',
    data_criacao DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- inserir dados para teste
-- INSERT INTO
--     usuarios (uuid, nome, email, senha)
-- VALUES (
--         'teste',
--         'teste',
--         'teste@teste',
--         'teste'
--     );