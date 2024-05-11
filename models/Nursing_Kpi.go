package models

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 建结构体(定义了数据库中的字段)
type Nursing_Kpi struct {
	Id         int64     `json:"id"`                                  //ID自增长
	Nursing_id string    `xorm:"comment('护理动作编号')" json:"nursing_id"` //护理动作编号
	Kpi_id     string    `xorm:"comment('题目id')" json:"kpi_id"`       //题目id
	P_id       int64     `xorm:"comment('大类小类')" json:"p_id"`         //大类小类
	Created    time.Time `json:"createtime" xorm:"int"`
	Updated    time.Time `json:"updatetime" xorm:"int"`
}

// tableName（）方法，返回数据库表的名称
func (a *Nursing_Kpi) TableName() string {
	return "nursing_kpi" //数据库表的名
}

// 添加方法（用于向数据库中插入数据）
func AddNursing_Kpi(a *Nursing_Kpi) error {
	_, err := orm.Insert(a)

	return err
}

// 修改（用于更新数据库中的数据）
func UpNursing_Kpi(a *Nursing_Kpi) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err
}

// 根据id查询数据
func SelectNursing_Kpiid(Id int64) (*Nursing_Kpi, error) {
	a := new(Nursing_Kpi)
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
func SelectNursing_Kpilist(limit int, pagesize int, search *Nursing_Kpi) []*Nursing_Kpi {
	//定义变量 listdata 用于存储查询结果的Dictionary切片
	listdata := []*Nursing_Kpi{}
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
	session := orm.Table("nursing_kpi")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}

	// if search.Nursing_id > 0 {
	// 	session = session.And(" nursing_id = ?", search.Nursing_id)
	// }

	// if search.Kpi_id > 0 {
	// 	session = session.And(" kpi_id = ?", search.Kpi_id)
	// }
	var byorder string
	byorder = "id DESC"
	//执行查询（将结果存储在listdata中，最后返回该列表）
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

// 定义一个函数用于获取符合特定条件的Dictionary总数
func GetNursing_kpitotal(search *Nursing_Kpi) int64 {
	//定义变量num用于存储总数
	var num int64
	num = 0
	//创建一个orm.Table("news")的会话
	session := orm.Table("nursing_kpi")
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	//
	//创建一个新的Dictionary实列a
	//使用session.Count(a)计算符合条件的总数
	a := new(Nursing_Kpi)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

// 定义一个函数用于删除指定id的数据 （int64的id作为参数）
func DeleteNursing_Kpi(id int64) int {
	//定义一个新的Dictionary实列a
	a := new(Nursing_Kpi)
	//使用orm.ID(id)指定要删除的id
	//Delete（a）执行删除操作 并存储在outnum中
	outnum, _ := orm.ID(id).Delete(a)
	//将outnum转换为int类型并返回
	return int(outnum)
}
