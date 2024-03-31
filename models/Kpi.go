package models

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Kpi struct {
	Id       int64     `json:"id"`                                       //ID自增长
	Pid      int64     `xorm:"comment('姓名')" json:"pid"`                 //ID自增长
	Kname    string    `xorm:"varchar(64) comment('指标名称')" json:"kname"` //指标名称
	Level    int       `xorm:"int(3) comment('级别')" json:"level"`        //指标等级
	Kcontent string    `xorm:"LONGTEXT comment('级别')" json:"kcontent"`   //指标的描述内容
	Created  time.Time `xorm:"int comment('创建时间')" json:"createtime"`    //创建时间戳
	Updated  time.Time `xorm:"int comment('修改时间')"json:"updatetime"`     //更改时间戳
	Weigh    int       `xorm:"comment('排序')" json:"weigh"`               //排序，倒序，
	// Is_hassorce int       `xorm:"comment('是否为分值')" json:"weigh"`            //排序，倒序，
	Status int `xorm:"comment('状态')" json:"status"` //状态
}

type Kpitree struct {
	Id       int64     `json:"id"`
	Pid      int64     `json:"pid"`        //ID自增长
	Kname    string    `json:"kname"`      //指标名称
	Level    int       `json:"level"`      //指标等级
	Kcontent string    `json:"kcontent"`   //指标的描述内容
	Created  time.Time `json:"createtime"` //创建时间戳
	Updated  time.Time `json:"updatetime"` //更改时间戳
	Weigh    int       `json:"weigh"`      //排序，倒序，
	Status   int       `json:"status"`     //状态
	Children []*Kpitree
}

func (a *Kpi) TableName() string {
	return "kpi"
}

// 获取树状数据
func GetKpitree(pid int) []*Kpitree {
	m := new(Kpi)
	//不new一个新的，采用结构体，外部无法访问()getruletreee() []*tt这样子，只能new一个，然后去访问
	return m.TreeKpilist(pid)

}

// 全部菜单
func (m *Kpi) TreeKpilist(pid int) []*Kpitree {

	var menus []*Kpi
	orm.Where("pid = ?", pid).Find(&menus)
	treelist := []*Kpitree{}
	for _, v := range menus {
		child := v.TreeKpilist(int(v.Id))
		node := &Kpitree{
			Id:    v.Id,
			Pid:   v.Pid,
			Kname: v.Kname,
			// Level:  v.Level,
			// Level:  v.Level,
			// Level:  v.Level,
			Status: v.Status,
		}
		node.Children = child
		treelist = append(treelist, node)
	}
	return treelist

}

// 添加
func Addkpi(a *Kpi) error {
	_, err := orm.Insert(a)
	return err
}

// 修改
func Upkpi(a *Kpi) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a) //通过添加Cols函数指定需要更新结构体中的哪些值，未指定的将不更新，指定了的即使为0也会更新。
	return affected, err

}
func GetkpiList(limit int, pagesize int, search *Kpi, order string) []*Kpi {
	//fmt.Println("搜索关键词",search)
	//    limit,_ := strconv.Atoi(limits)
	//
	//   if pagesize-1<1 {
	page := pagesize - 1
	//   }
	listdata := []*Kpi{}
	//拼接搜索分页查询语句
	var byorder string
	byorder = "weigh DESC"
	if order == "-id" {
		byorder = "weigh DESC"
	}
	orm.Table("kpi").
		Where("kname like ?", "%"+search.Kname+"%").
		OrderBy(byorder).
		// Orderby(byorder).
		Limit(limit, limit*page).
		Find(&listdata)
	//  orm.Where("username like ?", "%"+search+"%").Limit(limit*pagesize, pagesize).Find(&listadmin)
	//    fmt.Println(listadmin)
	return listdata
}

func Getkpitotal(search *Kpi) int64 {
	var num int64
	num = 0
	a := new(Kpi)
	total, err := orm.Cols("id", "kname").Where("kname like ?", "%"+search.Kname+"%").Count(a)
	if err == nil {
		num = total

	}
	return num
}

func Deletekpi(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Kpi)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}

// 根据
func SelectkpiByTitle(Kname string) (*Kpi, error) {
	a := new(Kpi)
	has, err := orm.Where("kname = ?", "%"+Kname+"%").Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}

func GetkpiapiList(limit int, pagesize int, search *Kpi, order string) []*Kpi {
	//fmt.Println("搜索关键词",search)
	//    limit,_ := strconv.Atoi(limits)
	//
	//   if pagesize-1<1 {
	page := pagesize - 1
	//   }
	listdata := []*Kpi{}
	//拼接搜索分页查询语句
	var byorder string
	byorder = "weigh DESC"
	if order == "-id" {
		byorder = "weigh DESC"
	}
	session := orm.Table("kpi")
	if search.Id >= 1 {
		session = session.And("id = ?", search.Id)
	}
	if search.Pid == 0 {
		session = session.And("pid = ?", search.Pid)
	}
	if search.Pid >= 1 {
		session = session.And("pid = ?", search.Pid)
	}

	if search.Kname != "" {
		Kname := "%" + search.Kname + "%"
		session = session.And("Kname LIKE ?", Kname)
	}
	if search.Level >= 1 {
		session = session.And("level = ?", search.Level)
	}
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
