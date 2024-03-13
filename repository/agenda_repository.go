package repository

import (
	"database/sql"
)

type AgendaRepository struct {
	DB *sql.DB
}
