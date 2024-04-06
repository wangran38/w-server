package models

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 建结构体
type Number struct {
	Id          int64     `json:"id"`                          //ID自增长
	Code        string    `xorm:"comment('编号')" json:"code"`   //编号
	Codetime    string    `xorm:"comment('编号拼音')" json:"date"` //评估日期
	Reasons     string    `xorm:"reasons" json:"reasons"`      //评估原因
	Assessorsid int64     `xorm:"assessors_id" json:"assessors_id"`
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
}

func (a *Number) TableName() string {
	return "number" //数据库表的名
}
func AddNumber(a *Number) error {
	_, err := orm.Insert(a)

	return err
}

// 根据用户id找用户返回数据
func SelectNumberCode(code string) (*Number, error) {
	a := new(Number)
	has, err := orm.Where("code = ?", code).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("数据出错！")
	}
	return a, nil

}
