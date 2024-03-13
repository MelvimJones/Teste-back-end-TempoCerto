// models/empresa.go
package models

import (
	"encoding/json"
)

type Empresa struct {
	Cnpj        string `json:"cnpj"`
	RazaoSocial string `json:"razao_social"`
}

func ParseJSON(jsonData []byte, v interface{}) error {
	return json.Unmarshal(jsonData, v)
}
