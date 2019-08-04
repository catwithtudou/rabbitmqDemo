package model

import (
	"github.com/jinzhu/gorm"
	"summer/rabbitmq/errno"
)

var  db *gorm.DB

func ModelInit(){
	var err error
	db,err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/summer?parseTime=true&charset=utf8&loc=Local")
	errno.FailOnError(err,"Can't open the database")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
}

func ModelClose(){
	db.DB().Close()
}

func DatabaseCreate(){
	db.Table("product").CreateTable(&Product{})
	db.Table("stock").CreateTable(&Stock{})
	db.Table("record").CreateTable(&Record{})
}