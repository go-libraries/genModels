package models

type UserWallet struct {
	Id         uint    `gorm:"column:id;type:int(11) unsigned" json:"id"`
	UserId     uint64  `gorm:"column:user_id;type:bigint(20) unsigned" json:"user_id"`
	BindId     uint64  `gorm:"column:bind_id;type:bigint(20) unsigned;default:'0'" json:"bind_id"` // 绑定的外部id(活动id、充值流水)
	Worth      float64 `gorm:"column:worth;type:decimal(10,2) unsigned;default:'0.00'" json:"worth"`
	Residue    float64 `gorm:"column:residue;type:decimal(10,2) unsigned;default:'0.00'" json:"residue"`
	Type       uint8   `gorm:"column:type;type:tinyint(1) unsigned;default:'1'" json:"type"`         // 1 充值币 6 赠币,预留2 3 4 5为其他类型
	Platform   uint8   `gorm:"column:platform;type:tinyint(1) unsigned;default:'1'" json:"platform"` // 1 安卓 2 ios 3 web
	Status     uint8   `gorm:"column:status;type:tinyint(1) unsigned;default:'1'" json:"status"`     // 1 正常 2 失效 3 过期
	CreateTime string  `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime string  `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"update_time"`
	ExpireTime string  `gorm:"column:expire_time;type:datetime;default:CURRENT_TIMESTAMP" json:"expire_time"` // 果冻才有过期时间
}

//get real primary key name
func (userWallet *UserWallet) GetKey() string {
	return "id"
}

//get primary key in model
func (userWallet *UserWallet) GetKeyProperty() uint {
	return userWallet.Id
}

//set primary key
func (userWallet *UserWallet) SetKeyProperty(id uint) {
	userWallet.Id = id
}

//get real table name
func (userWallet *UserWallet) TableName() string {
	return "user_wallet"
}

func (userWallet *UserWallet) GetById(id string) {
	Orm.Model(userWallet).First(userWallet, userWallet.GetKey()+" = '"+id+"'")
}

func (userWallet *UserWallet) GetOne(condition ...interface{}) (err []error) {
	err = Orm.Model(userWallet).First(userWallet, condition).GetErrors()
	if len(err) > 0 {
		return err
	}
	return
}

func (userWallet *UserWallet) GetList(page, limit int64, condition interface{}) (list []*UserWallet) {
	err := Orm.Model(userWallet).Limit(limit).Offset((page-1)*limit).Find(list, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}

func (userWallet *UserWallet) Create() []error {
	return Orm.Model(userWallet).Create(userWallet).GetErrors()
}

func (userWallet *UserWallet) Update(update UserWallet) []error {
	return Orm.Model(userWallet).UpdateColumns(info).GetErrors()
}

func (userWallet *UserWallet) Delete() {
	Orm.Model(userWallet).Delete(userWallet)
}
