package models

import (
	"beego.su77.cn/utils"
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

//去拼去相信 敢闯创不凡
type Users struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Mobile string `json:"mobile"`
	Password string `json:"-"`
	Gender int `json:"gender"`
	City string `json:"city"`
	Ip string `json:"ip"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
	UserProfile []*UserProfile `json:"user_profile" orm:"reverse(many)"`
}

func add(a, b *int) (c int, err error) {
	if (*a < 0 || *b < 0) {
		err = errors.New("只支持非负整数相加")
		return
	}
	*a *= 2
	*b *= 3
	c = *a + *b
	return
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