package admin

import (
	"beego.su77.cn/models"
	"beego.su77.cn/utils"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
)

const SuccessCode int = 0
const Message string = "登录成功"

type LoginController struct {
	BaseController
}

type Data struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func (this *LoginController) Login() {
	this.Data["name"] = "登录"

	this.TplName = "admin/login.html"
}

//执行登录操作
func (this *LoginController) DoLogin() {

	data := make(map[string] interface{})
	data["mobile"] = this.GetString("mobile")
	data["password"] = this.GetString("password")
	users, ok := models.CheckPassword(this.GetString("mobile"), this.GetString("password"))

	data["user"] = users
	data["ok"] = ok

	if  ok {
		this.Data["json"] = Data{SuccessCode, Message, data}
	} else {

		this.Data["json"] = Data{100001, "登录失败", data}
	}

	this.ServeJSON()
}


//注册
func (this *LoginController) Register() {
	this.Data["title"] = "注册"

	this.TplName = "admin/register.html"
}

func (this *LoginController) Curl() {

	req := httplib.Post("http://t-userapi.acc.cn/api/v1/regions")
	req.Header("Accept", "application/json")
	req.Header("Content-Type", "application/json")

	req.Param("parent_id", "110000")

	//base12, _ := req.Response()

	req.Debug(true)

	data := make(map[string] interface{})
	data["mobile"] = "123"

	var str interface{}
	err := req.ToJSON(&str)

	//fmt.Printf("json",str["code"])

	data["str"] = str

	if err != nil {
		this.Data["json"] = Data{1, "错误", data}
	} else {
		this.Data["json"] = Data{SuccessCode, Message, data}

	}

	this.ServeJSON()
}

//执行注册
func (this *LoginController) DoRegister() {
	var mobile, password string
	mobile = this.GetString("mobile")
	password = this.GetString("password")
	//注册
	ormer := orm.NewOrm()
	//user := new(models.Users)
	user := models.Users{}
	user.Mobile = mobile
	user.Password = utils.PasswordMD5(mobile, password)
	user.Name = "ryan"
	user.Gender = 1

	created, id, err := ormer.ReadOrCreate(&user, "Mobile")

	data := make(map[string] interface{})
	data["user"] = user
	data["id"] = id

	//this.Data["json"] = Data{SUCCESS_CODE, MESSAGE, data}

	if err != nil {
		data["err"] = err
		this.Data["json"] = Data{100004, "系统错误", data}

	} else {

		if created {
			this.Data["json"] = Data{SuccessCode, Message, data}
		} else {
			this.Data["json"] = Data{100003, "该手机号码已存在", data}
		}

	}

	this.ServeJSON()
}
