package main

import (
	"base-project/config"
	"base-project/routes"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/thedevsaddam/renderer"
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
	render := renderer.New()
	routes := setupRoutes(render, sqlDb, validator)
	routes.Run(cfg.AppPort)
}

func setupRoutes(render *renderer.Render, myDb *sql.DB, validator *validator.Validate) *routes.Routes {

	return &routes.Routes{}
}
