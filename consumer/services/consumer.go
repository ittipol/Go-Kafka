package services

import "github.com/Shopify/sarama"

type consumerGroupHandler struct {
	eventHandler EventHandler
}

func NewConsumerGroupHandler(eventHandler EventHandler) sarama.ConsumerGroupHandler {
	return &consumerGroupHandler{eventHandler}
}

func (obj *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj *consumerGroupHandler) ConsumeClaim(consumerGroupSession sarama.ConsumerGroupSession, consumerGroupClaim sarama.ConsumerGroupClaim) error {

	for message := range consumerGroupClaim.Messages() {
		obj.eventHandler.Handle(message.Topic, message.Value)
		consumerGroupSession.MarkMessage(message, "")
	}

	return nil
}
