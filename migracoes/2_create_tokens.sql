-- Active: 1778527634874@@127.0.0.1@3306
-- CONFIGURACAO DA TABELA DE TOKENS

-- resetar a tabela
DROP TABLE IF EXISTS tokens;
 
    
-- criacao da tabela
CREATE TABLE IF NOT EXISTS tokens (
    uuid TEXT UNIQUE PRIMARY KEY,
    usuario_uuid TEXT 
    perfis TEXT NOT NULL DEFAULT 'colaborador',
    data_criacao DATETIME DEFAULT CURRENT_TIMESTAMP,
    validade TEXT DEFAULT 'temporario',
    comentarios TEXT,
    FOREIGN KEY (usuario_uuid) REFERENCES usuarios(uuid)
    CONSTRAINT tipos_validade CHECK (
        validade IN ('permanente', 'temporario')
    )
);

-- inserir dados para teste
-- INSERT INTO tokens (uuid, usuario_uuid) VALUES ('teste','teste');
