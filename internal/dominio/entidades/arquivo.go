package entidades

import (
	"time"

	"github.com/google/uuid"
)

type Arquivo struct {
	UUID        uuid.UUID
	Nome        string
	Usuario     *Usuario
	Processo    *Processo
	DataCriacao time.Time
	URL         string
}
