package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Regions struct {
	Id int `json:"id"`
	AreaName string `json:"area_name"`
	ParentId int `json:"parent_id"`
	CityCode string `json:"city_code"`
	Lng string `json:"lng"`
	Lat string `json:"lat"`
	Level int `json:"level"`
	Sort int `json:"sort"`
}

func GetRegions(ids []int) []Regions {

	o := orm.NewOrm()

	var regions []Regions
	_, err := o.Raw("SELECT * FROM regions WHERE id In (?, ?, ?)", ids).QueryRows(&regions)

	if err != nil {
		fmt.Println("没查到任何内容")
	}

	return regions
}

func GetRegionsByParentId(id int) []Regions {
	o := orm.NewOrm()
	var regions []Regions

	_, _ = o.QueryTable(new(Regions)).Filter("ParentId", id).All(&regions)

	return regions
}

func GetLevelOne(id int) []Regions {
	o := orm.NewOrm()

	var regions []Regions

	if id > 0 {
		_, _ = o.QueryTable(new(Regions)).Filter("Level", 1).Filter("Id", id).All(&regions)

	} else {
		_, _ = o.QueryTable(new(Regions)).Filter("Level", 1).All(&regions)
	}
	//ids := [2]int{360000, 110000}

	return regions
}

func GetLevelTwo() []Regions {
	o := orm.NewOrm()

	var regions []Regions

	_, _ = o.QueryTable(new(Regions)).Filter("Level", 2).All(&regions)

	return regions
}

func GetLevelThree() []Regions {
	o := orm.NewOrm()

	var regions []Regions

	_, _ = o.QueryTable(new(Regions)).Filter("Level", 3).All(&regions)

	return regions
}