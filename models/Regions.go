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