package models

import "errors"

// 老人基础信息表
type Olduser struct {
	Id                      int64  //ID自增长
	Name                    string `xorm:"varchar(64) comment('姓名')" json:"name"`                                             //姓名
	Sex                     int    `json:"sex" xorm:"not null default 1 comment('是否启用 默认1 是男 2 女') TINYINT"`                  //性别
	Nation                  string `xorm:"varchar(64)" json:"nation"`                                                         //民族
	Birth                   string `xorm:"varchar(20)" json:"birrth"`                                                         //出生日期
	Marriage                string `xorm:"varchar(50)" json:"marriage"`                                                       //婚姻状况
	Marriagecode            string `xorm:"varchar(50)" json:"marriagecode"`                                                   //婚姻状况编码
	Idnumber                string `xorm:"varchar(64)" json:"idnumber"`                                                       //身份证号
	Ssnumber                string `xorm:"varchar(64)" json:"ssnumber"`                                                       //社保号码
	Financial               string `xorm:"varchar(32)" json:"financial"`                                                      //经济来源
	Payment                 string `xorm:"varchar(32)" json:"payment"`                                                        //医疗费用支付方式
	Kind                    int    `xorm:"int(11)" json:"Kind"`                                                               //类别（居家，机构）
	Status                  int    `xorm:"int(11)" json:"status"`                                                             //状态
	Weiglatest_eval_gradeht int    `json:"weiglatest_eval_gradeht" xorm:"not null default 1 comment('层级1 2 3 4 等级') TINYINT"` //最近评级机构
	D1                      string `xorm:"varchar(50)" json:"d1"`                                                             //省
	D2                      string `xorm:"varchar(50)" json:"d2"`                                                             //市
	D3                      string `xorm:"varchar(50)" json:"d3"`                                                             //县
	D4                      string `xorm:"varchar(50)" json:"d4"`                                                             //乡
	Address                 string `xorm:"varchar(50)" json:"address"`                                                        //地址
	H1                      string `xorm:"varchar(50)" json:"h1"`                                                             //省（户籍地）
	H2                      string `xorm:"varchar(50)" json:"h2"`                                                             //市（户籍地）
	H3                      string `xorm:"varchar(50)" json:"h3"`                                                             //县（户籍地）
	H4                      string `xorm:"varchar(50)" json:"h4"`                                                             //乡（户籍地）
	Place                   string `xorm:"varchar(50)" json:"place"`                                                          //地址（户籍地）
	Previous_work           string `xorm:"varchar(50)" json:"Previous_work"`                                                  //既往工作
	Education_code          string `xorm:"varchar(50)" json:"education_code"`                                                 //教育程度代码
	Education_name          string `xorm:"varchar(50)" json:"education_name"`                                                 //教育程度名称
	Residency_code          string `xorm:"varchar(50)" json:"residency_code"`                                                 //居住状况代码
	Residency_name          string `xorm:"varchar(50)" json:"residency_name"`                                                 //居住状况名称
	Live_with_who_code      string `xorm:"varchar(50)" json:"live_with_who_code"`                                             //与谁住在一起
	Live_with_who_name      string `xorm:"varchar(50)" json:"live_with_who_name"`                                             //与谁住在一起
	In_time                 int64  `xorm:"int(11)" json:"in_time"`                                                            //入院时间
	Room_num                string `xorm:"varchar(50)" json:"room_num"`                                                       //房号
}

func (a *Olduser) TableName() string {
	return "elder" //数据库表的别名
}

func AddOlduser(a *Olduser) error {
	_, err := orm.Insert(a)
	return err
}

// //修改
// func UpUser(a *User) (int64,error) {
// 	affected, err := orm.Id(a.Id).Update(a)
// 	return affected, err

// }
// //根据
func SelectOlduserbyname(Name string) (*Olduser, error) {
	a := new(Olduser)
	has, err := orm.Where("name = ?", Name).Get(a)
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
func Deleteold(id int64) int {
	a := new(Olduser)
	outnum, _ := orm.Where("id = ?", id).Delete(a)
	return int(outnum)

}
