package model

import (
	"fmt"

	"github.com/go-spatial/tegola/internal/ttools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgDb *gorm.DB

func GetPgDb() (*gorm.DB, error) {
	if PgDb == nil {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			ttools.GetEnvDefault("DB_HOST", "localhost"),
			ttools.GetEnvDefault("DB_USER", "postgres"),
			ttools.GetEnvDefault("DB_PASSWORD", "postgres"),
			ttools.GetEnvDefault("DB_NAME", "sobo"),
			ttools.GetEnvDefault("DB_PORT", "5432"),
			ttools.GetEnvDefault("TEGOLA_POSTGIS_SSL", "disable"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			return nil, err
		}

		// migrate table
		// db.AutoMigrate(&TLayer{})
		PgDb = db
	}

	return PgDb, nil
}
