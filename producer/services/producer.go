package services

import (
	"encoding/json"
	"events"
	"log"
	"reflect"

	"github.com/Shopify/sarama"
)

type ProducerHandler interface {
	Handle(event events.Event) error
}

type producerHandler struct {
	producer sarama.SyncProducer
}

func NewProducerHandler(producer sarama.SyncProducer) ProducerHandler {
	return &producerHandler{producer}
}

func (obj *producerHandler) Handle(event events.Event) error {

	topic := reflect.TypeOf(event).Name()

	value, err := json.Marshal(event)

	if err != nil {
		log.Printf("Error: %v \n", err)
		return err
	}

	message := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	p, o, err := obj.producer.SendMessage(&message)

	if err != nil {
		log.Printf("Error: %v \n", err)
		return err
	}

	log.Printf("Message sent -> partition=%v, offset=%v \n", p, o)

	return nil
}
