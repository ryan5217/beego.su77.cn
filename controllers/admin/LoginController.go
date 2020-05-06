package admin

import (
	"beego.su77.cn/models"
	"beego.su77.cn/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
	"github.com/dgrijalva/jwt-go"
)

const SuccessCode int = 0
const Message string = "登录成功"

type LoginController struct {
	BaseController
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
		auth := make(map[string] interface{})
		auth["authorization"] = CreateToken(users.Mobile)
		auth["exp"] = TokenExp()
		data["auth"] =  auth

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
	user := new(models.Users)
	//user := models.Users{}
	user.Mobile = mobile
	user.Password = utils.PasswordMD5(mobile, password)
	user.Name = "ryan"
	user.Gender = 1
	user.Ip = this.Ctx.Request.Header.Get("X-Real-ip")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	//userProfile := models.UserProfile{}
	userProfile := new(models.UserProfile)
	userProfile.Sex = true
	userProfile.UpdatedAt = time.Now()
	userProfile.CreatedAt = time.Now()
	userProfile.Users = user

	created, _, err := ormer.ReadOrCreate(user, "Mobile")
	_, _, _ = ormer.ReadOrCreate(userProfile, "Users")

	//_, _ = ormer.Insert(userProfile)

	data := make(map[string] interface{})
	data["user"] = user

	//this.Data["json"] = Data{SUCCESS_CODE, MESSAGE, data}

	if err != nil {
		data["err"] = err
		this.Data["json"] = Data{100004, "系统错误", data}

	} else {
		auth := make(map[string] interface{})
		auth["authorization"] = CreateToken(user.Mobile)
		auth["exp"] = TokenExp()

		data["auth"] =  auth

		if created {
			this.Data["json"] = Data{SuccessCode, Message, data}
		} else {
			this.Data["json"] = Data{100003, "该手机号码已存在", data}
		}

	}

	this.ServeJSON()
}

func TokenExp() int {
	tokenExp,_ :=strconv.Atoi(beego.AppConfig.String("TokenExp"))
	expTime := time.Now().Add(time.Hour * time.Duration(tokenExp)).Unix()

	return int(expTime)
}

func CreateToken(Phone string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims :=make(jwt.MapClaims)
	claims["exp"] = TokenExp()
	claims["iat"] = time.Now().Unix()
	claims["phone"] = Phone
	token.Claims = claims
	tokenString,_ := token.SignedString([]byte(beego.AppConfig.String("TokenSecrets")))
	return tokenString
}

func (this *LoginController) GetUser() {

	//this.Abort("404")

	token := this.Ctx.Request.Header.Get("authorization")

	defer func() {
		//捕获到异常错误 返回信息401未授权
		if r := recover(); r != nil {
			this.Data["json"] = Data{401, "未查询到用户", make(map[string] string)}
			this.Ctx.Output.SetStatus(401)
			this.ServeJSON()
			this.StopRun()
		}
	}()

	mobile := CheckToken(token)

	//查询用户
	ormer := orm.NewOrm()
	user := models.Users{}
	user.Mobile = mobile
	//err := ormer.Read(&user, "Mobile")
	//err := ormer.Read(&user)
	err := ormer.QueryTable(user).Filter("Mobile", mobile).Filter("UserProfile__Sex", true).One(&user)
	_, _ = ormer.LoadRelated(&user, "UserProfile")

	//user.UserProfile =
	if err != nil {
		this.Data["json"] = Data{401, "未查询到用户", make(map[string] string)}
	}

	if user.Id > 0 {
		this.ApiJsonReturn(0, "success", user)
	}

	this.ApiJsonReturn(1, "未找到", make(map[string]string))
}

func CheckToken(tokenString string) string {
	var mobile string

	token,_ :=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method")
		}
		return []byte(beego.AppConfig.String("TokenSecrets")),nil
	})
	claims,_:=token.Claims.(jwt.MapClaims)
	mobile =claims["phone"].(string)
	return mobile
}

func (this *LoginController) TestError() {

	phone := this.GetString("phone")

	//_ = errors.New("dsdasd")
	//
	//errors.Is()

	eee := [2][2]string {
		{"wq",phone},
		{"dsad", "adasd"},
	}


	this.ApiJsonReturn(100, "success", eee)

	//this.ApiJsonReturn(100, "success", eee)
}
