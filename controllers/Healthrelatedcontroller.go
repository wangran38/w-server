package controllers

import (
	// "fmt"
	"time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

// type NewsController struct{}
type Healthrelated struct {
	Id             int64     `json:"id"`                                     //id
	Senior_id      int64     `xorm:"comment('老者id')" json:"senior_id"`       //老者id
	Assessors_id   int64     `xorm:"comment('所属评估员id')" json:"assessors_id"` //所属评估员id
	Number_id      int64     `xorm:"comment('编号id')" json:"number_id"`       //编号id
	Pressureinjury string    `xorm:"comment('压力性损伤')" json:"pressureinjury"` //压力性损伤
	Joint          string    ` xorm:"comment('关节活动度度')" json:"Joint"`        //关节活动度
	Affectedarea   string    ` xorm:"comment('关节影响部位')" json:"affectedarea"` //关节影响部位
	Woundcondition string    ` xorm:"comment('伤口情况')" json:"woundcondition"` //伤口情况
	Specialcare    string    ` xorm:"comment('用药剂量')" json:"specialcare"`    //特殊护理
	Painsensation  string    ` xorm:"comment('疼痛感')" json:"painsensation"`   //疼痛感
	Toothloss      string    ` xorm:"comment('牙齿缺失')" json:"toothloss"`      //牙齿缺失
	Wearing        string    ` xorm:"comment('义齿佩戴')" json:"wearing"`        //义齿佩戴
	Swallow        string    ` xorm:"comment('吞咽困难的情形和症状')" json:"swallow"`  //吞咽困难的情形和状况
	Malnutrition   string    ` xorm:"comment('营养不良')" json:"malnutrition"`   //营养不良
	Cleaning       string    ` xorm:"comment('清理呼吸道无效')" json:"cleaning"`    //清理呼吸道无效
	Coma           string    ` xorm:"comment('昏迷')" json:"coma"`             //昏迷
	Other          string    ` xorm:"comment('其他')" json:"other"`            //其他
	Created        time.Time `xorm:"int" json:"createtime"`
	Updated        time.Time `xorm:"int" json:"updatetime"`
	Limit          int       `json:"limit"`
	Page           int       `json:"page"`
}

// type Any interface{}
// 获取当前用户信息
func GetHealthrelatedlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Healthrelated
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	search := &models.Healthrelated{
		Id:             searchdata.Id,
		Senior_id:      searchdata.Senior_id,
		Assessors_id:   searchdata.Assessors_id,
		Number_id:      searchdata.Number_id,
		Pressureinjury: searchdata.Pressureinjury,
		Joint:          searchdata.Joint,
		Affectedarea:   searchdata.Affectedarea,
		Woundcondition: searchdata.Woundcondition,
		Specialcare:    searchdata.Specialcare,
		Painsensation:  searchdata.Painsensation,
		Toothloss:      searchdata.Toothloss,
		Wearing:        searchdata.Wearing,
		Swallow:        searchdata.Swallow,
		Malnutrition:   searchdata.Malnutrition,
		Cleaning:       searchdata.Cleaning,
		Coma:           searchdata.Coma,
		Other:          searchdata.Other,
		// Created:        searchdata.Created,
		// Updated:        searchdata.Updated,
	}
	listdata := models.GetHealthrelatedList(limit, page, search)
	// listnum := models.GetHealthList(search)

	result["page"] = page
	// result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取菜单失败1",
			"data":    "",
		})
		return
	} else {
		result["listdata"] = listdata
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data":    result,
		})
		return
	}
}

// //添加用户组
func AddHealthrelated(c *gin.Context) {
	var formdata models.Healthrelated
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Healthrelated)
	Intodata.Id = formdata.Id
	Intodata.Senior_id = formdata.Senior_id
	Intodata.Assessors_id = formdata.Assessors_id
	Intodata.Number_id = formdata.Number_id
	Intodata.Pressureinjury = formdata.Pressureinjury
	Intodata.Joint = formdata.Joint
	Intodata.Affectedarea = formdata.Affectedarea
	Intodata.Woundcondition = formdata.Woundcondition
	Intodata.Specialcare = formdata.Specialcare
	Intodata.Painsensation = formdata.Painsensation
	Intodata.Toothloss = formdata.Toothloss
	Intodata.Wearing = formdata.Wearing
	Intodata.Swallow = formdata.Swallow
	Intodata.Malnutrition = formdata.Woundcondition
	Intodata.Cleaning = formdata.Cleaning
	Intodata.Coma = formdata.Coma
	Intodata.Other = formdata.Other
	Intodata.Created = time.Now()
	info, _ := models.SelectHealthrelatedid(Intodata.Number_id) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "已经存在！",
		})
		return
	}
	err := models.AddHealthrelated(Intodata) //判断账号是否存在！
	if err != nil {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "添加数据出错！",
			"data": err,
		})
		return
	} else {
		// result := make(map[string]interface{})
		// result["id"] = Rulestable.Id //返回当前总数
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "数据添加成功！",
			"data": "",
		})
	}
}

// //修改用户组
func EdiHealthrelated(c *gin.Context) {
	var formdata models.Healthrelated
	c.ShouldBind(&formdata)
	updata := new(models.Healthrelated)
	updata.Id = formdata.Id
	updata.Senior_id = formdata.Senior_id
	updata.Assessors_id = formdata.Assessors_id
	updata.Number_id = formdata.Number_id
	updata.Pressureinjury = formdata.Pressureinjury
	updata.Joint = formdata.Joint
	updata.Affectedarea = formdata.Affectedarea
	updata.Woundcondition = formdata.Woundcondition
	updata.Specialcare = formdata.Specialcare
	updata.Painsensation = formdata.Painsensation
	updata.Toothloss = formdata.Toothloss
	updata.Wearing = formdata.Wearing
	updata.Swallow = formdata.Swallow
	updata.Malnutrition = formdata.Woundcondition
	updata.Cleaning = formdata.Cleaning
	updata.Coma = formdata.Coma
	updata.Other = formdata.Other
	updata.Updated = time.Now()
	if formdata.Id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.UpHealthrelated(updata) //判断账号是否存在！
		if err != nil {
			c.JSON(201, gin.H{
				"code": 201,
				"msg":  "修改数据出错！",
				"data": err,
			})
			return
		} else {

			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "数据修改成功！",
				"data": res,
			})

		}
	}
}

func DeleteHealthrelated(c *gin.Context) {
	var searchdata models.Healthrelated
	c.BindJSON(&searchdata)
	delnum := models.DeleteHealthrelated(searchdata.Id)
	if delnum > 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除成功！",
			"data":    delnum,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "操作失败！",
		})

	}

}
