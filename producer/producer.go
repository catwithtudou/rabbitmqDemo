package producer

import (
	"github.com/streadway/amqp"
	"log"
	"summer/rabbitmq/errno"
)

func ProduceInit()(conn *amqp.Connection,ch *amqp.Channel){
	//建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	errno.FailOnError(err, "Failed to connect to RabbitMQ")

	//创建channel
	ch, err = conn.Channel()
	errno.FailOnError(err, "Failed to connect to channel")

	//创建消息队列,queue，并分配默认binding，empty exchange
	_, err = ch.QueueDeclare(
		"buy", //消息队列名字
		true,    //队列持久化
		false,   //自动删除
		false,   //优先
		false,   //不等待
		nil,
	)
	errno.FailOnError(err, "Failed to declare a queue")
	return
}

func ProducePublish(ch  *amqp.Channel,consumer []byte){

	// 发布消息，第一个参数表示路由名称（exchange），""则表示使用默认消息路由
	err:=ch.Publish(
		"",
		"buy",
		false,
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			Body:            consumer,
		})
	errno.FailOnError(err,"发布信息失败")
	log.Println(" [x] Sent %s", string(consumer))
}

func ProducerClose(ch *amqp.Channel,conn *amqp.Connection)(){
	ch.Close()
	conn.Close()
}