package main

import (
	"account/controllers"
	"account/repositories"
	"account/services"
	"fmt"
	"migration/orm/db"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.EnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

func main() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)

	db := db.GetConnection(dsn, false)

	accountRepository := repositories.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)
	accountController := controllers.NewAccountController(accountService)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	app.Get("health", func(c *fiber.Ctx) error {

		c.Status(fiber.StatusOK)
		c.JSON(fiber.Map{
			"message": "ok",
		})

		return nil
	})

	app.Get("getAccounts", accountController.GetAccounts)

	app.Listen(":5000")

}
