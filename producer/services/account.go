package services

import (
	"events"
	"github.com/google/uuid"
	"producer/commands"
)

// ทำแบบรับที่เดียวแล้วใช้ switch แยกเอาก็ได้
// คล้าย service ของ comsumer
// มี Func เดียว แต่แยก topic ใน Func
type I_AccountService interface {
	// struct command สร้างเอาไว้ให้  producer ใช้เท่านั้น ***
	OpenAccount(command commands.OpenAccount_Command) (id string, err error)
	DeposiFund(command commands.DepositFund_Command) error
	WithdrawFun(command commands.WithdrawFund_Command) error
	CloseAccount(command commands.CloseAccount_Command) error
}

type accountService struct {
	eventProducer I_EventProducer
}

func NewAccountService(eventProducer I_EventProducer) I_AccountService {
	return &accountService{
		eventProducer: eventProducer,
	}
}

// CloseAccount implements AccountService
func (obj *accountService) CloseAccount(command commands.CloseAccount_Command) error {

	// validator

	// ปั้น events เตียมส่งเข้าไปหา porducer
	event := events.CloseAccount_Event{
		ID: command.ID,
	}

	return obj.eventProducer.Producer(event)
}

// DeposiFund implements AccountService
func (obj *accountService) DeposiFund(command commands.DepositFund_Command) error {

	// validator

	// ปั้น events เตียมส่งเข้าไปหา porducer
	event := events.DepositFund_Event{
		ID:     command.ID,
		Amount: command.Amount,
	}

	// return แบบนี้ก็ได้ เพราะ obj.eventProducer.Producer(event) มัน return error อยู่ แล้ว
	return obj.eventProducer.Producer(event)
}

// OpenAccount implements AccountService
func (obj *accountService) OpenAccount(command commands.OpenAccount_Command) (id string, err error) {

	// validater command
	// command.AccountHolder ==""
	// command.AccountTYpe == 0
	// command.OpeningBalance == 0

	// ปั้น events เตียมส่งเข้าไปหา porducer
	// event คือ struct อ่ะและ
	// ต้องใช้ events จาก evets เท่านั้น เพราะ ตั้งค่า paramiter ไว้
	event := events.OpenAccount_Event{
		ID:             uuid.NewString(),
		AccountHolder:  command.AccountHolder,
		AccountTYpe:    command.AccountTYpe,
		OpeningBalance: command.OpeningBalance,
	}

	// err ประกาศไว้ด้านบนแล้ว เลยไม่ต้องใช้ :=
	// ส่ง event
	err = obj.eventProducer.Producer(event)

	return event.ID, err
}

// WithdrawFun implements AccountService
func (obj *accountService) WithdrawFun(command commands.WithdrawFund_Command) error {

	// validator

	// ปั้น events เตียมส่งเข้าไปหา porducer
	event := events.WithdrawFund_Event{
		ID:     command.ID,
		Amount: command.Amount,
	}

	return obj.eventProducer.Producer(event)
}
