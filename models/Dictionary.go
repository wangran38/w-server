package models

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

// 建结构体(定义了数据库中的字段)
type Dictionary struct {
	Id       int64  `json:"id"`                                       //ID自增长
	Code     int64  `xorm:"comment('编号')" json:"code"`                //编号
	Codename string `xorm:"comment('编号拼音')" json:"codename"`          //编号拼音
	Name     string `xorm:"varchar(200) comment('字典名称')" json:"name"` //字典名称
}

// tableName（）方法，返回数据库表的名称
func (a *Dictionary) TableName() string {
	return "dictionary" //数据库表的名
}

// 添加方法（用于向数据库中插入数据）
func AddDictionary(a *Dictionary) error {
	_, err := orm.Insert(a)

	return err
}

// 修改（用于更新数据库中的数据）
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

// 根据前端参数查询Dictionary表的全部内容（并进行分页）
func SelectDictionarylist(limit int, pagesize int, search *Dictionary) []*Dictionary {
	//定义变量 listdata 用于存储查询结果的Dictionary切片
	listdata := []*Dictionary{}
	//page用于计算分页的页面索引
	var page int
	//如pagesize-1< 1 则设置页面为0 否则计算页面索引
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	//设置限制条件 限制数量<=6则将限制数量设置为6
	if limit <= 6 {
		limit = 6

	}
	//构建查询语句
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

	var byorder string
	byorder = "id DESC"
	//执行查询（将结果存储在listdata中，最后返回该列表）
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

// 定义一个函数用于获取符合特定条件的Dictionary总数
func GetDictionarytotal(search *Dictionary) int64 {
	//定义变量num用于存储总数
	var num int64
	num = 0
	//创建一个orm.Table("news")的会话
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
	//创建一个新的Dictionary实列a
	//使用session.Count(a)计算符合条件的总数
	a := new(Dictionary)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

// 定义一个函数用于删除指定id的数据 （int64的id作为参数）
func DeleteDictionary(id int64) int {
	//定义一个新的Dictionary实列a
	a := new(Dictionary)
	//使用orm.ID(id)指定要删除的id
	//Delete（a）执行删除操作 并存储在outnum中
	outnum, _ := orm.ID(id).Delete(a)
	//将outnum转换为int类型并返回
	return int(outnum)
}
