package model

import (
	"github.com/jinzhu/gorm"
	"summer/rabbitmq/errno"
	"time"
)
type Product struct{
	ID uint
	ProductId uint
	Name string
}
type Stock struct {
	ID uint
	StoreId uint
	ProductId uint
	ProductNum uint
}

type Record struct{
	ID uint
	UserId uint
	ProductId uint
	StoreId uint
	Num uint
	BuyTime time.Time
}


type Consumer struct{
	UserId uint
	ProductId uint
	StoreId uint
	Num uint
	BuyTime time.Time
}

type Storer struct{

}

//查询商品id
func SelectProductId(productId uint)bool{
	if err:=db.Table("product").Where("product_id = ?",productId).First(&Product{}).Error;err!=nil{
		errno.FailOnError(err,"the product is not existed")
		return false
	}
	return true
}



//查询商品存储
func SelectProductNum(productId uint)(num uint){
	var stock Stock
    if err:=db.Table("stock").Where("product_id = ?",productId).First(&stock).Error;err!=nil{
    	errno.FailOnError(err,"select the stock false")
    	return
	}
	return stock.ProductNum
}

//减少订单
func DecUpdateProductNum(productId uint,num uint)bool{
	if err:=db.Table("stock").Where("product_id = ?",productId).Update("product_num",gorm.Expr("product_num - ?",num)).Error;err!=nil{
		errno.FailOnError(err,"declined false")
		return false
	}
	return true
}

//插入订单
func CreateRecord(record Record)bool{
	if err:=db.Table("record").Create(&record).Error;err!=nil{
		return false
	}
	return true
}

//更新库存
func UpdateProductNum(stock Stock)bool{
	if err:=db.Table("stock").Update("product_num",stock.ProductNum).Where("product_id",stock.ProductId).Error;err!=nil{
		errno.FailOnError(err,"update the stock false")
		return false
	}
	return true
}

//插入商品
//插入库存
func InsertProduct(product Product,stock Stock)bool{
	db.Begin()
	if err:=db.Table("product").Create(&product).Error;err!=nil{
		errno.FailOnError(err,"insert the product false")
		return false
	}
	if err:=db.Table("stock").Create(&stock).Error;err!=nil{
		errno.FailOnError(err,"insert the product false")
		return false
	}
	db.Commit()
	return true
}
