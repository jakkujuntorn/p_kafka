package main

import (
	_ "fmt"
	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"producer/controllers"
	"producer/services"
	"strings"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")                               // อ้างการเข้าถึงข้อมูลด้วย . เช่น kafak.server
	viper.AutomaticEnv()                                   // เจอค่าใน config ก่อนจะเอาค่ามาใช้ ถ้าไม่เจอจะเอาใน yaml มาใช้
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // ใน shell ใช้ . ไม่ได้

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// producer ไม่ต้องใช้ DB

func main() {
	// kafka producer

	// Set UP Config ******
	// config  chatGPT
	// config := sarama.NewConfig()
	// config.Producer.RequiredAcks = sarama.WaitForAll
	// config.Producer.Retry.Max = 5
	// config.Producer.Return.Successes = true

	// config.Admin.Timeout = 10

	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}

	defer producer.Close()
	// chat Gpt
	// defer func() {
	// 	if err := producer.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	//*************** port and adaptor ***********
	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountControllse := controllers.NewAccountController(accountService)

	// fiber
	app := fiber.New()

	app.Post("/openaccount", accountControllse.OpenAccount)
	// app.Post("/dipositfund")
	// app.Post("/withdrawfund")
	// app.Post("/closeaccount")

	app.Listen(":8000")
}

// func main() {
// 	server := []string{"localhost:9092"}

// 	producer, err := sarama.NewSyncProducer(server, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer producer.Close()

// เตรียม ปั้น message *****
// 	msg := sarama.ProducerMessage{
// 		Topic: "russy",
// 		Value: sarama.StringEncoder("Hello word jack5"),
// 	}

// ********** ส่ง ********
// 	partition, offset, err := producer.SendMessage(&msg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Partition:%v, offset:%v", partition, offset)

// }
