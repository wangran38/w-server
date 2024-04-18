package models

import (
	"errors"
	"time"
)

type Information struct {
	Id           int64     `json:"id"`                                     //id
	Senior_id    int64     `xorm:"comment('老者id')" json:"senior_id"`       //老者id
	Assessors_id int64     `xorm:"comment('所属评估员id')" json:"assessors_id"` //所属评估员id
	Provider     string    `xorm:"comment('提供者姓名')" json:"provider"`       //所属评估员id
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

// 添加
func AddInformation(a *Information) error {
	_, err := orm.Insert(a)
	return err
}

// 修改
func UpInformation(a *Information) (int64, error) {
	affected, err := orm.Id(a.Senior_id).Update(a)
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
	session := orm.Table("Information")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Senior_id > 0 {
		session = session.And("senior_id = ?", search.Senior_id)
	}
	// fmt.Println(stringid)

	// if search.Title != "" {
	// 	title := "%" + search.Title + "%"
	// 	session = session.And("title LIKE ?", title)
	// 	// session = session.And("pid", rules.Title)
	// }
	// if search.Categoryid > 0 {
	// 	session = session.And("category_id = ?", search.Categoryid)
	// }

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
		session = session.And(" Provider_id = ?", search.Senior_id)
	}
	a := new(Information)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

func DeleteInformation(Senior_id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Information)
	outnum, _ := orm.ID(Senior_id).Delete(a)

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
