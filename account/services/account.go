package services

import "account/repositories"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountService interface {
	GetAll() ([]BankAccount, error)
}

type accountService struct {
	accountRepository repositories.AccountRepository
}

func NewAccountService(accountRepository repositories.AccountRepository) AccountService {
	return &accountService{accountRepository}
}

func (obj *accountService) GetAll() ([]BankAccount, error) {

	var bankAccounts []BankAccount

	results, err := obj.accountRepository.GetAll()

	if err != nil {
		return nil, err
	}

	for _, value := range results {
		bankAccounts = append(bankAccounts, BankAccount{
			ID:            value.ID,
			AccountHolder: value.AccountHolder,
			AccountType:   value.AccountType,
			Balance:       value.Balance,
		})
	}

	return bankAccounts, nil
}
