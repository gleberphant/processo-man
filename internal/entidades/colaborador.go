package entidades

type Colaborador struct {
	Usuario `json:"usuario,omitempty"  db:"usuario"`
}
