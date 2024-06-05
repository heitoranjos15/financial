package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

  "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)



func GetConnection(context context.Context) *pgxpool.Pool {
  databaseURL := "root:root@localhost:5432/financial?sslmode=disable"

  conn, err := pgxpool.New(context, "postgres://"+databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func RunMigrations() {
  databaseURL := "root:root@localhost:5432/financial?sslmode=disable"
  m, err := migrate.New("file://database/migrations", "postgres://"+databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}

func DownMigrations() {
  databaseURL := "root:root@localhost:5432/financial?sslmode=disable"
  m, err := migrate.New("file://database/migrations", "postgres://"+databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err := m.Down(); err != nil {
		log.Println(err)
	}
}
