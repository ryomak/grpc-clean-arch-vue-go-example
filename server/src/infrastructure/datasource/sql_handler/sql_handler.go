package sql_handler

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	DB *sql.DB
}

func NewSqlHandler(path, env string) *SqlHandler {
	cs, err := NewConfigsFromFile(path)
	if err != nil {
		log.Fatalf("cannot open database configuration. -> %s", err)
	}
	db, err := cs.OpenDB(env)
	return &SqlHandler{DB: db}
}
