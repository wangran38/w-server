package models

import (
	"errors"
	"time"
)

// 建结构体
type Healthrelated struct {
	Id             int64     `json:"id"`                                                  //id
	Senior_id      int64     `xorm:"comment('老者id')" json:"senior_id"`                    //老者id
	Assessors_id   int64     `xorm:"comment('所属评估员id')" json:"assessors_id"`              //所属评估员id
	Number_id      int64     `xorm:"comment('编号id')" json:"number_id"`                    //编号id
	Pressureinjury string    `xorm:"varchar(100) comment('压力性损伤')" json:"pressureinjury"` //压力性损伤
	Joint          string    ` xorm:"varchar(100) comment('关节活动度')" json:"Joint"`         //关节活动度
	Affectedarea   string    ` xorm:"varchar(100) comment('关节影响部位')" json:"affectedarea"` //关节影响部位
	Woundcondition string    ` xorm:"varchar(100) comment('伤口情况')" json:"woundcondition"` //伤口情况
	Specialcare    string    ` xorm:"varchar(100) comment('用药剂量')" json:"specialcare"`    //特殊护理
	Painsensation  string    ` xorm:"varchar(100) comment('疼痛感')" json:"painsensation"`   //疼痛感
	Toothloss      string    ` xorm:"varchar(100) comment('牙齿缺失')" json:"toothloss"`      //牙齿缺失
	Wearing        string    ` xorm:"varchar(100) comment('义齿佩戴')" json:"wearing"`        //义齿佩戴
	Swallow        string    ` xorm:"varchar(100) comment('吞咽困难的情形和状况')" json:"swallow"`  //吞咽困难的情形和状况
	Malnutrition   string    ` xorm:"varchar(100) comment('营养不良')" json:"malnutrition"`   //营养不良
	Cleaning       string    ` xorm:"varchar(100) comment('清理呼吸道无效')" json:"cleaning"`    //清理呼吸道无效
	Coma           string    ` xorm:"varchar(100) comment('昏迷')" json:"coma"`             //昏迷
	Other          string    ` xorm:"varchar(100) comment('其他')" json:"other"`            //其他
	Created        time.Time `xorm:"int" json:"createtime"`
	Updated        time.Time `xorm:"int" json:"updatetime"`
}

// Table Name()方法 返回数据库表的名称
func (a *Healthrelated) TableName() string {
	return "healthrelated"
}

// 根据用户id找用户返回数据
func SelectHealthrelatednumber_id(number_id int64) (*Healthrelated, error) {
	a := new(Healthrelated)
	has, err := orm.Where(" number_id = ?", number_id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("数据出错！")
	}
	return a, nil

}

// 添加（用于向数据库中插入数据）
func AddHealthrelated(a *Healthrelated) error {
	_, err := orm.Insert(a)
	return err
}

// 修改（用于更新数据库中的数据）
func UpHealthrelated(a *Healthrelated) (int64, error) {
	affected, err := orm.Id(a.Number_id).Update(a)
	return affected, err

}

// 根据前端参数查询Healthrelated表的全部内容（并进行分页）
func GetHealthrelatedList(limit int, pagesize int, search *Healthrelated) []*Healthrelated {
	//page用于计算分页的页面索引
	var page int
	//定义变量listdata用于存储查询结果的Healthrelated切片
	listdata := []*Healthrelated{}
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
	session := orm.Table("Healthrelated")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Number_id > 0 {
		session = session.And("number_id = ?", search.Number_id)
	}
	// fmt.Println(stringid)

	// if search.Title != "" {
	// 	title := "%" + search.Title + "%"
	// 	session = session.And("title LIKE ?", title)
	// 	// session = session.And("pid", rules.Title)
	// }
	// if search.Categoryid > 0 {
	// 	session = session.And("category_id = ?", search.Categoryid)
	// }

	var byorder string
	byorder = "id ASC"
	// if order != "" {
	// 	byorder = "id DESC"
	// }
	//执行查询（将结果存储在listdata中，最后返回该列表）
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

func GetHealthrelatedtotal(search *Healthrelated) int64 {
	var num int64
	num = 0
	session := orm.Table("healthrelated")
	if search.Senior_id > 0 {
		session = session.And(" Number_id = ?", search.Number_id)
	}
	a := new(Healthrelated)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

// 删除
func DeleteHealthrelated(Number_id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Healthrelated)
	outnum, _ := orm.ID(Number_id).Delete(a)

	return int(outnum)

}

// 根据
func SelectHealthrelatedByTitle(Drugname string) (*Healthrelated, error) {
	a := new(Healthrelated)
	has, err := orm.Where("title = ?", Drugname).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}
