package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	controllers "manjo-test/controllers/transaction"
	"manjo-test/domain/models"
	"manjo-test/middlewares"
	repositories "manjo-test/repositories/transaction"
	services "manjo-test/services/transaction"
	"net/http"
	"os"
)

func main() {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	appPort := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Successfully connected to database")

	Migrations(db)

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		c.Set("Access-Control-Allow-Headers", "Content-Type, x-api-key, x-request-at")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Status":  "success",
			"Message": "QR Generator API",
		},
		)
	})

	group := app.Group("/api/v1")
	group.Use(middlewares.Authenticate())
	group.Post("/qr/generate", transactionController.Create)
	group.Post("/qr/payment", transactionController.Update)

	if err = app.Listen(":" + appPort); err != nil {
		log.Fatal(err)
	}
}

func Migrations(db *gorm.DB) {
	var err error
	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")
}
