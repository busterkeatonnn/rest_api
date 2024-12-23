package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(dbAlias string, dbUser string, dbPassword string, dbPort string, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		dbAlias, dbUser, dbPassword, dbPort, dbName,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}
