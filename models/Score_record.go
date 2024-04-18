package models

import (
	"errors"
	"time"
)

type Score_record struct {
	Id           int64     `json:"id"`                                           //id
	Seniorid     int64     `xorm:"'senior_id' comment('老者id')" json:"senior_id"` //老者id
	Assessors_id int64     `xorm:"comment('所属评估员id')" json:"assessors_id"`       //所属评估员id
	Number_id    int64     `xorm:"comment('编号id')" json:"number_id"`             //编号id
	Kpi_id       int64     `xorm:"comment('指标id')" json:"kpi_id"`                //KPI的id
	Kpiinfo_id   int64     `xorm:"comment('指标分题id')" json:"kpiinfo_id"`          //KPI的id
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
// func GetHealthList(limit int, pagesize int, search *Health) []*Health {
// 	var page int
// 	listdata := []*Health{}
// 	if pagesize-1 < 1 {
// 		page = 0
// 	} else {
// 		page = pagesize - 1
// 	}
// 	if limit <= 6 {
// 		limit = 6

// 	}
// 	session := orm.Table("Health")
// 	// stringid := strconv.FormatInt(search.Id, 10)
// 	if search.Number_id > 0 {
// 		session = session.And("number_id = ?", search.Number_id)
// 	}
// 	if search.Senior_id > 0 {
// 		session = session.And("senior_id = ?", search.Senior_id)
// 	}
// 	if search.Assessors_id > 0 {
// 		session = session.And("assessors_id = ?", search.Assessors_id)
// 	}

// 	// fmt.Println(stringid)

// 	// if search.Title != "" {
// 	// 	title := "%" + search.Title + "%"
// 	// 	session = session.And("title LIKE ?", title)
// 	// 	// session = session.And("pid", rules.Title)
// 	// }
// 	// if search.Categoryid > 0 {
// 	// 	session = session.And("category_id = ?", search.Categoryid)
// 	// }

// 	var byorder string
// 	byorder = "id ASC"
// 	// if order != "" {
// 	// 	byorder = "id DESC"
// 	// }
// 	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
// 	return listdata
// }

// func GetHealthtotal(search *Health) int64 {
// 	var num int64
// 	num = 0
// 	session := orm.Table("health")
// 	if search.Senior_id > 0 {
// 		session = session.And(" Number_id = ?", search.Number_id)
// 	}
// 	a := new(Health)
// 	total, err := session.Count(a)
// 	if err == nil {
// 		num = total

// 	}
// 	return num
// }

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
