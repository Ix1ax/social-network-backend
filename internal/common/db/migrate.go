package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ix1ax/social-network-backend/internal/common/config"
)

func RunMigrations(cfg *config.Config) error {

	dsn := fmt.Sprintf(
		"pgx5://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser,     // USER
		cfg.DbPassword, // PASSWORD
		cfg.DbHost,     // HOST
		cfg.DbPort,     // PORT
		cfg.DbName,     // DBNAME
	)

	m, err := migrate.New("file://db/migration", dsn)

	if err != nil {
		return fmt.Errorf("failed to create migration: %w", err)
	}

	defer m.Close()

	/*
	 * err живет внутри if , блок if с инициализацией в условии
	 */
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to create migration: %w", err)
	}

	return nil

}
