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
func DisponibilidadeHandler(c *gin.Context) {
	// Lógica para verificar a disponibilidade de horários aqui
	c.JSON(http.StatusOK, gin.H{
		"message": "Verificação de disponibilidade",
	})
}
