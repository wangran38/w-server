package controllers

import (
	// "fmt"
	"time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

// type NewsController struct{}
type Health struct {
	Id           int64 `json:"id`        //id
	Senior_id    int64 `json:"senior_id` //老者id
	Assessors_id int64 `json:"assessors_id"`

	//所属评估员id
	Number_id  int64     ` json:"number_id"`                         //编号id
	Disease    string    `xorm:"comment('疾病诊断')" json:"disease"`     //疾病诊断
	Drugname   string    ` xorm:"comment('药物名称')" json:"drugname"`   //药物名称
	Medication string    ` xorm:"comment('服药方法')" json:"medication"` //服药方法
	Dosage     string    ` xorm:"comment('用药剂量')" json:"dosage"`     //用药剂量
	Frequency  string    ` xorm:"comment('用药频率')" json:"frequency"`  //用药频率
	Created    time.Time `json:"createtime" xorm:"int"`
	Updated    time.Time `json:"updatetime" xorm:"int"`
	Limit      int       `json:"limit"`
	Page       int       `json:"page"`
}

// type Any interface{}
// 获取当前用户信息
func GetoHealthlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Health
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	search := &models.Health{
		Id:           searchdata.Id,
		Senior_id:    searchdata.Senior_id,
		Assessors_id: searchdata.Assessors_id,
		Number_id:    searchdata.Number_id,
		Disease:      searchdata.Disease,
		Drugname:     searchdata.Drugname,
		Frequency:    searchdata.Frequency,
		Created:      searchdata.Created,
		Updated:      searchdata.Updated,
	}
	listdata := models.GetHealthList(limit, page, search)
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
func AddHealth(c *gin.Context) {
	var formdata models.Health
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Health)
	Intodata.Id = formdata.Id
	Intodata.Senior_id = formdata.Senior_id
	Intodata.Assessors_id = formdata.Assessors_id
	Intodata.Number_id = formdata.Number_id
	Intodata.Disease = formdata.Disease
	Intodata.Drugname = formdata.Drugname
	Intodata.Frequency = formdata.Frequency
	Intodata.Created = time.Now()
	info, _ := models.SelectHealthnumber_id(Intodata.Number_id) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "已经存在！",
		})
		return
	}
	err := models.AddHealth(Intodata) //判断账号是否存在！
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
func EdiHealth(c *gin.Context) {
	var formdata models.Health
	c.ShouldBind(&formdata)
	updata := new(models.Health)
	updata.Id = formdata.Id
	updata.Senior_id = formdata.Senior_id
	updata.Assessors_id = formdata.Assessors_id
	updata.Number_id = formdata.Number_id
	updata.Disease = formdata.Dosage
	updata.Drugname = formdata.Drugname
	updata.Frequency = formdata.Frequency
	updata.Updated = time.Now()
	if formdata.Number_id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.UpHealth(updata) //判断账号是否存在！
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

func DeleteHealth(c *gin.Context) {
	var searchdata models.Health
	c.BindJSON(&searchdata)
	delnum := models.DeleteHealth(searchdata.Senior_id)
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
