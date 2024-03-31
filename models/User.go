package models
import (
	"errors"
	"time"
)
type User struct {
	Id   int64
	Openid string `xorm:"varchar(200)" json:"openid"`
	Nickname string `xorm:"varchar(200)" json:"nickname"`
	Sex int `json:"sex" xorm:"not null default 1 comment('是否启用 默认1 是 0 无') TINYINT"`
	Province string `xorm:"varchar(200)" json:"province"`
	City string `xorm:"varchar(200)" json:"city"`
	Country string `xorm:"varchar(200)" json:"country"`
	Headimgurl string `xorm:"varchar(200)" json:"headimgurl"`
	Unionid string `xorm:"varchar(200)" json:"unionid"`
	Token string `xorm:"varchar(200)" json:"token"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func (a *User) TableName() string {
	return "user"
}
func AddUser(a *User) error {
	_, err := orm.Insert(a)
	return err
}
//修改
func UpUser(a *User) (int64,error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}
//根据
func SelectUserByOpenid(Openid string) (*User, error) {
	a := new(User)
	has, err := orm.Where("openid = ?", Openid).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}