package service

import (
	"consumer/repository"
	"encoding/json"
	"events"
	"fmt"
	"log"
	"reflect"
)

type I_EventHandler_AccountService interface {
	Handle_Events(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accoutRepo repository.IAccountRepository
}

func NewAccountEventHandler(accoutRepo repository.IAccountRepository) I_EventHandler_AccountService {
	return &accountEventHandler{accoutRepo: accoutRepo}
}

// Handle implements EventHandler
// สร้าง Func เดียว  และใช้ Switch มาแยกเอาว่าจะไป Event ไหนด้วย reflect.TypeOf
func (ae *accountEventHandler) Handle_Events(topic string, eventBytes []byte) {

	//*********** แยก topic ว่าเข้า เงื่อนไขไหน **********
	switch topic {
	case reflect.TypeOf(events.OpenAccount_Event{}).Name():
		fmt.Println("step 2 Handlet switch ")

		// Event
		event := &events.OpenAccount_Event{}
		// แปลง Event eventBytes ที่รับมา แปลงเป็น events.OpenAccount_Event{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		// validater data ก่อน

		// ปั้นข้อมูลก่อนลง DB
		bankAccount := repository.Bank_Aaccount{
			ID:             event.ID,
			AccountHolder:  event.AccountHolder,
			AccountTYpe:    event.AccountTYpe,
			Balance: event.OpeningBalance,
		}

		// ส่งไป REpo  เข้า DB
		result, errCreate := ae.accoutRepo.SaveAccount(bankAccount)
		if errCreate != nil {
			log.Println(errCreate)
		}
		_ = result

		// fmt.Printf("[%v] %#v",topic,event) // แกะ  prointf ว่า ปริ้นค่าอะไร **
		// topic output ->  [OpenAccount_Event]

		// event output -> &events.OpenAccount_Event
		//{ID:"92730f4a-28da-452f-9cfa-f0e754a3df7b",
		// AccountHolder:"Jack", AccountTYpe:1, OpeningBalance:1500}

	case reflect.TypeOf(events.DepositFund_Event{}).Name():
		fmt.Println("DepositFund_Event")
		event := &events.DepositFund_Event{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}

		// ดึงข้อมูบขึ้นมาก่อน
		bankAccount, err := ae.accoutRepo.GetAccountById(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		_ = bankAccount

		// บวกเงินฝากเข้าไป
		bankAccount.Balance += event.Amount

		// SAave ลง DB

		_, err = ae.accoutRepo.SaveAccount(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}

	case reflect.TypeOf(events.WithdrawFund_Event{}).Name():
	case reflect.TypeOf(events.CloseAccount_Event{}).Name():
	default:
		fmt.Println("No Event Hanlder")
	}
}
