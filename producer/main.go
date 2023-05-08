package main

import (
	"fmt"
	"producer/controllers"
	"producer/services"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

func main() {

	addrs := viper.GetStringSlice("kafka.servers")

	producer, err := sarama.NewSyncProducer(addrs, nil)

	if err != nil {
		panic(err)
	}
	defer producer.Close()

	producerHandler := services.NewProducerHandler(producer)
	accountService := services.NewAccountService(producerHandler)
	accountController := controllers.NewAccountController(accountService)

	app := fiber.New()

	app.Get("health", func(c *fiber.Ctx) error {

		c.Status(fiber.StatusOK)
		c.JSON(fiber.Map{
			"message": "ok",
		})

		return nil
	})

	app.Post("openAccount", accountController.OpenAccount)
	app.Post("depositFund", accountController.DepositFund)
	app.Post("withdrawFund", accountController.WithdrawFund)
	app.Post("closeAccount", accountController.CloseAccount)

	app.Listen(":4000")
}

// func main() {

// 	addrs := []string{
// 		"localhost:9092",
// 	}

// 	producer, err := sarama.NewSyncProducer(addrs, nil)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer producer.Close()

// 	message := sarama.ProducerMessage{
// 		Topic: "mtopic",
// 		Value: sarama.StringEncoder("Hello World......"),
// 	}

// 	partition, offset, err := producer.SendMessage(&message)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Partition: %v, Offset: %v \n", partition, offset)
// }
