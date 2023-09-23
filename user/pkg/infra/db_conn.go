package infra

import (
	"fmt"
	"go-micro-sample/user/pkg/lib/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
)

func NewPostgresConnector(cfg *config.DBConfig) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)
	conn, err := sqlx.Connect(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
