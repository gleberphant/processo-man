package entidades

type Token struct {
	ID    int    `json:"id,omitempty"  db:"id"`
	Nome  string `json:"nome,omitempty"  db:"nome"`
	Chave string `json:"chave,omitempty"  db:"chave"`
}
