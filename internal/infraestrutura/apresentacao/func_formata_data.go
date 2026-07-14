package apresentacao

import (
	"time"
)

// 1. Crie a função de formatação de formatação de data
func formatarData(data time.Time) string {
	if data.IsZero() {
		return "Sem Data" // ou retorne "" se preferir vazio
	}
	// Formato padrão brasileiro
	return data.Format("02/01/2006")
}
