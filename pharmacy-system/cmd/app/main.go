package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"pharmacy-system/internal/repository"
	"pharmacy-system/internal/service"
)

func main() {
	fmt.Println("Pharmacy System")

	// Подключение к базе данных
	db, err := connectToDB()
	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}

	// Инициализация репозитория
	newRepository := repository.NewRepository(db)
	// Инициализация сервиса
	newService := service.NewService(*newRepository)
	//
	mux := http.NewServeMux()
	// Инициализация обработчика
	newHandler := handler.NewHandler(mux, newService)
	newHandler.InitRoutes()

	fmt.Printf("Server is starting... address: %v", ":8080\n")
	err = http.ListenAndServe("localhost:8080", newHandler)
	if err != nil {
		panic(err)
	}

}

// connectToDB connects to the PostgresSQL database using the provided DSN.
//
// It returns a pointer to a gorm.DB object and an error if any occurred.
func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=humo_db port=5432 sslmode=disable TimeZone=Asia/Dushanbe"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
