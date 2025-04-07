package database

import (
	"ToGoList/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	*gorm.DB
}

func InitDB(dsn string) *DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connection to Postgres DB, caused by: ", err)
	}
	registerModels(db)

	return &DB{db}
}

// auto migrations for ALL models
func registerModels(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Book{},
		// another models
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}

// Close connection with DB
func (d *DB) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
