package models

import (
	"errors"
	"time"
)

type Information struct {
	Id           int64     `json:"id"`                                     //id
	Senior_id    int64     `xorm:"comment('老者id')" json:"senior_id"`       //老者id
	Assessors_id int64     `xorm:"comment('所属评估员id')" json:"assessors_id"` //所属评估员id
	Number_id    int64     `xorm:"comment('评估编号id')" json:"number_id"`     //评估编号id
	Provider     string    `xorm:"comment('提供者姓名')" json:"provider"`       //提供者姓名
	Relationship string    ` xorm:"varchar(200)" json:"relationship"`      //关系
	Contactname  string    `xorm:"TEXT " json:"contactname"`               //提供者姓名
	Phone        string    `xorm:"TEXT" json:"phone"`                      //电话
	Created      time.Time `xorm:"int" json:"createtime"`
	Updated      time.Time `xorm:"int" json:"updatetime"`
}

func (a *Information) TableName() string {

	return "information"
}

// 根据用户id找用户返回数据
func SelectInformationid(id int64) (*Information, error) {
	a := new(Information)
	has, err := orm.Where(" id = ?", id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("数据出错！")
	}
	return a, nil

}

// 根据评估员id返回数据
func SelectInformationId(id int64) (*Information, error) {
	a := new(Information)
	has, err := orm.Where(" id = ?", id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("数据出错！")
	}
	return a, nil

}

// 添加
func AddInformation(a *Information) error {
	_, err := orm.Insert(a)
	return err
}

// 修改
func UpInformation(a *Information) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}
func GetInformationList(limit int, pagesize int, search *Information) []*Information {
	var page int
	listdata := []*Information{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6

	}
	session := orm.Table("information")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	if search.Number_id > 0 {
		session = session.And("number_id = ?", search.Number_id)
	}
	if search.Senior_id > 0 {
		session = session.And("senior_id = ?", search.Senior_id)
	}
	if search.Assessors_id > 0 {
		session = session.And("assessors_id = ?", search.Assessors_id)
	}

	var byorder string
	byorder = "id ASC"
	// if order != "" {
	// 	byorder = "id DESC"
	// }
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

func GetInformationtotal(search *Information) int64 {
	var num int64
	num = 0
	session := orm.Table("information")
	if search.Senior_id > 0 {
		session = session.And(" Id = ?", search.Id)
	}
	a := new(Information)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

func DeleteInformation(Id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Information)
	outnum, _ := orm.ID(Id).Delete(a)

	return int(outnum)

}

// 根据
func SelectInformatioByTitle(Contactname string) (*Information, error) {
	a := new(Information)
	has, err := orm.Where("title = ?", Contactname).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}
