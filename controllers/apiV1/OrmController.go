package apiV1

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"time"
)

//id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
//  `status` tinyint(1) unsigned DEFAULT '0',

type Test struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Status int `json:"status"`
}

type User struct {
	Id int64
	Name string `orm:"size(128)"`
	UserName string `orm:"index;size(128)"`
	Password string `orm:"size(128)"`
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
	ApiToken string `orm:"size(256);column(token)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	DeletedAt time.Time `orm:"null;type(datetime)"`
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"` //设置多对多反向关系
}


type GetOrmController struct {
	beego.Controller
}

func init() {
	//192.168.10.10
	//_ = orm.RegisterDataBase("test", "mysql", "homestead:secret@tcp(192.168.10.10:3306)/test", 30)

	orm.RegisterModel(new(Test), new(User), new(Profile), new(Post), new(Tag))

	//defer orm.DebugLog

	//_ = orm.RunSyncdb("default", false, true)
}

func (c *GetOrmController) GetOrm() {

	ormer := orm.NewOrm()
	//
	param, _ := c.GetInt64("id")

	//test := Test{Name:"ryan123", Status:param}

	//_, _ = orm.Insert(&test)
	//fmt.Println(param)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)

	//// insert
	//id, err := o.Insert(&test)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//
	//// update
	//test.Name = "astaxie"
	//num, err := o.Update(&test)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//// read one
	testRead := Test{Id: param}
	_ = ormer.Read(&testRead)
	//if ss {
	//	c.Data["json"] =
	//}
	//fmt.Printf("ERR: %v\n", read)

	c.Data["json"] = testRead
	c.ServeJSON()
	//c.TplName = "apiV1/getorm.tpl"
	//fmt.Printf("ERR: %v\n", read)
	//
	//// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

}

func (c *GetOrmController) TestOrm() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i
	}

	qiepian := [2][2]string{}
	qiepian[0][0] = "你好"
	qiepian[0][1] = "qwewq"

	qiepian[1][0] = "你好时候死额"
	qiepian[1][0] = "霓虹"

	monster := make(map[string]interface{})
	monster["student"] = "1233123"
	monster["geenv"] = os.Getenv("GOPATH")
	monster["hhe"] = qiepian
	monster["length"] = len(qiepian)
	monster["arr"] = arr

	c.Data["json"] = monster
	c.ServeJSON()

}