package entidades

type Cliente struct {
	Usuario `json:"usuario,omitempty"  db:"usuario"`
}
