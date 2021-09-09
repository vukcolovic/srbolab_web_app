package dbRepo

import (
	"github.com/jmoiron/sqlx"
	"srbolab/internal/repository"
)

type sqliteDBRepo struct {
	DB  *sqlx.DB
}

// NewSqliteRepo creates the repository
func NewSqliteRepo(Conn *sqlx.DB) repository.DatabaseRepo {
	return &sqliteDBRepo{
		DB:  Conn,
	}
}
