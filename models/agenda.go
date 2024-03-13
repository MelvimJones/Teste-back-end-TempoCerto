package models

// Agenda representa a estrutura de dados para uma agenda
type Agenda struct {
	Empresa struct {
		Cnpj string `json:"cnpj"`
	} `json:"empresa"`
	Horario string `json:"horario"`
}
