-- CONFIGURACAO DA TABELA DE USUARIOS

DROP TABLE IF EXISTS usuarios;

-- criação da tabela
CREATE TABLE IF NOT EXISTS usuarios (
    uuid TEXT UNIQUE PRIMARY KEY,
    nome TEXT NOT NULL DEFAULT 'NOME',
    email TEXT UNIQUE NOT NULL,
    senha TEXT NOT NULL,
    perfis TEXT NOT NULL DEFAULT 'colaborador',
    data_criacao DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- insere usuario teste
INSERT INTO
    usuarios (uuid, email, senha)
VALUES ('00000000-0000-0000-0000-000000000000', 'teste@teste', 'teste');