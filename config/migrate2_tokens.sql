-- CONFIGURACAO DA TABELA DE TOKENS

-- resetar a tabela
DROP TABLE IF EXISTS tokens;

    
-- criacao da tabela
CREATE TABLE IF NOT EXISTS tokens (
    uuid TEXT UNIQUE PRIMARY KEY,
    token TEXT NOT NULL,
    usuario_uuid TEXT 
    perfis TEXT NOT NULL DEFAULT 'colaborador',
    data_criacao DATETIME DEFAULT CURRENT_TIMESTAMP,
    ativo BOOLEAN NOT NULL DEFAULT 0,
    validade TEXT DEFAULT 'permanente',
    comentarios TEXT,
    FOREIGN KEY (usuario_uuid) REFERENCES usuarios(uuid)
    CONSTRAINT tipos_validade CHECK (
        validade IN ('permanente', 'temporario')
    )
);

-- inserir dados para teste
INSERT INTO tokens (uuid, token) VALUES ('teste','teste');
