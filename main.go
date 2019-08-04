package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/streadway/amqp"
	"summer/rabbitmq/controller"
	"summer/rabbitmq/model"
	"summer/rabbitmq/producer"
)


func main() {
	var conn *amqp.Connection
	var chanel *amqp.Channel
	model.ModelInit()
	conn,chanel=producer.ProduceInit()
    controller.ControllerInit(chanel)
	defer producer.ProducerClose(chanel,conn)
	defer model.ModelClose()
}
