package models

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id          int64     `xorm:"pk autoincr int(11)" json:"id"`                  // ID 自增长
	Pid         int64     `xorm:"int(11) comment('父级ID')" json:"pid"`             // ID 自增长
	Chaptername string    `xorm:"varchar(64) comment('章节名称')" json:"chaptername"` // 章节名称
	Level       int       `xorm:"int(3) comment('级别')" json:"level"`              //级别
	Name        string    `xorm:"varchar(64) comment('事件名称')" json:"name"`        // 事件名称
	Content     string    `xorm:"LONGTEXT comment('详细内容')" json:"content"`        // 哪本书
	Isdel       int       `json:"isdel" xorm:"not null default 1 comment('是否启用 默认1 正常 2 已删除') TINYINT"`
	Created     time.Time `xorm:"int(11) comment('创建时间')" json:"createtime"` // 创建时间戳
	Updated     time.Time `xorm:"int(11) comment('修改时间')" json:"updatetime"` // 更改时间戳
}

type Booktree struct {
	Id          int64     `json:"id"`          // ID 自增长
	Pid         int64     `json:"pid"`         // 子级ID 自增长
	Chaptername string    `json:"chaptername"` // 章节名称
	Level       int       `json:"level"`       //级别
	Name        string    `json:"name"`        // 事件名称
	Content     string    `json:"content"`     // 哪本书
	Created     time.Time `json:"createtime"`  // 创建时间戳
	Updated     time.Time `json:"updatetime"`  // 更改时间戳
	Children    []*Booktree
}

func (a *Book) TableName() string {
	return "book"
}

// 获取树状数据
func GetBooktree(pid int64) []*Booktree {
	m := new(Book)
	// 不 new 一个新的，采用结构体，外部无法访问() getruletreee() []*tt 这样子，只能 new 一个，然后去访问
	return m.TreeBooklist(pid)
}

// 全部菜单
func (m *Book) TreeBooklist(pid int64) []*Booktree {

	var menus []*Book
	orm.Cols("id", "pid", "chaptername").Where("pid =?", pid).Find(&menus)
	treelist := []*Booktree{}
	for _, v := range menus {
		child := v.TreeBooklist(int64(v.Id))
		node := &Booktree{
			Id:          v.Id,
			Pid:         v.Pid,
			Chaptername: v.Chaptername,
		}
		node.Children = child
		treelist = append(treelist, node)
	}
	return treelist
}

// 添加
func Addbook(a *Book) error {
	_, err := orm.Insert(a)
	return err
}

// 修改
func Upbook(a *Book) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a) // 通过添加 Cols 函数指定需要更新结构体中的哪些值，未指定的将不更新，指定了的即使为 0 也会更新
	return affected, err
}

// 删除
func Deletebook(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Book)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}

// 获取列表
func GetbookList(limit int, pagesize int, search *Book, order string) []*Book {

	page := pagesize - 1
	listdata := []*Book{}
	// 拼接搜索分页查询语句
	var byorder string
	byorder = "id DESC"
	if order == "-id" {
		byorder = "id DESC"
	}
	orm.Table("book").
		// Where("chaptername like ?", "%"+search.Chaptername+"%").
		OrderBy(byorder).
		// Orderby(byorder).
		Limit(limit, limit*page).
		Find(&listdata)
	//  orm.Where("username like?", "%"+search+"%").Limit(limit*pagesize, pagesize).Find(&listadmin)
	//    fmt.Println(listadmin)
	return listdata
}

// 根据chaptername查找
func SelectbookByTitle(Chaptername string) (*Book, error) {
	a := new(Book)
	has, err := orm.Where("chaptername = ?", "%"+Chaptername+"%").Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}

// 根据用户id找用户返回数据
func Selectbookid(Id int64) (*Book, error) {
	a := new(Book)
	has, err := orm.Cols("id", "content").Where("id = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("数据出错！")
	}
	return a, nil

}
func GetbookapiList(limit int, pagesize int, search *Book, order string) []*Book {
	//fmt.Println("搜索关键词",search)
	//    limit,_ := strconv.Atoi(limits)
	//
	var page int
	// var limit int
	if limit < 1 {
		limit = 6
	}
	// listdata := []*News{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	listdata := []*Book{}
	//拼接搜索分页查询语句
	var byorder string
	byorder = "id DESC"
	if order == "-id" {
		byorder = "id DESC"
	}
	session := orm.Table("book").Cols("id", "pid", "chaptername", "level", "name", "content", "created")
	if search.Id >= 1 {
		session = session.And("id = ?", search.Id)
	}
	// if search.Pid == 0 {
	// 	session = session.And("pid = ?", search.Pid)
	// }
	if search.Pid >= 1 {
		session = session.And("pid = ?", search.Pid)
	}

	if search.Chaptername != "" {
		Chaptername := "%" + search.Chaptername + "%"
		session = session.And("chaptername LIKE ?", Chaptername)
	}
	if search.Level >= 1 {
		session = session.And("level = ?", search.Level)
	}
	if search.Name != "" {
		Name := "%" + search.Name + "%"
		session = session.And("name LIKE ?", Name)
	}
	// if search.Content != "" {
	// 	Content := "%" + search.Content + "%"
	// 	session = session.And("content LIKE ?", Content)
	// }
	// if search.Name != "" {
	// 	name := "%" + search.Name + "%"
	// 	session = session.And("name LIKE ?", name)

	// }
	// Where("kname like ?", "%"+search.Kname+"%").
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata

	//  orm.Where("username like ?", "%"+search+"%").Limit(limit*pagesize, pagesize).Find(&listadmin)
	//    fmt.Println(listadmin)
	// return listdata
}
func Getbooktotal(search *Book) int64 {
	var num int64
	num = 0
	session := orm.Table("book")
	if search.Pid >= 1 {
		session = session.And("pid = ?", search.Pid)
	}

	if search.Chaptername != "" {
		Chaptername := "%" + search.Chaptername + "%"
		session = session.And("chaptername LIKE ?", Chaptername)
	}
	if search.Level >= 1 {
		session = session.And("level = ?", search.Level)
	}
	if search.Name != "" {
		Name := "%" + search.Name + "%"
		session = session.And("name LIKE ?", Name)
	}
	if search.Content != "" {
		Content := "%" + search.Content + "%"
		session = session.And("content LIKE ?", Content)
	}
	a := new(Book)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}
