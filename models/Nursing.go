package models

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 建结构体(定义了数据库中的字段)
type Nursing struct {
	Id          int64     `json:"id"`                                   //ID自增长
	Shift       string    `xorm:"comment('班次')" json:"shiftshift"`      //班次
	Timeperiod  string    `xorm:"comment('时间段')" json:"timeperiod"`     //时间段
	Nursingid   int64     `xorm:"comment('编号')" json:"nursingid"`       //护理动作id
	Nursingname string    `xorm:"comment('护理动作名称')" json:"nursingname"` //护理动作名称
	Created     time.Time `json:"createtime" xorm:"int"`
	Updated     time.Time `json:"updatetime" xorm:"int"`
}

// tableName（）方法，返回数据库表的名称
func (a *Nursing) TableName() string {
	return "nursing" //数据库表的名
}

// 添加方法（用于向数据库中插入数据）
func AddNursing(a *Nursing) error {
	_, err := orm.Insert(a)

	return err
}

// 修改（用于更新数据库中的数据）
func UpNursing(a *Nursing) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err
}

// 根据id查询数据
func SelectNursingid(Id int64) (*Nursing, error) {
	a := new(Nursing)
	has, err := orm.Where("id = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {

		return nil, errors.New("未找到！")
	}
	return a, nil
}

// 根据前端参数查询Dictionary表的全部内容（并进行分页）
func SelectNursinglist(limit int, pagesize int, search *Nursing) []*Nursing {
	//定义变量 listdata 用于存储查询结果的Dictionary切片
	listdata := []*Nursing{}
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
	session := orm.Table("nursing")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	// fmt.Println(stringid)

	if search.Shift != "" {
		title := "%" + search.Shift + "%"
		session = session.And("shift LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}

	if search.Timeperiod != "" {
		title1 := "%" + search.Timeperiod + "%"
		session = session.And("timeperiod LIKE ?", title1)
		// session = session.And("pid", rules.Title)
	}
	if search.Nursingid > 0 {
		// title := "%" + search.Codename + "%"
		session = session.And("nursingid = ?", search.Nursingid)
		// session = session.And("pid", rules.Title)
	}
	if search.Nursingname != "" {
		title := "%" + search.Nursingname + "%"
		session = session.And("nursingname  LIKE ?", title)
	}

	var byorder string
	byorder = "id DESC"
	//执行查询（将结果存储在listdata中，最后返回该列表）
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

// 定义一个函数用于获取符合特定条件的Dictionary总数
func GetNursingtotal(search *Nursing) int64 {
	//定义变量num用于存储总数
	var num int64
	num = 0
	//创建一个orm.Table("news")的会话
	session := orm.Table("nursing")
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	if search.Shift != "" {
		title := "%" + search.Shift + "%"
		session = session.And("shift LIKE ?", title)
	}
	if search.Timeperiod != "" {
		title1 := "%" + search.Timeperiod + "%"
		session = session.And("timeperiod LIKE ?", title1)
		// session = session.And("pid", rules.Title)
	}
	if search.Nursingid > 0 {
		// title := "%" + search.Codename + "%"
		session = session.And("nursingid = ?", search.Nursingid)
		// session = session.And("pid", rules.Title)
	}
	if search.Nursingname != "" {
		title := "%" + search.Nursingname + "%"
		session = session.And("nursingname  LIKE ?", title)
	}
	//创建一个新的Dictionary实列a
	//使用session.Count(a)计算符合条件的总数
	a := new(Nursing)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

// 定义一个函数用于删除指定id的数据 （int64的id作为参数）
func DeleteNursing(id int64) int {
	//定义一个新的Dictionary实列a
	a := new(Nursing)
	//使用orm.ID(id)指定要删除的id
	//Delete（a）执行删除操作 并存储在outnum中
	outnum, _ := orm.ID(id).Delete(a)
	//将outnum转换为int类型并返回
	return int(outnum)
}
