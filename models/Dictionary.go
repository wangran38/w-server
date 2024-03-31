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
func AddDictionary(a *Dictionary) error {
	_, err := orm.Insert(a)

	return err
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
func SelectDictionarylist(wk *Dictionary) []*Dictionary {
	// a := new(Dictionary)
	listdata := []*Dictionary{}
	// session := orm.Table("education")
	session := orm.Table("dictionary")
	if wk.Codename != "" {
		session = session.And("codename = ?", wk.Codename)
		// session = session.And("pid", rules.Title)
	}
	// if wk.Code <= 1 {
	// 	session = session.And("code = ?", wk.Code)
	// 	// session = session.And("pid", rules.Title)
	// }
	// if wk.Code == 0 {
	// 	session = session.And("code = ?", wk.Code)
	// 	// session = session.And("pid", rules.Title)
	// }
	// orm.Table("Dictionary").
	// Where("codename = ?",wk.Codename)
	session.Find(&listdata)
	return listdata
}

func DeleteDictionary(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Dictionary)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
