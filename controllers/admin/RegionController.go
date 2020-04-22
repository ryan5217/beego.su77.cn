package admin

import (
	"beego.su77.cn/models"
	"beego.su77.cn/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type RegionController struct {
	BaseController
}

func (this *RegionController) Index() {
	ids := []int{110000, 110102, 110107}

	res := models.GetRegions(ids)
	emptys := make(map[string]interface{})

	fmt.Println(len(res))

	if len(res)  > 0 {
		emptys["data"] = res
		this.Data["json"] = Data{0, "succss", emptys}
	} else {
		emptys["data"] = []int{}
		this.Data["json"] = Data{0, "succss", emptys}
	}

	//this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *RegionController) Values() {
	o := orm.NewOrm()

	var posts [] interface{}
	num, err := o.Raw("SELECT name,title FROM users left join post on post.user_id = users.id WHERE users.id = ?", 4).QueryRows(&posts)
	fmt.Println(posts)

	var maps []orm.Params
	num, err = o.Raw("SELECT area_name, id FROM regions WHERE level = ?", 1).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["area_name"]) // slene
	}
	fmt.Println(maps)

	var lists []orm.ParamsList
	num, err = o.Raw("SELECT area_name, id FROM regions WHERE level = ?", 1).ValuesList(&lists)
	if err == nil && num > 0 {
		fmt.Println(lists[0][0]) // slene
	}
	fmt.Println(lists)

	var list orm.ParamsList
	num, err = o.Raw("SELECT area_name, id  FROM regions WHERE level = ?", 1).ValuesFlat(&list)
	if err == nil && num > 0 {
		fmt.Println(list) // []{"1","2","3",...}
	}
	fmt.Println(list)

	this.StopRun()
	//this.Data["json"] = maps
	//this.ServeJSON()
}

func (this *RegionController) Curd() {
	o := orm.NewOrm()
	user := new(models.Users)
	user.Name = "小刘"
	user.Mobile = "1321312223"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Password = utils.PasswordMD5(user.Mobile, "123456")

	userProfile := new(models.UserProfile)
	userProfile.Sex = true
	userProfile.UpdatedAt = time.Now()
	userProfile.CreatedAt = time.Now()

	userProfile.Users = user

	_, _ = o.Insert(user)
	_, _ = o.Insert(userProfile)
	//
	//_, err := o.Insert(user)
	//if err != nil {
	//	fmt.Println("insert失败")
	//}

	_ = o.Read(user, "id")

	_ = o.Read(user.UserProfile)

	fmt.Println(user, user.UserProfile)
	this.StopRun()
}

func (this *RegionController) Select() {

	//o := orm.NewOrm()
	//user := models.Users{}
	//_ = o.QueryTable(user).Filter("UserProfile__Id", 2).One(&user)
	//fmt.Println(user)
	//fmt.Println(user.UserProfile)
	//this.StopRun()

	o := orm.NewOrm()
	var userProfile []*models.UserProfile
	num, err := o.QueryTable("user_profile").Filter("Users", 3).RelatedSel().OrderBy("-id").All(&userProfile)

	if err == nil {
		fmt.Printf("%d posts read\n", num)
		for _, post := range userProfile {
			fmt.Printf("Id: %d, UserName: %s, Title: %s\n", post.Id, post.Users.Name, post.Users.Mobile)
			fmt.Println(post.Users)
		}
	}
	this.StopRun()

}