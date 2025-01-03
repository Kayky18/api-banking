package database

import (
	"errors"
	"kayky18/api-banking/internal/entity"

	"gorm.io/gorm"
)

type Transaction struct {
	DB *gorm.DB
}

var (
	ErrInvalidTransactionAmount = errors.New("invalid transaction amount")
	ErrInvalidTransactionType   = errors.New("invalid transaction type")
)

func NewTransactionDB(db *gorm.DB) *Transaction {
	return &Transaction{DB: db}
}

func (t *Transaction) CreateTransaction(userid int, description string, amount float64, transactionType string) (*entity.Transaction, error) {

	if amount <= 0 {
		return nil, ErrInvalidTransactionAmount
	}
	if transactionType != "income" && transactionType != "expense" {
		return nil, ErrInvalidTransactionType
	}

	if userid == 0 {
		return nil, errors.New("invalid user id")
	}

	transaction := entity.Transaction{
		UserID:      userid,
		Description: description,
		Amount:      amount,
		Type:        transactionType,
	}
	result := t.DB.Model(&entity.Transaction{}).Create(&transaction)

	return &transaction, result.Error
}

func (t *Transaction) GetTransactions(userid int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	result := t.DB.Model(&entity.Transaction{}).Find(&transactions)
	return transactions, result.Error
}

func (t *Transaction) GetTransaction(userid int, id int) (entity.Transaction, error) {
	var transaction entity.Transaction
	result := t.DB.Model(&entity.Transaction{}).Where("id = ?", id).Where("user_id = ?", userid).First(&transaction)

	return transaction, result.Error
}

func (t *Transaction) UpdateTransaction(userid int, id int, description string, amount float64, transactionType string) error {
	if amount <= 0 {
		return ErrInvalidTransactionAmount
	}
	if transactionType != "income" && transactionType != "expense" {
		return ErrInvalidTransactionType
	}

	transaction, err := t.GetTransaction(userid, id)
	if err != nil {
		return err
	}

	transaction.Description = description
	transaction.Amount = amount
	transaction.Type = transactionType

	result := t.DB.Model(&entity.Transaction{}).Where("id = ?", id).Updates(&transaction)
	return result.Error
}

func (t *Transaction) DeleteTransaction(userid int, id int) error {
	result := t.DB.Model(&entity.Transaction{}).Where("id = ?", id).Delete(&entity.Transaction{})

	return result.Error
}

func (t *Transaction) GetByType(userid int, transactionType string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	result := t.DB.Model(&entity.Transaction{}).Where("type = ?", transactionType).Find(&transactions)
	return transactions, result.Error
}

func (t *Transaction) GetTotalBalance(userid int) (float64, error) {
	var total float64
	result := t.DB.Model(&entity.Transaction{}).Select("sum(amount)").Row().Scan(&total)
	return total, result
}
