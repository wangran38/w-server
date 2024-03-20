package models

//城市后端模型
import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type Country struct {
	Id        int64  `json:"id"`
	Pid       int   `json:"pid"`
	Shortname string `xorm:"varchar(200)" json:"shortname"`
	Name      string `xorm:"varchar(200)" json:"name"`
	Mergename string `xorm:"varchar(200)" json:"mergename"`
	Level     int    `json:"level" xorm:"not null default 1 comment('层级 0 1 2 省市区县') TINYINT"`
	Pinyin    string    `xorm:"varchar(200)" json:"pingyin"`
	Code      string `xorm:"varchar(200)" json:"code"`
	Zip       string `xorm:"varchar(200)" json:"zip"`
	First     string `xorm:"varchar(200)" json:"first"`
	Lng       string `xorm:"varchar(200)" json:"lng"`
	Lat       string `xorm:"varchar(200)" json:"lat"`
}
type Countrytree struct {
	Id        int64  `json:"id"`
	Pid       int   `json:"pid"`
	Shortname string `xorm:"varchar(200)" json:"shortname"`
	Name      string `xorm:"varchar(200)" json:"name"`
	Mergename string `xorm:"varchar(200)" json:"mergename"`
	Level     int    `json:"level" xorm:"not null default 1 comment('层级 0 1 2 省市区县') TINYINT"`
	Pinyin    string    `xorm:"varchar(200)" json:"pingyin"`
	Code      string `xorm:"varchar(200)" json:"code"`
	Zip       string `xorm:"varchar(200)" json:"zip"`
	First     string `xorm:"varchar(200)" json:"first"`
	Lng       string `xorm:"varchar(200)" json:"lng"`
	Lat       string `xorm:"varchar(200)" json:"lat"`
	Children  []*Countrytree
}
func (a *Country) TableName() string {
	return "country"
}
//获取树状数据
func Getcountrytree(pid int) []*Countrytree {
	m := new(Country)
	//不new一个新的，采用结构体，外部无法访问()getruletreee() []*tt这样子，只能new一个，然后去访问
	return m.Treecountrylist(pid)

}

//全部菜单
func (m *Country) Treecountrylist(pid int) []*Countrytree {

	var menus []*Country
	orm.Where("pid = ?", pid).Find(&menus)
	treelist := []*Countrytree{}
	for _, v := range menus {
		child := v.Treecountrylist(int(v.Id))
		node := &Countrytree{
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
func GetCountryList(limit int, pagesize int, search *Country, order string) []*Country {
	// orm.DB()
	// listdata = *[]Authrule
	var page int
	listdata := []*Country{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6
	}

	session := orm.Table("country")
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
	if search.Level >= 1 {
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
func GetCountrytotal(search *Country) int64 {
	var num int64
	num = 0
	session := orm.Table("country")
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	if search.Pid > 0 {
		session = session.And("pid = ?", search.Pid)
	}
	if search.Shortname != "" {
		shortname := "%" + search.Shortname + "%"
		session = session.And("shortname LIKE ?", shortname)
	}
	if search.Level > 0 {
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
	a := new(Country)
	total, err := session.Count(a)
	if err == nil {
		num = total
	}
	return num
}
//add
//添加
func AddCountry(a *Country) error {
	_, err := orm.Insert(a)
	return err
}
//修改
func UpCountry(a *Country) (int64,error) {
	affected, err := orm.Id(a.Id).Update(a) //通过添加Cols函数指定需要更新结构体中的哪些值，未指定的将不更新，指定了的即使为0也会更新。
	return affected, err

}

func DeleteCountry(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Country)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
//根据
func SelectCountryByTitle(Title string) (*Country, error) {
	a := new(Country)
	has, err := orm.Where("name = ?", Title).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}

func Selectcountryid(Id int64) (*Country, error) {
	a := new(Country)
	has, err := orm.Where("id = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("数据出错！")
	}
	return a, nil

}