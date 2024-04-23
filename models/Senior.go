package models

import (
	"errors"
	"time"
)

// 老人基础信息表
type Senior struct {
	Id              int64  //ID自增长
	Assessor_id     int64  `xorm:"int(11)" json:"assessor_id"`
	Number_id       int64  `xorm:"int(11)" json:"number_id"`
	Senior_name     string `xorm:"varchar(64) comment('姓名')" json:"senior_name"`                               //姓名
	Senior_gender   int    `json:"senior_gender" xorm:"not null default 1 comment('是否启用 默认1 是男 2 女') TINYINT"` //性别
	Senior_dob      string `xorm:"varchar(100)" json:"senior_dob"`                                             //出生日期
	Birthday        int    `xorm:"int(11)" json:"birthday"`                                                    //出生日期时间戳
	Height          int    `xorm:"int(11)" json:"height"`                                                      //身高
	Weight          int    `xorm:"int(11)" json:"weight"`                                                      //体重
	Ethnic          string `xorm:"varchar(150)" json:"ethnic"`                                                 //民族
	Religion        string `xorm:"varchar(150)" json:"religion"`                                               //宗教
	Idnumber        string `xorm:"varchar(18)" json:"idnumber"`                                                //身份证号
	Education_level string `xorm:"varchar(100)" json:"education_level"`                                        //文化程度
	Live_way        string `xorm:"varchar(100)" json:"live_way"`                                               //居家方式
	Is_marriage     string `xorm:"varchar(100)" json:"is_marriage"`                                            //婚姻状况
	// Ssnumber      string `xorm:"varchar(64)" json:"ssnumber"`                                                //社保号码
	Payment   string    `xorm:"varchar(100)" json:"payment"`                                        //医疗费用支付方式
	Financial string    `xorm:"varchar(100)" json:"financial"`                                      //经济来源
	Is_del    int       `json:"is_del" xorm:"not null default 1 comment('是否启用 默认1 是 2 否') TINYINT"` //性别
	Created   time.Time `json:"createtime" xorm:"int"`
	Updated   time.Time `json:"updatetime" xorm:"int"`
}

func (a *Senior) TableName() string {
	return "senior" //数据库表的别名
}

func AddSenior(a *Senior) error {
	_, err := orm.Insert(a)
	return err
}

// //修改
// func UpUser(a *User) (int64,error) {
// 	affected, err := orm.Id(a.Id).Update(a)
// 	return affected, err

// }
// //根据
func SelectSeniorbyidnum(Idnum string) (*Senior, error) {
	a := new(Senior)
	// title := "%" + Idnum + "%"
	has, err := orm.Where("idnumber = ?", Idnum).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {

		return nil, errors.New("未找到！")
	}
	return a, nil
}

// func AddAuthaccess(a *Olduser) error {
// 	_, err := orm.Insert(a)
// 	return err
// }

// 删除
func DelSenior(id int64) int {
	a := new(Senior)
	outnum, _ := orm.Where("id = ?", id).Delete(a)
	return int(outnum)

}
func GetSeniorList(limit int, pagesize int, search *Senior, order string) []*Senior {
	//fmt.Println("搜索关键词",search)
	//    limit,_ := strconv.Atoi(limits)
	//
	//   if pagesize-1<1 {
	page := pagesize - 1
	//   }
	listdata := []*Senior{}
	//拼接搜索分页查询语句
	var byorder string
	byorder = "weigh DESC"
	if order == "-id" {
		byorder = "weigh DESC"
	}
	orm.Table("senior").
		// Where("kname like ?", "%"+search.Kname+"%").
		OrderBy(byorder).
		// Orderby(byorder).
		Limit(limit, limit*page).
		Find(&listdata)
	//  orm.Where("username like ?", "%"+search+"%").Limit(limit*pagesize, pagesize).Find(&listadmin)
	//    fmt.Println(listadmin)
	return listdata
}

func GetSeniortotal(search *Senior) int64 {
	var num int64
	num = 0
	a := new(Senior)
	total, err := orm.Cols("id", "kname").Where("kname like ?", "%"+search.Senior_name+"%").Count(a)
	if err == nil {
		num = total

	}
	return num
}
