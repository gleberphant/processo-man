-- CONFIGURACAO DA TABELA DE USUARIOS

-- DESABILITA TEMPORARIAMENTE AS CONSTRAINTS
PRAGMA foreign_keys=off;

-- INICIO DA TRANSAÇÃO
BEGIN TRANSACTION;


-- backup da tabela - para migração de dados
ALTER TABLE IF EXISTS usuarios RENAME TO usuarios_old;

-- criação da tabela
CREATE TABLE IF NOT EXISTS usuarios (
    uuid TEXT UNIQUE PRIMARY KEY,
    nome TEXT,
    email TEXT UNIQUE NOT NULL,
    senha TEXT NOT NULL,
    perfis TEXT NOT NULL DEFAULT 'colaborador',
    data_criacao DATETIME DEFAULT CURRENT_TIMESTAMP
);



DROP TABLE IF EXISTS usuarios_old;