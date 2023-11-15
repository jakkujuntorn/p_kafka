package commands

// ที่ไม่ใช้ของ event เพราะ มันแค่คล้ายๆกัน เช่น OpenAccount_Command ไม่ต้องใช้  ID
type OpenAccount_Command struct { // ใช้คล้ายๆ evet แต่ไม่เอา ID เดียวให้ UUID สร้าง
	AccountHolder  string
	AccountTYpe    int
	OpeningBalance float64
}

type DepositFund_Command struct {
	ID     string
	Amount float64
}

type WithdrawFund_Command struct {
	ID     string
	Amount float64
}

type CloseAccount_Command struct {
	ID string
}