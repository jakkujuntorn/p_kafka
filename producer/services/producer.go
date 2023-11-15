package services

import (
	"encoding/json"
	"events"
	"reflect"

	"github.com/Shopify/sarama"
)

type I_EventProducer interface {
	// ใช้ interface{} เปล่า แต่มีชื่อเพราะ จะทำให้มันดูง่าย
	Producer(event events.Event) error // events.Event เป็น interface เปล่า

	// ใช้แบบนี้ก็ได้
	// Producer(event interface{}) error // events.Event เป็น interface เปล่า
}

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) I_EventProducer {
	return eventProducer{producer}
}

// Producer implements EventProducere
// ที่ Func Producer รับ interface{} เปล่า เพราะ Struct หลายแบบที่ส่งเข้ามา
// อาจใช้ generic ลองเอาไปทำดู
func (obj eventProducer) Producer(event events.Event) error {
// func (obj eventProducer) Producer(event interface{}) error {

	// check topic หรือเช็ค ชื่อของ struct
	topic := reflect.TypeOf(event).Name()

	// event to []byte
	//  ทำไมต้อง marshal ด้วย ***
	// เพราะ ข้อูลเป็น json เลยต้องแปลงเป็น string
	value, err := json.Marshal(event)
	if err != nil {
		return err
	}



	// *********** ปั้น message ***********
	msg := sarama.ProducerMessage{
		Topic: topic,
		// sarama.ByteEncoder สำหรับ ข้อมูล struct หรือซับซ้อน
		Value: sarama.ByteEncoder(value),

		// sarama.StringEncode สำหรับ string ทั่วไป
		// Value: sarama.StringEncoder(value), // chat gpt
		// Value: sarama.StringEncoder("SS"), // ถ้าเป็น string ส่งเข้าไปได้เลย
	}

	//************* ส่ง message ************
	// ได้รับ patition, offset, error
	_, _, err = obj.producer.SendMessage(&msg)
	if err != nil {
		return err
	}

	return nil
}
