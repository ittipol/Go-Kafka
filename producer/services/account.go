package services

import (
	"events"
	"log"
	"producer/commands"

	"github.com/google/uuid"
)

type AccountService interface {
	OpenAccount(command commands.OpenAccountCommand) (id string, err error)
	DepositFund(command commands.DepositFundCommand) error
	WithdrawFund(command commands.WithdrawFundCommnd) error
	CloseAccount(command commands.CloseAccountCommand) error
}

type accountService struct {
	producerHandler ProducerHandler
}

func NewAccountService(producerHandler ProducerHandler) AccountService {
	return &accountService{producerHandler}
}

func (obj *accountService) OpenAccount(command commands.OpenAccountCommand) (id string, err error) {

	event := events.OpenAccountEvent{
		ID:             uuid.New().String(),
		AccountHolder:  command.AccountHolder,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}

	log.Printf("Producer Event: %#v\n", event)

	err = obj.producerHandler.Handle(event)

	if err != nil {
		log.Printf("Error: %v \n", err)
		return "", err
	}

	return event.ID, nil
}

func (obj *accountService) DepositFund(command commands.DepositFundCommand) error {

	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("Producer Event: %#v\n", event)

	err := obj.producerHandler.Handle(event)

	if err != nil {
		log.Printf("Error: %v \n", err)
		return err
	}

	return nil
}

func (obj *accountService) WithdrawFund(command commands.WithdrawFundCommnd) error {

	event := events.WithdrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("Producer Event: %#v\n", event)

	err := obj.producerHandler.Handle(event)

	if err != nil {
		log.Printf("Error: %v \n", err)
		return err
	}

	return nil
}

func (obj *accountService) CloseAccount(command commands.CloseAccountCommand) error {

	event := events.CloseAccountEvent{
		ID: command.ID,
	}

	log.Printf("Producer Event: %#v\n", event)

	err := obj.producerHandler.Handle(event)

	if err != nil {
		log.Printf("Error: %v \n", err)
		return err
	}

	return nil
}
