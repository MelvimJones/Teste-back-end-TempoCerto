package handlers

import (
	"AgendaService/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	//"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Adicione a função de manipulação para a verificação de disponibilidade
func VerificarDisponibilidade(c *gin.Context) {
	disponibilidade := ObterDisponibilidade()
	c.JSON(200, disponibilidade)
}

// Função auxiliar para obter a disponibilidade
func ObterDisponibilidade() []map[string]interface{} {
	horariosDisponiveis := []map[string]interface{}{}

	// Disponibilidade de horários
	horaAtual := time.Now().Hour()
	minutoAtual := time.Now().Minute()

	for hora := 8; hora < 18; hora++ {
		inicio := fmt.Sprintf("%02d:00", hora)
		fim := fmt.Sprintf("%02d:00", hora+1)

		// Verificar disponibilidade
		horaNoFuturo := hora > horaAtual || (hora == horaAtual && minutoAtual < 30)

		horarioDisponivel := map[string]interface{}{
			"inicio":     inicio,
			"fim":        fim,
			"disponivel": horaNoFuturo,
		}
		horariosDisponiveis = append(horariosDisponiveis, horarioDisponivel)
	}

	return horariosDisponiveis
}

/* ------------------------------------------------------ */

func AgendamentoHandler(c *gin.Context) {
	// Lógica para tratar o agendamento aqui
	c.JSON(http.StatusOK, gin.H{
		"message": "Agendamento realizado com sucesso",
	})
}

// Nova função para listar agendas
func ListarAgendasHandler(c *gin.Context) {
	// Lógica para listar as agendas
	c.JSON(http.StatusOK, gin.H{
		"message": "Listagem de agendas",
	})
}

// Função para verificar a disponibilidade de horários
func DisponibilidadeHandler(c *gin.Context, db *sql.DB) {
	// Consulta o banco de dados para obter os horários existentes
	agendas, err := database.ListarAgendas(db)
	if err != nil {
		log.Println("Erro ao obter as agendas do banco de dados:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter as agendas do banco de dados"})
		return
	}

	// Cria um mapa para armazenar os horários disponíveis
	disponibilidade := make(map[string]bool)
	// Define todos os horários como disponíveis inicialmente
	for hora := 8; hora < 18; hora++ {
		inicio := fmt.Sprintf("%02d:00", hora)
		disponibilidade[inicio] = true
	}

	// Marca os horários existentes como não disponíveis
	for _, agenda := range agendas {
		disponibilidade[agenda.Horario] = false
	}

	// Cria uma lista de objetos JSON para representar a disponibilidade
	var disponibilidadeJSON []gin.H
	for hora, disponivel := range disponibilidade {
		// Adiciona uma hora ao horário de início para obter o horário de fim
		horarioInicio, _ := time.Parse("15:04", hora)
		horarioFim := horarioInicio.Add(time.Hour) // Adiciona uma hora
		fim := horarioFim.Format("15:04")

		disponibilidadeJSON = append(disponibilidadeJSON, gin.H{
			"inicio":     hora,
			"fim":        fim,
			"disponivel": disponivel,
		})
	}

	// Retorna a disponibilidade como resposta JSON
	c.JSON(http.StatusOK, disponibilidadeJSON)
}
