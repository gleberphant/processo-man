-- Active: 1778958784571@@127.0.0.1@3306
-- CONFIGURACAO DA TABELA DE TOKENS

-- resetar a tabela
DROP TABLE IF EXISTS tarefas;
 
    
-- criacao da tabela
CREATE TABLE IF NOT EXISTS tarefas (
    uuid TEXT UNIQUE PRIMARY KEY,
    processo_uuid TEXT,
    responsavel_uuid TEXT,
    nome TEXT NOT NULL,
    concluida BOOLEAN DEFAULT false,
    data_criacao DATETIME DEFAULT CURRENT_TIMESTAMP,
    data_conclusao DATETIME DEFAULT NULL,
    comentarios TEXT,
    FOREIGN KEY (processo_uuid) REFERENCES processos(uuid),
    FOREIGN KEY (responsavel_uuid) REFERENCES usuarios(uuid)
);

-- inserir dados para teste
-- INSERT INTO tokens (uuid, usuario_uuid) VALUES ('teste','teste');
