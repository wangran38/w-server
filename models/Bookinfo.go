package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Bookinfo struct {
	Id          int64     `json:"id"`                                             // ID 自增长
	Pid         int64     `xorm:"comment('姓名')" json:"pid"`                       // ID 自增长
	Chaptername string    `xorm:"varchar(64) comment('章节名称')" json:"chaptername"` // 章节名称
	Level       int       `xorm:"int(3) comment('级别')" json:"level"`              //级别
	Name        string    `xorm:"varchar(64) comment('事件名称')" json:"name"`        // 事件名称
	Content     string    `xorm:"LONGTEXT comment('哪本书')" json:"content"`         // 哪本书
	Created     time.Time `xorm:"int comment('创建时间')" json:"createtime"`          // 创建时间戳
	Updated     time.Time `xorm:"int comment('修改时间')" json:"updatetime"`          // 更改时间戳
}

func (a *Bookinfo) TableName() string {
	return "bookinfo"
}

// 添加
func Addbookinfo(a *Bookinfo) error {
	_, err := orm.Insert(a)
	return err
}

// 修改
func Upbookinfo(a *Bookinfo) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a) //通过添加Cols函数指定需要更新结构体中的哪些值，未指定的将不更新，指定了的即使为0也会更新。
	return affected, err

}
func Deletebookinfo(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Bookinfo)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
func GetbookinfoList(limit int, pagesize int, search *Bookinfo, order string) []*Bookinfo {
	// orm.DB()
	// listdata = *[]Authrule
	var page int
	listdata := []*Bookinfo{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 1 {
		limit = 1000

	}

	session := orm.Table("bookinfo")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	// fmt.Println(stringid)
	// stringpid := strconv.FormatInt(search.Id, 10)
	if search.Pid == 0 {
		session = session.And("Pid = ?", search.Pid)
	}
	if search.Pid >= 1 {
		session = session.And("Pid = ?", search.Pid)
	}
	// fmt.Println(stringpid)

	if search.Chaptername != "" {
		chaptername := "%" + search.Chaptername + "%"
		session = session.And("chaptername LIKE ?", chaptername)
		// session = session.And("pid", rules.Title)
	}
	var byorder string
	byorder = "id ASC"
	if order != "" {
		byorder = "id DESC"
	}
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata

}

func GetbookinfoNum(search *Bookinfo) int64 {
	var num int64
	num = 0
	session := orm.Table("bookinfo")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	// fmt.Println(stringid)
	// stringpid := strconv.FormatInt(search.Id, 10)
	if search.Pid > 0 {
		session = session.And("Pid = ?", search.Pid)
	} else {
		session = session.And("Pid < ?", 1)
	}
	// fmt.Println(stringpid)

	if search.Chaptername != "" {
		chaptername := "%" + search.Chaptername + "%"
		session = session.And("chaptername LIKE ?", chaptername)
		// session = session.And("pid", rules.Title)
	}
	a := new(Bookinfo)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num

}
