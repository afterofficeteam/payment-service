package main

import (
	"database/sql"
	"payment-service/config"
	"payment-service/routes"

	"github.com/go-playground/validator/v10"

	handler "payment-service/handlers"
	repo "payment-service/repository"
	svc "payment-service/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	sqlDb, err := config.ConnectToDatabase(config.Connection{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		return
	}
	defer sqlDb.Close()

	validator := validator.New()
	routes := setupRoutes(sqlDb, validator)
	routes.Run(cfg.AppPort)
}

func setupRoutes(myDb *sql.DB, validator *validator.Validate) *routes.Routes {
	store := repo.NewStore(myDb)
	services := svc.Newsvc(store)
	handler := handler.NewHandler(services, validator)

	return &routes.Routes{}
}
