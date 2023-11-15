package main

import (
	"context"
	"events" // events ที่อ้างมาจากโปรเจค events

	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"

	"consumer/repository"
	"consumer/service"
)

// func นี้จะ run ก่อน main
func init() {
	viper.SetConfigName("config")                          // ชื่อไฟล์
	viper.SetConfigType("yaml")                            // type
	viper.AddConfigPath(".")                               // อ้างการเข้าถึงข้อมูลด้วย . เช่น kafak.server
	viper.AutomaticEnv()                                   // ใช้ใน Env ด้วย ถ้าเจอใน Env ก่อนจะเอาใน Env มาใช้                                  // เจอค่าใน config ก่อนจะเอาค่ามาใช้ ถ้าไม่เจอจะเอาใน yaml มาใช้
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // ใน shell ใช้ . ไม่ได้

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// ใช้ gorm
func initDataBase() {
	//  set up DB
	//viper ดึงค่า  config ต่างๆ
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)
	_ = dsn

	// dial:= mysql.Open(dsn)
	// db,err:= gorm.Open(dial, &gorm.Config{
	// Logger: logger.Default.LogMode(logger.Silent)
	// })
	// if err != nil {
	// 	panic(err)
	// }

}

func main() {
	top1 := events.OpenAccount_Event{}
	top2 := events.Topics
	_ = top1
	_ = top2

	// group
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}

	// close sarama
	// defer consumer.Close()
	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	// Set up DB
	initDataBase()

	// vdo code bangkok ใช้ DB ด้วย
	accountRepo := repository.NewAccountRepo()

	// 2 อันนี้ อยู่ layer เดียวกัน
	// EventHandler ทำหน้าที่คล้าย layer Service
	accountEventService := service.NewAccountEventHandler(accountRepo)
	// NewConsumerHandler ทำหน้าที่ Handler ใน layer Controller
	// ทำหน้าที่ เป็น sarama.ConsumerGroupHandler
	accountCusumerControll := service.NewConsumerHandler(accountEventService) //accountCusumerHandler คือ ConsumerGroupHandler

	// events.Topic เอามาจาก events มี  Topic ที่ เซตไว้
	// for _, value := range events.Topics {
	// 	fmt.Println(value) // OpenAccount_Event | DepositFund_Event | WithdrawFund_Event | CloseAccount_Event
	// }

	// *************** check error ********
	go func() {
		for err := range consumer.Errors() {
			fmt.Println("Error:", err)
		}
	}()

	fmt.Println("Account Consumer Satrt .")

	//  loop เพื่อที่จะ consumer ตลอด (รอรับ message ตลอด)
	for {
		// consumer น่าจะ ติดตาม topic ตรงนี้ เช็คว่ามี topic อะไรบ้าง ***
		// events.Topics มาจาก reflect.TypeOf(OpenAccount_Event{}).Name() - Topic ที่เราสร้างเตรียมไว้
		consumer.Consume(context.Background(), events.Topics, accountCusumerControll)
	}

	// stram chat gpt รึ ป่าว
	// 	server := []string{"localhost:9092"}
	// consumerStream, err := sarama.NewConsumer(server, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// defer consumer.Close()

	// partitionCon,_:=consumerStream.Partitions("my-topic")

	// for msg:= range partitionCon.Message{
	// 	fmt.Println(msg)
	// }

}

// func main() {

// 	server := []string{"localhost:9092"}
// 	consumer, err := sarama.NewConsumer(server, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer consumer.Close()

//ConsumePartition(topic,padition, เลือก message)
// sarama.OffsetNewest เลือก ว่าจะเอา message ล่าสุด
// 	partitionConsumer, err := consumer.ConsumePartition("russy", 0, sarama.OffsetNewest)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer partitionConsumer.Close()

// 	fmt.Println("Consumer Start.")

// 	for {
// 		select {
// case Error
// 		case err := <-partitionConsumer.Errors():
// 			fmt.Println(err)
// Case Message
// 		case msg := <-partitionConsumer.Messages():
// 			fmt.Println(string(msg.Value))
// 		}
// 	}

// }
