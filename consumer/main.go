package main

import (
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"log"
	"migration/orm/db"
	"strings"

	"github.com/Shopify/sarama"
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
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)

	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)

	db := db.GetConnection(dsn, false)

	err = db.AutoMigrate(repositories.BankAccount{})

	if err != nil {
		panic(err)
	}

	accountRepository := repositories.NewAccountRepository(db)
	accountEventHandler := services.NewAccountEventHandler(accountRepository)
	consumerGroupHandler := services.NewConsumerGroupHandler(accountEventHandler)

	log.Printf("Start Consume from topic")
	for {
		consumer.Consume(context.Background(), events.Topics, consumerGroupHandler)
	}

}

// func main() {

// 	addrs := []string{
// 		"localhost:9092",
// 	}

// 	consumer, err := sarama.NewConsumer(addrs, nil)

// 	if err != nil {
// 		panic("cannot connect to server")
// 	}
// 	defer consumer.Close()

// 	partition, err := consumer.ConsumePartition("mtopic", 0, sarama.OffsetNewest)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer partition.Close()

// 	fmt.Println("Start consumer....")
// 	for {
// 		select {
// 		case err := <-partition.Errors():
// 			fmt.Printf("Error: %v \n", err)
// 			// return
// 		case message := <-partition.Messages():
// 			fmt.Printf("Topic: %v \n", message.Topic)
// 			fmt.Printf("Message: %v \n", string(message.Value))
// 		}
// 	}

// }
