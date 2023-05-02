package repositories

import (
	"database/sql"

	"gorm.io/gorm"
)

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() ([]BankAccount, error)
	FindById(id string) (BankAccount, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	// db.AutoMigrate(BankAccount{})
	return &accountRepository{db}
}

func (obj *accountRepository) Save(bankAccount BankAccount) error {
	return obj.db.Save(bankAccount).Error
}

func (obj *accountRepository) Delete(id string) error {
	return obj.db.Where("id=@id", sql.Named("id", id)).Delete(BankAccount{}).Error
}

func (obj *accountRepository) FindAll() ([]BankAccount, error) {
	var bankAccounts []BankAccount

	tx := obj.db.Find(&bankAccounts)
	return bankAccounts, tx.Error
}

func (obj *accountRepository) FindById(id string) (BankAccount, error) {
	var bankAccount BankAccount

	tx := obj.db.Where("id=@id", sql.Named("id", id)).First(&bankAccount)

	return bankAccount, tx.Error
}
