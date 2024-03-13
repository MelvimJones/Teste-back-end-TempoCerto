package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Adicione a função de manipulação para a verificação de disponibilidade
func VerificarDisponibilidade(c *gin.Context) {
	disponibilidade := ObterDisponibilidade()
	c.JSON(200, disponibilidade)
}

// GetAvailability é uma função auxiliar para obter a disponibilidade
func GetAvailability() []map[string]interface{} {
	availabilities := []map[string]interface{}{}

	// Lógica para verificar a disponibilidade de horários
	currentHour, currentMinute := time.Now().Hour(), time.Now().Minute()

	for hour := 8; hour < 18; hour++ {
		startTime := fmt.Sprintf("%02d:00", hour)
		endTime := fmt.Sprintf("%02d:00", hour+1)

		// Verificar disponibilidade
		futureTime := hour > currentHour || (hour == currentHour && currentMinute < 30)

		availability := map[string]interface{}{
			"start":     startTime,
			"end":       endTime,
			"available": futureTime,
		}
		availabilities = append(availabilities, availability)
	}

	return availabilities
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
func DisponibilidadeHandler(c *gin.Context) {
	// Lógica para verificar a disponibilidade de horários aqui
	c.JSON(http.StatusOK, gin.H{
		"message": "Verificação de disponibilidade",
	})
}
