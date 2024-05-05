package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Bookfile struct {
	Id      int64     `xorm:"pk autoincr int(11)" json:"id"`                 // ID 自增长
	Bookid  int64     `xorm:"'book_id' comment('所属章节')" json:"book_id"`      // ID 自增长
	Fileurl string    `xorm:"varchar(250) comment('图片视频地址')" json:"fileurl"` // 章节名称
	Title   string    `xorm:"varchar(150) comment('文件名称')" json:"title"`     // 事件名称
	Created time.Time `xorm:"int(11) comment('创建时间')" json:"createtime"`     // 创建时间戳
	Updated time.Time `xorm:"int(11) comment('修改时间')" json:"updatetime"`     // 更改时间戳
}

func (a *Bookfile) TableName() string {
	return "bookfile"
}

// 添加
func Addbookfile(a *Bookfile) error {
	_, err := orm.Insert(a)
	return err
}

// 修改
func Upbookfile(a *Bookfile) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a) //通过添加Cols函数指定需要更新结构体中的哪些值，未指定的将不更新，指定了的即使为0也会更新。
	return affected, err

}
func Deletebookfile(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Bookfile)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
func GetbookfileList(limit int, pagesize int, search *Bookfile, order string) []*Bookfile {
	// orm.DB()
	// listdata = *[]Authrule
	var page int
	listdata := []*Bookfile{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 1 {
		limit = 1000

	}

	session := orm.Table("bookfile")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	// fmt.Println(stringid)
	// stringpid := strconv.FormatInt(search.Id, 10)
	if search.Bookid > 0 {
		session = session.And("book_id = ?", search.Bookid)
	}
	if search.Fileurl != "" {
		filename := "%" + search.Fileurl + "%"
		session = session.And("fileurl LIKE ?", filename)
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

func GetbookfileNum(search *Bookfile) int64 {
	var num int64
	num = 0
	session := orm.Table("bookfile")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id = ?", search.Id)
	}
	// fmt.Println(stringid)
	// stringpid := strconv.FormatInt(search.Id, 10)
	if search.Bookid > 0 {
		session = session.And("book_id = ?", search.Bookid)
	}
	if search.Fileurl != "" {
		filename := "%" + search.Fileurl + "%"
		session = session.And("fileurl LIKE ?", filename)
		// session = session.And("pid", rules.Title)
	}
	a := new(Bookfile)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num

}
