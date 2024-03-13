package main

// Config representa a estrutura de configuração do aplicativo.
type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

// GetConfig retorna uma instância de Config preenchida com os valores adequados.
func GetConfig() *Config {
	return &Config{
		DBUsername: "agenda_user",
		DBPassword: "senha_segura",
		DBHost:     "localhost",
		DBPort:     "3306", // Porta padrão do MySQL
		DBName:     "agenda_db",
	}
}
