package service

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type consumerHandler struct {
	eventHandler I_EventHandler_AccountService
}

// คล้าย layer controller
// confrom ตาม sarama.ConsumerGroupHandler
func NewConsumerHandler(eventHandler I_EventHandler_AccountService) sarama.ConsumerGroupHandler {
	return consumerHandler{
		eventHandler: eventHandler,
	}
}

// Setup implements sarama.ConsumerGroupHandler
func (obj consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup implements sarama.ConsumerGroupHandler
func (obj consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim implements sarama.ConsumerGroupHandler
// ใช้ อันนี้ อันเดียว *****
// ****************** Message จะวิ่งมาที่นี้ ****************
func (obj consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Println("step 1 ConsumeClaim ")

	for msg := range claim.Messages() { // message เข้ามาที่ claims.Message
		// fmt.Println(msg) // output []byte
		// fmt.Println(msg.Topic) // topic
		// fmt.Println(msg.Value) // []byte

		// ส่ง message ไปที่ Handler() / topic กับ value
		obj.eventHandler.Handle_Events(msg.Topic, msg.Value) // msg.Key ก็ได้รึ ป่าว

		// mark message เช็ค message ว่ามีการได้รับแล้ว
		session.MarkMessage(msg, "") // session จัดการเรื่อง message ที่รับไปทำงานแล้ว
	}

	return nil
}
