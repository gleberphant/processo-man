-- Active: 1779363427406@@127.0.0.1@3306
-- CONFIGURACAO DA TABELA DE TOKENS

-- resetar a tabela
DROP TABLE IF EXISTS administradores;
 
    
-- criacao da tabela
CREATE TABLE IF NOT EXISTS administradores (
    id INTEGER PRIMARY KEY,
    usuario_uuid TEXT UNIQUE NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    cpf TEXT NOT NULL DEFAULT '',
    cargo TEXT NOT NULL DEFAULT '',
    data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_uuid) REFERENCES usuarios(uuid) ON DELETE CASCADE
);

-- inserir dados para teste
INSERT INTO
    administradores (usuario_uuid, cpf, cargo)
VALUES ('00000000-0000-0000-0000-000000000000', 'teste', 'teste');
-- INSERT INTO tokens (uuid, usuario_uuid) VALUES ('teste','teste');
