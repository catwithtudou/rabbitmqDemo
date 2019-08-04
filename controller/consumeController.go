package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"net/http"
	"strconv"
	"summer/rabbitmq/consumer"
	"summer/rabbitmq/errno"
	"summer/rabbitmq/model"
	"summer/rabbitmq/producer"
	"time"
	"unsafe"
)

func ControllerInit(channel *amqp.Channel){
	route:=gin.Default()
	route.POST("/buy", func(context *gin.Context) {
		UserId,_:=strconv.ParseInt(context.PostForm("user-id"),10,64)
		ProductId,_:=strconv.ParseInt(context.PostForm("product-id"),10,64)
		ShopId,_:=strconv.ParseInt(context.PostForm("store-id"),10,64)
		Num,_:=strconv.ParseInt(context.PostForm("number"),10,64)
		consumer:=model.Consumer{
			UserId:    *(*uint)(unsafe.Pointer(&UserId)),
			ProductId: *(*uint)(unsafe.Pointer(&ProductId)),
			StoreId:   *(*uint)(unsafe.Pointer(&ShopId)),
			Num:       *(*uint)(unsafe.Pointer(&Num)),
			BuyTime:   time.Now(),
		}
		consume,err:=json.Marshal(consumer)
		errno.FailOnError(err,"序列化失败")
		producer.ProducePublish(channel,consume)
		context.String(http.StatusOK,string(consume))
	})
	route.GET("/get", func(context *gin.Context) {
		consumer.Consume(context)
	})
	route.Run(":9025")
}
