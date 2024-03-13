package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// RespostaAgenda representa a estrutura de dados para a resposta da API
type RespostaAgenda struct {
	Horario string `json:"horario"`
	Empresa struct {
		Cnpj string `json:"cnpj"`
		Nome string `json:"nome"`
	} `json:"empresa"`
}

// ListarAgendas retorna a lista de agendas do banco de dados
func ListarAgendas(db *sql.DB) ([]RespostaAgenda, error) {
	rows, err := db.Query("SELECT cnpj, horario FROM agendas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []RespostaAgenda
	for rows.Next() {
		var agenda RespostaAgenda
		err := rows.Scan(&agenda.Empresa.Cnpj, &agenda.Horario)
		if err != nil {
			return nil, err
		}

		// Consultar API Receita WS
		nomeFantasia, err := consultarReceitaWS(agenda.Empresa.Cnpj)
		if err != nil {
			// Tratar erro da consulta à API Receita WS
			return nil, err
		}

		// Atualizar o nome fantasia no objeto agenda
		agenda.Empresa.Nome = nomeFantasia
		agendas = append(agendas, agenda)
	}
	return agendas, nil
}

// consultarReceitaWS consulta a API Receita WS para obter o nome fantasia
func consultarReceitaWS(cnpj string) (string, error) {
	url := "https://receitaws.com.br/v1/cnpj/" + cnpj
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Adicione um log para imprimir a resposta da API Receita WS
	log.Printf("Resposta da API Receita WS: %s", body)

	// Parse da resposta JSON
	var respostaAPI map[string]interface{}
	err = json.Unmarshal(body, &respostaAPI)
	if err != nil {
		return "", err
	}

	// Extrair o nome fantasia da resposta
	nomeFantasia, ok := respostaAPI["fantasia"].(string)
	if !ok || nomeFantasia == "" {
		// Se o nome fantasia não estiver disponível, use o nome padrão
		nome, ok := respostaAPI["nome"].(string)
		if !ok {
			return "", fmt.Errorf("Campos 'fantasia' e 'nome' não encontrados na resposta da API")
		}
		return nome, nil
	}

	return nomeFantasia, nil
}

// CriarAgenda cria uma nova agenda no banco de dados
func CriarAgenda(db *sql.DB, cnpj string, horario string) (int64, error) {
	// Implementação de exemplo para criar uma agenda
	result, err := db.Exec("INSERT INTO agendas (cnpj, horario) VALUES (?, ?)", cnpj, horario)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
