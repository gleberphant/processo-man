-- Active: 1778958784571@@127.0.0.1@3306
-- CONFIGURACAO DA TABELA DE TOKENS

-- resetar a tabela
DROP TABLE IF EXISTS processos;
 
    
-- criacao da tabela
CREATE TABLE IF NOT EXISTS processos (
    uuid TEXT UNIQUE PRIMARY KEY,
    nome TEXT,
    dono TEXT, 
    data_criacao DATETIME DEFAULT CURRENT_TIMESTAMP,
    comentarios TEXT
);

-- inserir dados para teste
-- INSERT INTO tokens (uuid, usuario_uuid) VALUES ('teste','teste');
