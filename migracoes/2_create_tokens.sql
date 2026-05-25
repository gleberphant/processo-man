-- Active: 1779363427406@@127.0.0.1@3306
-- CONFIGURACAO DA TABELA DE TOKENS

-- resetar a tabela
DROP TABLE IF EXISTS tokens;

-- criacao da tabela
CREATE TABLE IF NOT EXISTS tokens (
    uuid TEXT UNIQUE PRIMARY KEY DEFAULT (
        printf(
            '%s-%s-%s-%s-%s',
            lower(hex(randomblob(4))),
            lower(hex(randomblob(2))),
            lower(hex(randomblob(2))),
            lower(hex(randomblob(2))),
            lower(hex(randomblob(6)))
        )
    ),
    usuario_uuid TEXT NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    perfis TEXT NOT NULL DEFAULT 'colaborador',
    data_criacao DATETIME DEFAULT CURRENT_TIMESTAMP,
    validade TEXT DEFAULT 'temporario',
    comentarios TEXT,
    FOREIGN KEY (usuario_uuid) REFERENCES usuarios (uuid) CONSTRAINT tipos_validade CHECK (
        validade IN ('permanente', 'temporario')
    )
);

-- inserir dados para teste
INSERT INTO
    tokens (comentarios)
VALUES ('token inicial para testes');