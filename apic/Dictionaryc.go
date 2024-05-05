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
