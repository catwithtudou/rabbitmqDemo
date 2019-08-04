package consumer

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"summer/rabbitmq/errno"
	"summer/rabbitmq/model"
)

func Consume(c *gin.Context) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	errno.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	errno.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(     // 消息队列
		"buy", // name
		true,   // durable 队列持久化
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	errno.FailOnError(err, "Failed to declare a queue")


	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack  自动ACK
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	errno.FailOnError(err, "Failed to register a consumer")

	//forever := make(chan bool)   // 创建bool型的channel
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			c.String(http.StatusOK,string(d.Body))
			var consumer model.Consumer
			err=json.Unmarshal(d.Body,&consumer)
			errno.FailOnError(err,"反序列化")
			model.BuyProduct(consumer)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-forever   // 从forever信道中取数据，必须要有数据流进来才可以，不然在此阻塞

}

