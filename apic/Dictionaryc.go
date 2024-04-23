package apic

import (
	// "fmt"
	// "net/http"
	_ "time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Dictionary struct {
	Id       int64  `json:"id"`
	Code     int64  `json:"code"`
	Codename string `json:"codename"`
	Name     string `json:"name"`
}

// 获取字典信息
func GetDictionaryclist(c *gin.Context) {
	//从header中获取到token
	var searchdata Dictionary
	c.ShouldBind(&searchdata)
	// //读取数据库

	// name:=""
	search := &models.Dictionary{
		Id:       searchdata.Id,
		Code:     searchdata.Code,
		Codename: searchdata.Codename,
		Name:     searchdata.Name,
	}
	listdata := models.SelectDictionarylist(10, 1, search)
	// listnum := models.GetNewstotal(search)
	result := make(map[string]interface{})
	// result["page"] = page
	// result["totalnum"] = listnum
	// result["limit"] = limit
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

// AddDoctionary 用于添加用户信息
// func AddDictionary(c *gin.Context) {
// 	// 从请求头中获取令牌
// 	// token := c.Request.Header.Get("Authorization")
// 	// if token == "" || len(token) == 0 {
// 	// 	c.JSON(201, gin.H{
// 	// 		"code":    201,
// 	// 		"message": "你没有权限,去远处玩！",
// 	// 		"data":    "",
// 	// 		// "permissions": menu,
// 	// 		// "roles":       role,
// 	// 	})
// 	// 	return
// 	// }
// 	// user, tokenerr := utils.GetLoginAssessorsc(token)
// 	// if tokenerr != nil {
// 	// 	c.JSON(201, gin.H{
// 	// 		"code":    201,
// 	// 		"message": "登录失效，请重新登录！",
// 	// 		"data":    "",
// 	// 		// "permissions": menu,
// 	// 		// "roles":       role,
// 	// 	})
// 	// 	return
// 	}
// 	var formdata Dictionary
// 	c.ShouldBind(&formdata)
// 	Intodata := new(models.Dictionary)
// 	Intodata.Id = formdata.Id
// 	Intodata.Code = formdata.Code
// 	Intodata.Codename = formdata.Codename
// 	Intodata.Name = formdata.Name

// 	err := models.AddDictionary(Intodata) // 判断账号是否存在！
// 	if err != nil {
// 		c.JSON(201, gin.H{
// 			"code": 201,
// 			"msg":  "添加数据出错！",
// 			"data": err,
// 		})
// 		return
// 	} else {
// 		c.JSON(200, gin.H{
// 			"code": 200,
// 			"msg":  "数据添加成功！",
// 			"data": "",
// 		})
// 	}
// }
