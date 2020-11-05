package models

type OrderGoods struct {
	Id        uint64  `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	OrderId   uint64  `gorm:"column:order_id;type:bigint(20) unsigned" json:"order_id"`
	UserId    uint64  `gorm:"column:user_id;type:bigint(20) unsigned" json:"user_id"`
	GoodsId   uint64  `gorm:"column:goods_id;type:bigint(20) unsigned" json:"goods_id"`
	GoodsName string  `gorm:"column:goods_name;type:varchar(255)" json:"goods_name"`
	Cost      float64 `gorm:"column:cost;type:decimal(10,2)" json:"cost"`
	Price     float64 `gorm:"column:price;type:decimal(10,2)" json:"price"`
	GoodsType uint8   `gorm:"column:goods_type;type:tinyint(1) unsigned;default:'1'" json:"goods_type"` // 商品类型  1 音频  2 专辑
	Number    uint    `gorm:"column:number;type:int(10) unsigned;default:'1'" json:"number"`            // 购买数量
}

//get real primary key name
func (orderGoods *OrderGoods) GetKey() string {
	return "id"
}

//get primary key in model
func (orderGoods *OrderGoods) GetKeyProperty() uint64 {
	return orderGoods.Id
}

//set primary key
func (orderGoods *OrderGoods) SetKeyProperty(id uint64) {
	orderGoods.Id = id
}

//get real table name
func (orderGoods *OrderGoods) TableName() string {
	return "order_goods"
}

func (orderGoods *OrderGoods) GetById(id string) {
	Orm.Model(orderGoods).First(orderGoods, orderGoods.GetKey()+" = '"+id+"'")
}

func (orderGoods *OrderGoods) GetOne(condition ...interface{}) (err []error) {
	err = Orm.Model(orderGoods).First(orderGoods, condition).GetErrors()
	if len(err) > 0 {
		return err
	}
	return
}

func (orderGoods *OrderGoods) GetList(page, limit int64, condition interface{}) (list []*OrderGoods) {
	err := Orm.Model(orderGoods).Limit(limit).Offset((page-1)*limit).Find(list, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}

func (orderGoods *OrderGoods) Create() []error {
	return Orm.Model(orderGoods).Create(orderGoods).GetErrors()
}

func (orderGoods *OrderGoods) Update(update OrderGoods) []error {
	return Orm.Model(orderGoods).UpdateColumns(info).GetErrors()
}

func (orderGoods *OrderGoods) Delete() {
	Orm.Model(orderGoods).Delete(orderGoods)
}
