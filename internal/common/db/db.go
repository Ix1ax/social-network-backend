package db

import (
	"fmt"

	// Ничего не используем из пакета, вызываем ради эффекта init
	"github.com/ix1ax/social-network-backend/internal/common/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewConnection(cfg *config.Config) (*sqlx.DB, error) {

	/*
	 * Создаем строку с нужным форматированием
	 * search_path - это в какой схеме искать
	 */
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable search_path=social_network",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
		cfg.DbUser,
		cfg.DbPassword,
	)

	// Подключаем с драйвером pgx и строкой раньше сформированной
	db, err := sqlx.Connect("pgx", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil

}
