package main

import (
	"fmt"
	"log"

	"github.com/Icarus-xD/AlgoMaster/internal/config"
	"github.com/Icarus-xD/AlgoMaster/internal/database"
	"github.com/Icarus-xD/AlgoMaster/internal/repository"
	"github.com/Icarus-xD/AlgoMaster/internal/service"
	"github.com/Icarus-xD/AlgoMaster/internal/transport/rest"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Config
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// DB
	dsn := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgresDriver, cfg.PostgresUser, cfg.PostgresPassword,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB, cfg.PostgresSSLMode,
	)
	db, err := database.Init(dsn) 
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// REPO's
	userRepo := repository.NewUserRepo(db)
	priceRepo := repository.NewTaskPriceRepo(db)

	// SERVICES
	emailService := service.NewEmailService(cfg.SMTPHost, cfg.SMTPUsername, cfg.SMTPPassword, cfg.SMTPPort)
	emailService.SendEmail("DebtExceeded", "email@email.com", "Debt Exceeded", struct{}{})

	userService := service.NewUserService(userRepo)
	taskService := service.NewTaskService(emailService, userRepo, priceRepo)

	// REST
	app := fiber.New()

	handler := rest.NewHandler(userService, taskService)
	handler.DefineRoutes(app)

	err = app.Listen(cfg.AppPort)
	if err != nil {
		log.Fatal(err)
	}
}