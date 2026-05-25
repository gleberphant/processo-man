-- Active: 1779363427406@@127.0.0.1@3306
-- CONFIGURACAO DA TABELA DE TOKENS

-- resetar a tabela
DROP TABLE IF EXISTS clientes;
 
    
-- criacao da tabela
CREATE TABLE IF NOT EXISTS clientes (
    id INTEGER PRIMARY KEY,  --TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    usuario_uuid TEXT UNIQUE NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    cpf TEXT NOT NULL DEFAULT '000.000.000-00',
    endereco TEXT NOT NULL DEFAULT '',
    tipo_pessoa TEXT NOT NULL DEFAULT 'fisica',
    data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_uuid) REFERENCES usuarios(uuid) ON DELETE CASCADE
);

-- inserir dados para teste
-- INSERT INTO tokens (uuid, usuario_uuid) VALUES ('teste','teste');
