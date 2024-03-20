package models

//城市后端模型
import (
	 "errors"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id        int64  `json:"id"`
	Pid       int   `json:"pid"`
	Shortname string `xorm:"varchar(200)" json:"shortname"`
	Name      string `xorm:"varchar(200)" json:"name"`
	Mergename string `xorm:"varchar(200)" json:"mergename"`
	Level     int    `json:"status" xorm:"not null default 1 comment('层级 0 1 2 省市区县') TINYINT"`
	Pinyin    string    `xorm:"varchar(200)" json:"pingyin"`
	Code      string `xorm:"varchar(200)" json:"code"`
	Zip       string `xorm:"varchar(200)" json:"zip"`
	First     string `xorm:"varchar(200)" json:"first"`
	Lng       string `xorm:"varchar(200)" json:"lng"`
	Lat       string `xorm:"varchar(200)" json:"lat"`
}
type Citytree struct {
	Id        int64  `json:"id"`
	Pid       int   `json:"pid"`
	Shortname string `xorm:"varchar(200)" json:"shortname"`
	Name      string `xorm:"varchar(200)" json:"name"`
	Mergename string `xorm:"varchar(200)" json:"mergename"`
	Level     int    `json:"status" xorm:"not null default 1 comment('层级 0 1 2 省市区县') TINYINT"`
	Pinyin    string    `xorm:"varchar(200)" json:"pingyin"`
	Code      string `xorm:"varchar(200)" json:"code"`
	Zip       string `xorm:"varchar(200)" json:"zip"`
	First     string `xorm:"varchar(200)" json:"first"`
	Lng       string `xorm:"varchar(200)" json:"lng"`
	Lat       string `xorm:"varchar(200)" json:"lat"`
	Children  []*Citytree
}
func (a *City) TableName() string {
	return "area"
}
//获取树状数据
func Getcitytree(pid int) []*Citytree {
	m := new(City)
	//不new一个新的，采用结构体，外部无法访问()getruletreee() []*tt这样子，只能new一个，然后去访问
	return m.Treecitylist(pid)

}

//全部菜单
func (m *City) Treecitylist(pid int) []*Citytree {

	var menus []*City
	orm.Where("pid = ?", pid).Find(&menus)
	treelist := []*Citytree{}
	for _, v := range menus {
		child := v.Treecitylist(int(v.Id))
		node := &Citytree{
			Id:        v.Id,
			Pid:       v.Pid,
			Name:     v.Name,
		}
		node.Children = child
		treelist = append(treelist, node)
	}
	return treelist

}
//根据用户名密码查询用户
func GetCityList(limit int, pagesize int, search *City, order string) []*City {
	// orm.DB()
	// listdata = *[]Authrule
	var page int
	listdata := []*City{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6
	}

	session := orm.Table("area")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id >= 1 {
		session = session.And("id = ?", search.Id)
	}
	if search.Pid >= 1 {
		session = session.And("pid = ?", search.Pid)
	}
	if search.Shortname != "" {
		shortname := "%" + search.Shortname + "%"
		session = session.And("shortname LIKE ?", shortname)
	}
	if search.Level >=1 {
		session = session.And("level = ?", search.Level)
	}
	if search.Name != "" {
		name := "%" + search.Name + "%"
		session = session.And("name LIKE ?", name)

	}
	if search.Pinyin != "" {
		pinyin := "%" + search.Pinyin + "%"
		session = session.And("pinyin LIKE ?", pinyin)
	}
	if search.Code != "" {
		// code := "%" + search.Code + "%"
		session = session.And("code = ?", search.Code)
	}
	if search.Zip != "" {
		zip := "%" + search.Zip + "%"
		session = session.And("zip LIKE ?", zip)
	}
	if search.First != "" {
		first := "%" + search.First + "%"
		session = session.And("first LIKE ?", first)
	}
	var byorder string
	byorder = "id ASC"
	if order == "+id" {
		byorder = "id ASC"
	}
	if order == "-id" {
		byorder = "id DESC"
	}
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata

}
func GetCitytotal(search *City) int64 {
	var num int64
	num = 0
	session := orm.Table("area")
	if search.Id >= 1 {
		session = session.And("id = ?", search.Id)
	}
	if search.Pid >= 1 {
		session = session.And("pid = ?", search.Pid)
	}
	if search.Level >= 1 {
		session = session.And("level = ?", search.Level)
	} else {
		session = session.And("level = ?", 1)
	}
	if search.Shortname != "" {
		shortname := "%" + search.Shortname + "%"
		session = session.And("shortname LIKE ?", shortname)
	}
	// if search.Level >=0 {
	// 	// level := "%" + search.Level + "%"
	// 	session = session.And("level", search.Level)

	// }
	if search.Name != "" {
		name := "%" + search.Name + "%"
		session = session.And("name LIKE ?", name)

	}
	if search.Pinyin != "" {
		pinyin := "%" + search.Pinyin + "%"
		session = session.And("pinyin LIKE ?", pinyin)
	}
	if search.Code != "" {
		// code := "%" + search.Code + "%"
		session = session.And("code = ?", search.Code)
	}
	if search.Zip != "" {
		zip := "%" + search.Zip + "%"
		session = session.And("zip LIKE ?", zip)
	}
	if search.First != "" {
		first := "%" + search.First + "%"
		session = session.And("first LIKE ?", first)
	}
	a := new(City)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}
//add
func SelectcityByTitle(Title string) (*City, error) {
	a := new(City)
	has, err := orm.Where("shortname = ?", Title).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}

func Selectcityid(Id int64) (*City, error) {
	a := new(City)
	has, err := orm.Where("id = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("数据出错！")
	}
	return a, nil

}