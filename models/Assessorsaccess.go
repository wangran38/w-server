package models

import (
	"errors"
)

type Assessorsaccess struct {
	Uid int64
	Gid int64
}

func (a *Assessorsaccess) TableName() string {
	return "Assessors_group_access"
}

// 根据用户id找用户返回数据
func SelectGid(Id int64) (*Assessorsaccess, error) {
	a := new(Assessorsaccess)
	has, err := orm.Where("uid = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到所在组别！")
	}
	return a, nil

}
func AddAssessorsaccess(a *Assessorsaccess) error {
	_, err := orm.Insert(a)
	return err
}

// 删除
func DeleteAssessorsaccess(uid int64) int {
	a := new(Assessorsaccess)
	outnum, _ := orm.Where("uid = ?", uid).Delete(a)
	return int(outnum)

}

// 修改
// 修改
func UpAssessorsaccess(a *Assessorsaccess) (int64, error) {
	affected, err := orm.Where("uid = ?", a.Uid).Update(a)
	return affected, err

}
