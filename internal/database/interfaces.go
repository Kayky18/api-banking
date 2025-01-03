package database

import "kayky18/api-banking/internal/entity"

type UserInterface interface {
	CreateUser(name string, email string, password string, role string) (*entity.User, error)
	FindUserById(id string) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	ValidatePassword(password string, hash string) bool
}

type TransactionInterface interface {
	CreateTransaction(userid int, description string, amount float64, transactionType string) (*entity.Transaction, error)

	GetTransactions(userid int) ([]entity.Transaction, error)

	GetTransaction(userid int, id int) (entity.Transaction, error)

	GetByType(userid int, transactionType string) ([]entity.Transaction, error)

	GetTotalBalance(userid int) (float64, error)

	UpdateTransaction(userid int, id int, description string, amount float64, transactionType string) error

	DeleteTransaction(userid int, id int) error
}

type AmountInterface interface {
	GetAmount(id string) (float64, error)
	UpdateAmount(id string, amount float64) error
	CreateAmount(id string, amount float64) error
}
