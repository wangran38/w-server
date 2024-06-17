package models

import (
	"errors"
	"time"
	"strings"
)

type Score_record struct {
	Id           int64     `json:"id"`                                           //id
	Seniorid     int64     `xorm:"'senior_id' comment('老者id')" json:"senior_id"` //老者id
	Assessors_id int64     `xorm:"comment('所属评估员id')" json:"assessors_id"`       //所属评估员id
	Number_id    int64     `xorm:"comment('编号id')" json:"number_id"`             //编号id
	Kpi_id       int64     `xorm:"comment('指标id')" json:"kpi_id"`                //KPI的id
	Kpiinfo_id   int64     `xorm:"comment('指标分题id')" json:"kpiinfo_id"`          //KPI的id
	Kpiids      string  `xorm:"-" json:"Kpiids"`  //KPI的id
	Score        int       ` xorm:"comment('分值')" json:"score"`                  //药物名称
	Created      time.Time `xorm:"int" json:"createtime"`
	Updated      time.Time `xorm:"int" json:"updatetime"`
}

func (a *Score_record) TableName() string {
	return "score_record"
}

// 根据用户id找用户返回数据
func SelectScore_record_id(id int64) (*Score_record, error) {
	a := new(Score_record)
	has, err := orm.Where("id = ?", id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("数据出错！")
	}
	return a, nil

}

// 添加
func AddScore_record(a *Score_record) error {
	_, err := orm.Insert(a)
	return err
}
func AddScore_record_arr(a []*Score_record) error {
	_, err := orm.Insert(&a)
	return err
}

// // 修改
// func UpHealth(a *Health) (int64, error) {
// 	affected, err := orm.Id(a.Number_id).Update(a)
// 	return affected, err

// }
func GetScore_recordList(limit int, pagesize int, search *Score_record) []*Score_record {
	var page int
	listdata := []*Score_record{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 100 {
		limit = 100

	}
		//拼接搜索分页查询语句
	var byorder string
	byorder = "kpi_id DESC"
	// if order == "-id" {
	// 	byorder = "id DESC"
	// }
	session := orm.Table("score_record")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Assessors_id > 0 {
		session = session.And("assessors_id = ?", search.Assessors_id)
	}
	if search.Number_id > 0 {
		session = session.And("number_id = ?", search.Number_id)
	}
	if search.Seniorid > 0 {
		session = session.And("senior_id = ?", search.Seniorid)
	}
	// if search.Kpi_id > 0 {
	// 	session = session.And("kpi_id = ?", search.Kpi_id)
	// }

	if search.Kpiinfo_id > 0 {
		session = session.And("kpiinfo_id = ?", search.Kpiinfo_id)
	}
	if search.Kpiids!="" {
		ids := strings.Split(search.Kpiids, ",") //转成数组用orm in
		session = session.In("kpi_id", ids)
	}
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

func GetScore_recordtotal(search *Score_record) int64 {
	var num int64
	num = 0
	session := orm.Table("score_record")
		if search.Assessors_id > 0 {
		session = session.And("assessors_id = ?", search.Assessors_id)
	}
	if search.Number_id > 0 {
		session = session.And("number_id = ?", search.Number_id)
	}
	if search.Seniorid > 0 {
		session = session.And("senior_id = ?", search.Seniorid)
	}
	// if search.Kpi_id > 0 {
	// 	session = session.And("kpi_id = ?", search.Kpi_id)
	// }
	if search.Kpiinfo_id > 0 {
		session = session.And("kpiinfo_id = ?", search.Kpiinfo_id)
	}
	if search.Kpiids!="" {
		ids := strings.Split(search.Kpiids, ",") //转成数组用orm in
		session = session.In("kpi_id", ids)
	}
	a := new(Score_record)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

// func DeleteHealth(Number_id int64) int {
// 	// intid, _ := strconv.ParseInt(id, 10, 64)
// 	a := new(Health)
// 	outnum, _ := orm.ID(Number_id).Delete(a)

// 	return int(outnum)

// }

// // 根据
// func SelectHealthByTitle(Drugname string) (*Health, error) {
// 	a := new(Health)
// 	has, err := orm.Where("title = ?", Drugname).Get(a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !has {
// 		return nil, errors.New("未找到！")
// 	}
// 	return a, nil

// }
