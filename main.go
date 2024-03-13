package main

import (
	"AgendaService/database"
	"AgendaService/handlers"
	"AgendaService/models"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Estrutura Config armazena as configurações do banco de dados
type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func GetConfig() Config {
	// Lógica para obter as configurações do banco de dados
	return Config{
		DBUsername: "agenda_user",
		DBPassword: "senha_segura",
		DBHost:     "localhost",
		DBPort:     "3306",
		DBName:     "agenda_db",
	}
}

func main() {
	config := GetConfig()

	// String de conexão do MySQL
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName)

	// Conexão com o banco de dados
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
	}

	// Teste da conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao testar a conexão com o banco de dados:", err)
	}

	defer db.Close()

	router := gin.Default()

	// Verificar Disponibilidade
	router.GET("/agendas:disponibilidade", handlers.VerificarDisponibilidade)

	// Rota teste de servidor 8080
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Servidor iniciado com sucesso!"})
	})

	// Rota para listar e criar agendas
	router.Handle("POST", "/agendas", func(c *gin.Context) {
		var novaAgenda models.Agenda
		if err := c.ShouldBindJSON(&novaAgenda); err != nil {
			log.Println("Erro ao fazer o bind do JSON:", err)
			c.JSON(400, gin.H{"error": "Dados inválidos"})
			return
		}

		id, err := database.CriarAgenda(db, novaAgenda.Empresa.Cnpj, novaAgenda.Horario)
		if err != nil {
			log.Println("Erro ao criar a agenda no banco de dados:", err)
			c.JSON(500, gin.H{"error": "Erro ao criar a agenda no banco de dados"})
			return
		}

		c.JSON(201, gin.H{"id": id})
	})

	// Rota que retorna as agendas
	router.GET("/agendas", func(c *gin.Context) {
		agendas, err := database.ListarAgendas(db)
		if err != nil {
			log.Println("Erro ao obter as agendas do banco de dados:", err)

			// Verifica se c.Errors e c.Errors.Last() não são nulos antes de acessar JSON()
			if c.Errors != nil && len(c.Errors) > 0 && c.Errors.Last() != nil {
				// Recupera a resposta JSON de erro
				errorResponse, err := c.Errors.Last().MarshalJSON()
				if err != nil {
					log.Println("Erro ao obter a resposta JSON de erro:", err)
				} else {
					// Imprime a resposta JSON de erro no log
					log.Printf("Resposta JSON de erro: %s", errorResponse)
				}

				// Verifica se a resposta contém a mensagem de erro "Too many requests"
				if strings.Contains(string(errorResponse), "Too many requests") {
					c.JSON(429, gin.H{"error": "Muitas solicitações, tente novamente mais tarde"})
					return
				}
			}

			c.JSON(500, gin.H{
				"status":   429,
				"mensagem": "Muitas solicitações foram feitas à API Receita WS. Por favor, aguarde alguns minutos.",
			})
			return
		}

		c.JSON(200, agendas)
	})

	// Rota para verificar disponibilidade
	router.Handle("GET", "/agendas/disponibilidade", func(c *gin.Context) {
		// Lógica de verificação de disponibilidade aqui
		c.JSON(200, gin.H{"message": "Verificar disponibilidade"})
	})

	router.Run(":8080")
}
