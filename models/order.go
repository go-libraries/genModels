package models

type Order struct {
	Id         uint64  `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	OrderId    uint64  `gorm:"column:order_id;type:bigint(20) unsigned" json:"order_id"`
	UserId     uint64  `gorm:"column:user_id;type:bigint(20) unsigned" json:"user_id"`
	Name       string  `gorm:"column:name;type:varchar(255)" json:"name"` // 订单内容的缩写或者主题描述
	Cost       float64 `gorm:"column:cost;type:decimal(10,2) unsigned;default:'0.00'" json:"cost"`
	Price      float64 `gorm:"column:price;type:decimal(10,2) unsigned;default:'0.00'" json:"price"`
	PayType    uint8   `gorm:"column:pay_type;type:tinyint(3) unsigned" json:"pay_type"` // 订单类型：0-未知，1-微信，2-支付宝，3-ios，4-果冻
	Status     uint8   `gorm:"column:status;type:tinyint(1) unsigned" json:"status"`     // 订单状态：0-待支付，1-支付成功，2-支付失败
	CreateTime string  `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime string  `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"update_time"`
}

//get real primary key name
func (order *Order) GetKey() string {
	return "id"
}

//get primary key in model
func (order *Order) GetKeyProperty() uint64 {
	return order.Id
}

//set primary key
func (order *Order) SetKeyProperty(id uint64) {
	order.Id = id
}

//get real table name
func (order *Order) TableName() string {
	return "order"
}

func (order *Order) GetById(id string) {
	Orm.Model(order).First(order, order.GetKey()+" = '"+id+"'")
}

func (order *Order) GetOne(condition ...interface{}) (err []error) {
	err = Orm.Model(order).First(order, condition).GetErrors()
	if len(err) > 0 {
		return err
	}
	return
}

func (order *Order) GetList(page, limit int64, condition interface{}) (list []*Order) {
	err := Orm.Model(order).Limit(limit).Offset((page-1)*limit).Find(list, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}

func (order *Order) Create() []error {
	return Orm.Model(order).Create(order).GetErrors()
}

func (order *Order) Update(update Order) []error {
	return Orm.Model(order).UpdateColumns(info).GetErrors()
}

func (order *Order) Delete() {
	Orm.Model(order).Delete(order)
}
