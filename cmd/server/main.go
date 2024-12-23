package main

import (
	"log"

	_ "tspo_final/docs"
	"tspo_final/internal/db"
	"tspo_final/internal/models"
	"tspo_final/internal/route"
)

func main() {
	db, err := db.ConnectToDB("db", "user", "user", "5432", "orders_db")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Order{})

	route := route.SetupRoutes(db)

	route.Run(":8080")
}
