package models

type OrderCharge struct {
	Id              uint64  `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	OrderId         uint64  `gorm:"column:order_id;type:bigint(20) unsigned" json:"order_id"`
	UserId          uint64  `gorm:"column:user_id;type:bigint(20) unsigned" json:"user_id"`
	PlatformOrderId string  `gorm:"column:platform_order_id;type:varchar(255)" json:"platform_order_id"`                            // 第三方订单绑定id
	ServeFee        float64 `gorm:"column:serve_fee;type:decimal(10,2);default:'0.00'" json:"serve_fee"`                            // 服务费
	ServeRate       float64 `gorm:"column:serve_rate;type:decimal(10,2) unsigned zerofill;default:'00000000.00'" json:"serve_rate"` // 抽成比例
}

//get real primary key name
func (orderCharge *OrderCharge) GetKey() string {
	return "id"
}

//get primary key in model
func (orderCharge *OrderCharge) GetKeyProperty() uint64 {
	return orderCharge.Id
}

//set primary key
func (orderCharge *OrderCharge) SetKeyProperty(id uint64) {
	orderCharge.Id = id
}

//get real table name
func (orderCharge *OrderCharge) TableName() string {
	return "order_charge"
}

func (orderCharge *OrderCharge) GetById(id string) {
	Orm.Model(orderCharge).First(orderCharge, orderCharge.GetKey()+" = '"+id+"'")
}

func (orderCharge *OrderCharge) GetOne(condition ...interface{}) (err []error) {
	err = Orm.Model(orderCharge).First(orderCharge, condition).GetErrors()
	if len(err) > 0 {
		return err
	}
	return
}

func (orderCharge *OrderCharge) GetList(page, limit int64, condition interface{}) (list []*OrderCharge) {
	err := Orm.Model(orderCharge).Limit(limit).Offset((page-1)*limit).Find(list, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}

func (orderCharge *OrderCharge) Create() []error {
	return Orm.Model(orderCharge).Create(orderCharge).GetErrors()
}

func (orderCharge *OrderCharge) Update(update OrderCharge) []error {
	return Orm.Model(orderCharge).UpdateColumns(info).GetErrors()
}

func (orderCharge *OrderCharge) Delete() {
	Orm.Model(orderCharge).Delete(orderCharge)
}
