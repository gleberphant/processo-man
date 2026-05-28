-- Active: 1778958784571@@127.0.0.1@3306
-- CONFIGURACAO DA TABELA DE PROCESSOS

-- DESABILITA TEMPORARIAMENTE AS CONSTRAINTS
PRAGMA foreign_keys=off;

-- INICIO DA TRANSAÇÃO
BEGIN TRANSACTION;

-- backup da tabela - para migração de dados
ALTER TABLE IF EXISTS processos RENAME TO processos_old;
    
-- criacao da tabela
CREATE TABLE IF NOT EXISTS  [processos] ( 
  [uuid] TEXT NOT NULL UNIQUE,
  [nome] TEXT NOT NULL DEFAULT '' ,
  [cliente_uuid] TEXT NULL,
  [colaborador_uuid] TEXT NULL,
  [data_criacao] DATETIME NULL DEFAULT CURRENT_TIMESTAMP ,
  [comentarios] TEXT NOT NULL DEFAULT '' ,
   PRIMARY KEY ([uuid]),
  FOREIGN KEY ([colaborador_uuid]) REFERENCES [usuarios] ([uuid]) ON DELETE NO ACTION ON UPDATE NO ACTION,
  FOREIGN KEY ([cliente_uuid]) REFERENCES [usuarios] ([uuid]) ON DELETE NO ACTION ON UPDATE NO ACTION
);

-- RESTAURA DADOS ANTIGOS
INSERT INTO processos(uuid, nome, cliente_uuid, data_criacao)
SELECT p.uuid, p.nome, u.uuid, p.data_criacao 
FROM processos_old p
LEFT JOIN usuarios u ON p.dono = u.uuid;

DROP TABLE IF EXISTS processos_old;

COMMIT;

PRAGMA foreign_keys=ON;