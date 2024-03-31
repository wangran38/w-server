package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Kpiinfo struct {
	Id      int64     `json:"id"`                                                //ID自增长
	Kpiid   int64     `xorm:"comment('姓名')" json:"kpi_id"`                       //ID自增长
	Title   string    `xorm:"not null varchar(64) comment('指标名称')" json:"title"` //指标名称
	Score   int       `xorm:"int(10) comment('分数')" json:"score"`                //指标等级
	Created time.Time `xorm:"int comment('创建时间')" json:"createtime"`             //创建时间戳
	Updated time.Time `xorm:"int comment('修改时间')"json:"updatetime"`              //更改时间戳
	Weigh   int       `xorm:"comment('排序')" json:"weigh"`                        //排序，倒序，
	// Is_hassorce int       `xorm:"comment('是否为分值')" json:"weigh"`            //排序，倒序，
	Status int `xorm:"comment('状态')" json:"status"` //状态
}

func (a *Kpiinfo) TableName() string {
	return "kpiinfo"
}

// 添加
func Addkpiinfo(a *Kpiinfo) error {
	_, err := orm.Insert(a)
	return err
}

// 修改
func Upkpiinfo(a *Kpiinfo) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a) //通过添加Cols函数指定需要更新结构体中的哪些值，未指定的将不更新，指定了的即使为0也会更新。
	return affected, err

}
func Deletekpiinfo(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Kpiinfo)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
func GetkpiinfoList(limit int, pagesize int, search *Kpiinfo, order string) []*Kpiinfo {
	// orm.DB()
	// listdata = *[]Authrule
	var page int
	listdata := []*Kpiinfo{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 1 {
		limit = 1000

	}

	session := orm.Table("kpiinfo")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	// fmt.Println(stringid)
	// stringpid := strconv.FormatInt(search.Id, 10)
	if search.Kpiid == 0 {
		session = session.And("kpiid = ?", search.Kpiid)
	}
	if search.Kpiid >= 1 {
		session = session.And("kpiid = ?", search.Kpiid)
	}
	// fmt.Println(stringpid)

	if search.Title != "" {
		title := "%" + search.Title + "%"
		session = session.And("title LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	var byorder string
	byorder = "score DESC"
	// if order != "" {
	// 	byorder = "id DESC"
	// }
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata

}

func GetkpiinfoNum(search *Kpiinfo) int64 {
	var num int64
	num = 0
	session := orm.Table("kpiinfo")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	// fmt.Println(stringid)
	// stringpid := strconv.FormatInt(search.Id, 10)
	if search.Kpiid > 0 {
		session = session.And("kpiid = ?", search.Kpiid)
	} else {
		session = session.And("kpiid < ?", 1)
	}
	// fmt.Println(stringpid)

	if search.Title != "" {
		title := "%" + search.Title + "%"
		session = session.And("title LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	a := new(Kpiinfo)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num

}
