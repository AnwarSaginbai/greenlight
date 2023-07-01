package data

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func OpenDB(cfg dbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.String())
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DefaultPostgresConfig() dbConfig {
	return dbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "greenlight",
		Password: "pa55word",
		Database: "greenlight",
		SSLMode:  "disable",
	}
}

func (cfg dbConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}
