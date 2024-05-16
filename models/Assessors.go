package models

import (
	"errors"
	"strconv"
	"time"

	// "fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Assessors struct {
	Id           int64
	Username     string    `xorm:"varchar(200)" json:"username"`
	Nickname     string    `xorm:"varchar(200)" json:"nickname"`
	Phone        string    `xorm:"varchar(15)" json:"phone"`
	Email        string    `xorm:"varchar(200)" json:"email"`
	Salt         string    `xorm:"varchar(200)" json:"salt"`
	Age          int       `xorm:"int(2)" json:"age"`
	Avatar       string    `xorm:"TEXT" json:"avatar"`
	Loginfailure int       `xorm:"int(10)" json:"loginfailure"`
	Logintime    int       `xorm:"int(10)" json:"logintime"`
	Loginip      string    `xorm:"varchar(200)" json:"loginip"`
	Token        string    `xorm:"varchar(59)"`
	Password     string    `xorm:"varchar(200)"`
	Created      time.Time `xorm:"created"`
	Updated      time.Time `xorm:"updated"`
}
type Assessorsjson struct {
	Id        int64     `json:"id"`
	Gid       int64     `json:"gid"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Password  string    `xorm:"varchar(200)"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Groupname string    `json:"groupname"`
	Created   time.Time `json:"createtime"`
}

func (a *Assessors) TableName() string {
	return "assessors"
}

type AssessorsGroup struct {
	Assessors  `xorm:"extends"`
	Authaccess `xorm:"extends"`
	Authgroup  `xorm:"extends"`
}

// users := make([]UserGroup, 0)
// engine.Join("INNER", "group", "group.id = user.group_id").Find(&users)
// 根据用户名密码查询用户
func SelectAssessorsbyphone(phone string) (*Assessors, error) {
	a := new(Assessors)
	has, err := orm.Where("phone = ?", phone).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户未找到！")
	}
	return a, nil

}

// 根据用户id找用户返回数据
func SelectAssessorsById(Id string) (*Assessors, error) {
	a := new(Assessors)
	id, _ := strconv.ParseInt(Id, 10, 64)
	has, err := orm.Where("id = ?", id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户未找到！")
	}
	return a, nil

}

// add
func AddAssessors(a *Assessors) error {
	_, err := orm.Insert(a)
	return err
}

// 分页列表
// 列表
func GetAssessorsList(limit int, pagesize int, search string, order string) []*Assessorsjson {
	//fmt.Println("搜索关键词",search)
	//    limit,_ := strconv.Atoi(limits)
	//
	//   if pagesize-1<1 {
	page := pagesize - 1
	//   }
	var listAssessors []*AssessorsGroup
	//拼接搜索分页查询语句
	var byorder string
	byorder = "a.id ASC"
	if order == "-id" {
		byorder = "a.id DESC"
	}
	if search != "" {
		orm.Table("assessors").Alias("a").
			Cols("a.*,ac.*, g.id,g.name").
			Join("INNER", []string{"auth_group_access", "ac"}, "ac.uid = a.id").
			Join("INNER", []string{"auth_group", "g"}, "g.id = ac.gid").
			Where("a.phone like ?", "%"+search+"%").
			OrderBy(byorder).
			// Orderby(byorder).
			Limit(limit, limit*page).
			Find(&listAssessors)
		//  orm.Where("username like ?", "%"+search+"%").Limit(limit*pagesize, pagesize).Find(&listAssessors)
	} else {
		orm.Table("assessors").Alias("a").
			Cols("a.*,ac.*, g.id,g.name").
			Join("INNER", []string{"auth_group_access", "ac"}, "a.id = ac.uid").
			Join("INNER", []string{"auth_group", "g"}, "g.id = ac.gid").
			// Where("a.username like ?", "%"+search+"%").
			OrderBy(byorder).
			Limit(limit, limit*page).
			Find(&listAssessors)
	}
	//    fmt.Println(listAssessors)
	Assessorslist := []*Assessorsjson{}
	for _, v := range listAssessors {
		// 结构体exends要用结构体里的结构体去读
		node := &Assessorsjson{
			Id:        v.Assessors.Id,
			Gid:       v.Authgroup.Id,
			Username:  v.Assessors.Username,
			Nickname:  v.Assessors.Nickname,
			Avatar:    v.Assessors.Avatar,
			Phone:     v.Assessors.Phone,
			Email:     v.Assessors.Email,
			Groupname: v.Authgroup.Name,
			Created:   v.Assessors.Created,
		}
		Assessorslist = append(Assessorslist, node)
	}
	return Assessorslist
}
func GetAssessorstotal(search string) int64 {
	var num int64
	num = 0
	a := new(Assessors)
	if search != "" {
		total, err := orm.Cols("id", "phone").Where("phone = ?", "%"+search+"%").Count(a)
		if err == nil {
			num = total

		}
	} else {
		total, err := orm.Cols("id", "phone").Count(a)
		if err == nil {
			num = total

		}
	}
	return num
}

// 删除
func DeleteAssessors(id int64) int {
	a := new(Assessors)
	outnum, _ := orm.ID(id).Delete(a)
	return int(outnum)

}

// 修改
func UpAssessors(a *Assessors) (int64, error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}
