package postgres

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	host     = "localhost"
	port     = "5051"
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func ConnectDefaultDataBase() (*bun.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	db.SetMaxOpenConns(90)

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
