package models

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

// 建结构体
type Dictionary struct {
	Id       int64  `json:"id"`                                       //ID自增长
	Code     int64  `xorm:"comment('编号')" json:"code"`                //编号
	Codename string `xorm:"comment('编号拼音')" json:"codename"`          //编号拼音
	Name     string `xorm:"varchar(200) comment('字典名称')" json:"name"` //字典名称
}

func (a *Dictionary) TableName() string {
	return "dictionary" //数据库表的名
}

// 添加方法
func AddDictionary(a *Dictionary) error {
	_, err := orm.Insert(a)

	return err
}

// 修改
func Updedictionary(a *Dictionary) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err
}

// 根据code查询数据
func SelectDictionarybycode(Code int64) (*Dictionary, error) {
	a := new(Dictionary)
	has, err := orm.Where("code = ?", Code).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {

		return nil, errors.New("未找到！")
	}
	return a, nil
}

// 根据前端参数查询Dictionary表的全部内容
func SelectDictionarylist(limit int, pagesize int, search *Dictionary) []*Dictionary {
	// a := new(Dictionary)
	listdata := []*Dictionary{}
	// session := orm.Table("education")
	var page int
	// listdata := []*News{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6

	}
	session := orm.Table("dictionary")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	// fmt.Println(stringid)

	if search.Codename != "" {
		title := "%" + search.Codename + "%"
		session = session.And("codename LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	if search.Code > 0 {
		// title := "%" + search.Codename + "%"
		session = session.And("code = ?", search.Code)
		// session = session.And("pid", rules.Title)
	}
	// if search.Categoryid > 0 {
	// 	session = session.And("category_id = ?", search.Categoryid)
	// }

	var byorder string
	byorder = "id DESC"
	// if order != "" {
	// 	byorder = "id DESC"
	// }
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

func GetDictionarytotal(search *Dictionary) int64 {
	var num int64
	num = 0
	session := orm.Table("news")
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	if search.Codename != "" {
		title := "%" + search.Codename + "%"
		session = session.And("codename LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	if search.Code > 0 {
		// title := "%" + search.Codename + "%"
		session = session.And("code = ?", search.Code)
		// session = session.And("pid", rules.Title)
	}
	a := new(Dictionary)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

func DeleteDictionary(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Dictionary)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)
}
