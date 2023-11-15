package repository

import (
	"errors"
	"fmt"
)

// models
type Bank_Aaccount struct {
	ID             string
	AccountHolder  string
	AccountTYpe    int
	Balance float64
}

type IAccountRepository interface {
	// CRUD
	SaveAccount(Bank_Aaccount) (*Bank_Aaccount, error)
	GetAccountAll() ([]Bank_Aaccount, error)
	GetAccountById(id string) (Bank_Aaccount, error)
	UpdateAccount(id string, bankaccount Bank_Aaccount) error
	DeleteAccount(id string) error
}

// type  `AccountRepository` is a struct that is not doing anything in this code snippet. It is not
// being used or referenced anywhere.
type accountRepository struct {
	// DB 
}

// ใน code bangkok ใช้ Gorm
// ใส พารามิเตอร์เป็น DB ด้วย
func NewAccountRepo() IAccountRepository {
	// autoMigrate()
	return &accountRepository{}
}

// CreateAccount implements IAccount
func (*accountRepository) SaveAccount(bankAccount Bank_Aaccount) (*Bank_Aaccount, error) {
	fmt.Println("Create Accont DB ")
	fmt.Println(bankAccount)
	return &bankAccount, nil
}

// GetAccountAll implements IAccount
func (*accountRepository) GetAccountAll() (bankAccount []Bank_Aaccount, err error) {
	bank_Ac := []Bank_Aaccount{}
	err = errors.New("Err Na ja")
	return bank_Ac, err
}

// GetAccountById implements IAccount
func (*accountRepository) GetAccountById(id string) (Bank_Aaccount, error) {
	panic("unimplemented")
}

// UpdateAccount implements IAccount
func (*accountRepository) UpdateAccount(id string, bankaccount Bank_Aaccount) error {
	panic("unimplemented")
}

// DeleteAccount implements IAccount
func (*accountRepository) DeleteAccount(ids string) error {
	dd := ids
	_ = dd
	// db.Where("id=?",id).Delete(&BankAccount) // id ที่เป็น string ต้อวง where ก่อน ไม่งั้น error  ยกเว้น id เป็น int
	return nil
}

// func NewAccoute_Repo() IAccountRepository {
// 	return &accountRepository{}
// }
