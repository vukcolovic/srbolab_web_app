package driver

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	SQL *sqlx.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 25
const maxIdleDbConn = 25

// ConnectSqliteDb creates database pool for db
func ConnectSqliteDb() (*DB, error) {
	db, err := sqlx.Connect("sqlite3", "../../srbolab.db3")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	dbConn.SQL = db

	return dbConn, err
}