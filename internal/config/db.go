package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/AbassAdeyemi/bookmarks/db"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
)

func GetDb(config AppConfig, logger *Logger) *pgx.Conn {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Db.Host, config.Db.Port, config.Db.UserName, config.Db.Password, config.Db.Database)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		logger.Fatal(err)
	}
	applyDbMigration(config, logger)
	return conn
}

func applyDbMigration(config AppConfig, logger *Logger) {
	d, err := iofs.New(db.MigrationFS, "migrations")
	if err != nil {
		logger.Fatalf("Error loading migrations from sources: %v", err)
	}
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Db.UserName, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Database)
	m, err := migrate.NewWithSourceInstance("iofs", d, databaseURL)
	if err != nil {
		logger.Fatalf("Error loading db migrations: %v", err)
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatalf("Error while applying db migrations: %v", err)
	}

	logger.Infof("Database migrations applied successfully")
}
