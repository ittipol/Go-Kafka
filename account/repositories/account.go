package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	GetAll() ([]BankAccount, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db}
}

func (obj *accountRepository) GetAll() ([]BankAccount, error) {

	var bankAccounts []BankAccount

	tx := obj.db.Find(&bankAccounts)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return bankAccounts, nil
}
