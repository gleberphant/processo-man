-- Active: 1779363427406@@127.0.0.1@3306
-- CONFIGURACAO DA TABELA DE TOKENS

-- resetar a tabela
DROP TABLE IF EXISTS tarefas;
 
    
-- criacao da tabela
CREATE TABLE IF NOT EXISTS tarefas (
    uuid TEXT UNIQUE PRIMARY KEY,
    processo_uuid TEXT DEFAULT '00000000-0000-0000-0000-000000000000',
    responsavel_uuid TEXT DEFAULT '00000000-0000-0000-0000-000000000000',
    nome TEXT NOT NULL,
    concluida BOOLEAN DEFAULT false,
    data_criacao DATETIME DEFAULT CURRENT_TIMESTAMP,
    data_conclusao DATETIME DEFAULT NULL,
    comentarios TEXT,
    FOREIGN KEY (processo_uuid) REFERENCES processos(uuid) ON DELETE CASCADE,
    FOREIGN KEY (responsavel_uuid) REFERENCES usuarios(uuid) ON DELETE CASCADE
);

-- inserir dados para teste
-- INSERT INTO tokens (uuid, usuario_uuid) VALUES ('teste','teste');
