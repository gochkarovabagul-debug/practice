package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func ConnectDB(config string) {
	conn, err := pgx.Connect(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Db error : %v\n", err)
		os.Exit(1)
	}
	db = conn
}

func GetDB() *pgx.Conn {
	return db
}
