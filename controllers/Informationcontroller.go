package controllers

import (
	// "fmt"
	"time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

// type NewsController struct{}
type Information struct {
	Id           int64     `json:"id"`
	Senior_id    int64     `json:"senior_id"`
	Assessors_id int64     `json:"assessors_id"`
	Provider     string    `json:"provider"`
	Relationship string    `json:"relationship" xorm:"varchar(200)"`
	Contactname  string    `json:"contactname" xorm:"TEXT "`
	Phone        string    `json:"phone" xorm:"varchar(200)"`
	Created      time.Time `json:"createtime" xorm:"int"`
	Updated      time.Time `json:"updatetime" xorm:"int"`
	Limit        int       `json:"limit"`
	Page         int       `json:"page"`
}

// type Any interface{}
// 获取当前用户信息
func GetInformationlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Information
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	search := &models.Information{
		Id:           searchdata.Id,
		Senior_id:    searchdata.Senior_id,
		Assessors_id: searchdata.Assessors_id,
		Provider:     searchdata.Provider,
		Relationship: searchdata.Relationship,
		Contactname:  searchdata.Contactname,
		Phone:        searchdata.Phone,
		Created:      searchdata.Created,
		Updated:      searchdata.Updated,
	}
	listdata := models.GetInformationList(limit, page, search)
	listnum := models.GetInformationtotal(search)

	result["page"] = page
	result["totalnum"] = listnum
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
func AddInformation(c *gin.Context) {
	var formdata models.Information
	// 检查 phone 字段是否为 11 位数字
	if len(formdata.Phone) != 11 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "无效的电话号码。电话号码应为 11 位数字。",
			"data": "",
		})
		return
	}
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Information)
	Intodata.Id = formdata.Id
	Intodata.Senior_id = formdata.Senior_id
	Intodata.Assessors_id = formdata.Assessors_id
	Intodata.Provider = formdata.Provider
	Intodata.Relationship = formdata.Relationship
	Intodata.Contactname = formdata.Contactname
	Intodata.Phone = formdata.Phone
	Intodata.Created = time.Now()
	info, _ := models.SelectInformationid(Intodata.Id) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "已经存在！",
		})
		return
	}
	err := models.AddInformation(Intodata) //判断账号是否存在！
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
func EdiInformation(c *gin.Context) {
	var formdata models.Information
	// 检查 phone 字段是否为 11 位数字
	if len(formdata.Phone) != 11 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "无效的电话号码。电话号码应为 11 位数字。",
		})
	}
	c.ShouldBind(&formdata)
	updata := new(models.Information)
	updata.Id = formdata.Id
	updata.Senior_id = formdata.Senior_id
	updata.Assessors_id = formdata.Assessors_id
	updata.Contactname = formdata.Contactname
	updata.Provider = formdata.Provider
	updata.Relationship = formdata.Relationship
	updata.Phone = formdata.Phone
	updata.Updated = time.Now()
	if formdata.Senior_id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.UpInformation(updata) //判断账号是否存在！
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

func DelInformation(c *gin.Context) {
	var searchdata models.Information
	c.BindJSON(&searchdata)
	delnum := models.DeleteInformation(searchdata.Senior_id)
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
