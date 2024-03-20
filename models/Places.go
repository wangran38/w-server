package models

import (
	"errors"
	"time"
)

type Places struct {
	Id      int64     `json:"id"`
    Cityid     int64       `json:"city_id"`
	Countryid     int64       `json:"country_id"`
	Title    string    `json:"title" xorm:"varchar(200)"`
	Image   string    `json:"image" xorm:"TEXT "`
	// Keywords   string  `json:"keywords" xorm:"TEXT "`
	// Description   string  `json:"description" xorm:" TEXT "`
    // Oldlink    string    `json:"oldlink" xorm:"varchar(200)"`
	Address    string    `json:"address" xorm:"varchar(200)"`
    Lng float64 `json:"lng"`  //地址经度
    Lat float64 `json:"lat"`  //地址维度
	Content   string  `json:"content" xorm:"LONGTEXT "`
	Isshow     int       `json:"isshow" xorm:"not null default 1 comment('是否启用 默认1 是 0 无') TINYINT"`
	Created time.Time `json:"createtime" xorm:"int"`
	Updated time.Time `json:"updatetime" xorm:"int"`
    Iscountry     int       `json:"iscountry" xorm:"not null default 1 comment('是否启用 默认2 是 1 无') TINYINT"`
	Weigh   int  `json:"weigh"`
	// Status  int       `json:"status"`
}


func (a *Places) TableName() string {
	return "places"
}

//根据用户id找用户返回数据
func SelectPlacesid(Id int64) (*Places, error) {
	a := new(Places)
	has, err := orm.Where("id = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("组别菜单数据出错！")
	}
	return a, nil

}

//添加
func AddPlaces(a *Places) error {
	_, err := orm.Insert(a)
	return err
}
//修改
func UpPlaces(a *Places) (int64,error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}
func GetPlacesList(limit int, pagesize int, search *Places, order string) []*Places {
	var page int
	listdata := []*Places{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6
	}
	session := orm.Table("places")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	// fmt.Println(stringid)

	if search.Title != "" {
		title := "%" + search.Title + "%"
		session = session.And("title LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}


	var byorder string
	byorder = "id DESC"
	if order != "" {
		byorder = "id"
	}
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

func GetPlacestotal(search *Places) int64 {
	var num int64
	num = 0
	session := orm.Table("places")
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	if search.Title != "" {
		name := "%" + search.Title + "%"
		session = session.And("title LIKE ?", name)
		// session = session.And("pid", rules.Title)
	}
	if search.Cityid > 0 {
		session = session.And("cityid = ?", search.Cityid)
	}
	if search.Isshow > 0 {
		session = session.And("isshow = ?", search.Isshow)
		// session = session.And("pid", rules.Title)
	}
	if search.Iscountry > 0 {
		session = session.And("iscountry = ?", search.Iscountry)
        	if search.Countryid > 0 {
		session = session.And("countryid = ?", search.Cityid)
	}
	}
	a := new(Places)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

func DeletePlaces(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Places)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
//根据
func SelectPlacesByTitle(Title string) (*Places, error) {
	a := new(Places)
	has, err := orm.Where("title = ?", Title).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}

//根据
func SelectPlacesByLink(Link string) (*Places, error) {
	a := new(Places)
	has, err := orm.Where("oldlink = ?", Link).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}