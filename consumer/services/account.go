package services

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepository repositories.AccountRepository
}

func NewAccountEventHandler(accountRepository repositories.AccountRepository) EventHandler {
	return &accountEventHandler{accountRepository}
}

func (obj *accountEventHandler) Handle(topic string, eventBytes []byte) {

	log.Printf("Incoming Message: %#v \n", string(eventBytes))

	switch topic {
	// OpenAccountEvent
	case reflect.TypeOf(events.OpenAccountEvent{}).Name():

		event := events.OpenAccountEvent{}

		err := json.Unmarshal(eventBytes, &event)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		bankAccount := repositories.BankAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBalance,
		}

		err = obj.accountRepository.Save(bankAccount)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		log.Printf("%#v \n", event)

	case reflect.TypeOf(events.DepositFundEvent{}).Name():

		event := events.DepositFundEvent{}

		err := json.Unmarshal(eventBytes, &event)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		bankAccount, err := obj.accountRepository.FindById(event.ID)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		bankAccount.Balance += event.Amount

		err = obj.accountRepository.Save(bankAccount)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		log.Printf("%#v \n", event)

	case reflect.TypeOf(events.WithdrawFundEvent{}).Name():

		event := events.DepositFundEvent{}

		err := json.Unmarshal(eventBytes, &event)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		bankAccount, err := obj.accountRepository.FindById(event.ID)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		bankAccount.Balance -= event.Amount

		err = obj.accountRepository.Save(bankAccount)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		log.Printf("%#v \n", event)

	case reflect.TypeOf(events.CloseAccountEvent{}).Name():

		event := events.CloseAccountEvent{}

		err := json.Unmarshal(eventBytes, &event)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		err = obj.accountRepository.Delete(event.ID)

		if err != nil {
			log.Printf("Error: %v \n", err)
			return
		}

		log.Printf("%#v \n", event)

	default:
		log.Panicln("No event handler")
	}
}
