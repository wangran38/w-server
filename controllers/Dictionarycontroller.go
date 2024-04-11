package controllers

import (
	// "fmt"
	"w-server/models"
	// "time"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

// type NewsController struct{}
type Dictionary struct {
	Id       int64  `json:"id"`
	Code     int64  `json:"code"`
	Codename string `json:"codename" xorm:"varchar(200)"`
	Name     string `json:"name" xorm:"TEXT "`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}

// type Any interface{}
// 获取当前用户信息
func GetDictionarylist(c *gin.Context) {
	//从header中获取到token
	var searchdata Dictionary
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// order := searchdata.Order
	search := &models.Dictionary{
		Id:       searchdata.Id,
		Code:     searchdata.Code,
		Codename: searchdata.Codename,
		// Isshow:   searchdata.Isshow,
	}
	listdata := models.SelectDictionarylist(limit, page, search)
	listnum := models.GetDictionarytotal(search)

	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取数据失败",
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
func AddDictionary(c *gin.Context) {
	var formdata models.Dictionary
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Dictionary)

	Intodata.Id = formdata.Id
	Intodata.Code = formdata.Code
	Intodata.Codename = formdata.Codename
	Intodata.Name = formdata.Name
	// Intodata.Createtime = time.Now()
	info, _ := models.SelectDictionarybycode(Intodata.Id) //判断账号是否存在！(Intodata.Codename)
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "该分类已经存在！",
		})
		return
	}
	err := models.AddDictionary(Intodata) //判断账号是否存在！
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
func Updedictionary(c *gin.Context) {
	var formdata models.Dictionary
	c.ShouldBind(&formdata)
	updata := new(models.Dictionary)
	updata.Id = formdata.Id
	updata.Code = formdata.Code
	updata.Codename = formdata.Codename
	updata.Name = formdata.Name
	if formdata.Id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.Updedictionary(updata) //判断账号是否存在！
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

// 删除
func DelDictionary(c *gin.Context) {
	var searchdata models.Dictionary
	c.BindJSON(&searchdata)
	delnum := models.DeleteDictionary(searchdata.Id)
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
