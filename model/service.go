package model

//顾客购买
//查询商品是否存在;查询商品库存是否足够;记录订单;减少库存
func BuyProduct(consumer Consumer)bool{

	if !SelectProductId(consumer.ProductId){
		return false
	}
	if SelectProductNum(consumer.ProductId)<1{
		return false
	}
	db.Begin()
	record:=Record{
		UserId:    consumer.UserId,
		ProductId: consumer.ProductId,
		StoreId:   consumer.StoreId,
		Num:       consumer.Num,
		BuyTime:   consumer.BuyTime,
	}
	if !CreateRecord(record) {
		return false
	}
	if !DecUpdateProductNum(consumer.ProductId,consumer.Num){
		return false
	}
	db.Commit()
	return true
}



