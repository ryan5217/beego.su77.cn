package models

import (
	"beego.su77.cn/utils"
	"github.com/astaxie/beego/orm"
	"time"
)

//去拼去相信 敢闯创不凡
type Users struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	Gender int `json:"gender"`
	City string `json:"city"`
	Ip string `json:"ip"`
	CreatedAt time.Time `json:"created_at",orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at",orm:"auto_now;type(datetime)"`
	UserProfile []*UserProfile `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Users))
}

func CheckPassword(mobile, password string) (*Users, bool) {
	ormer := orm.NewOrm()

	ok := false
	userRead := Users{Mobile:mobile}
	_ = ormer.Read(&userRead, "Mobile")

	password = utils.PasswordMD5(userRead.Mobile, password)
	if userRead.Password ==  password{
		ok = true
	}

	//var user User
	//err := o.QueryTable("user").Filter("name", "slene").One(&user)
	//if err == orm.ErrMultiRows {
	//	// 多条的时候报错
	//	fmt.Printf("Returned Multi Rows Not One")
	//}
	//if err == orm.ErrNoRows {
	//	// 没有找到记录
	//	fmt.Printf("Not row found")
	//}

	return &userRead,ok
}