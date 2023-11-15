package events

import "reflect"

// ได้ค่าอะไร
// ได้ชื่อ type struct
// มัยช้วยได้ค่าที่ถูกต้อง
var Topics = []string{
	reflect.TypeOf(OpenAccount_Event{}).Name(),
	reflect.TypeOf(DepositFund_Event{}).Name(),
	reflect.TypeOf(WithdrawFund_Event{}).Name(),
	reflect.TypeOf(CloseAccount_Event{}).Name(),
}

type Event interface {}

type OpenAccount_Event struct {
	ID             string
	AccountHolder  string
	AccountTYpe    int
	OpeningBalance float64
}

type DepositFund_Event struct {
	ID     string
	Amount float64
}

type WithdrawFund_Event struct {
	ID     string
	Amount float64
}

type CloseAccount_Event struct {
	ID string
}
