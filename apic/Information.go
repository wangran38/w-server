package apic

import (
	// "fmt"
	// "net/http"
	_ "time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Information struct {
	Id           int64  `json:"Id"`
	Senior_id    int64  `json:"senior_id"`
	Provider     string `json:"provider"`
	Assessors_id int64  `json:"assessors_id"`
	Relationship string `json:"relationship" xorm:"varchar(200)"`
	Contactname  string `json:"contactname" xorm:"TEXT "`
	Phone        string `json:"phone" xorm:"TEXT "`
}

// 获取展会信息
func GetInformationlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Information
	c.ShouldBind(&searchdata)
	// //读取数据库

	search := &models.Information{
		Id:           searchdata.Id,
		Senior_id:    searchdata.Senior_id,
		Assessors_id: searchdata.Assessors_id,
		Provider:     searchdata.Provider,
		Relationship: searchdata.Relationship,
		Contactname:  searchdata.Contactname,
		Phone:        searchdata.Phone,
	}
	listdata := models.GetInformationList(10, 1, search)
	listnum := models.GetInformationtotal(search)
	result := make(map[string]interface{})
	result["totalnum"] = listnum
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

func GetInformationinfo(c *gin.Context) {
	var searchdata Information
	c.ShouldBind(&searchdata)
	// fmt.Print(searchinfo.Id)
	// result := make(map[string]interface{})
	info, _ := models.SelectNewsid(searchdata.Senior_id)
	if info != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "获取成功",
			"data": info,
		})

	} else {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "获取数据失败",
			"data": searchdata.Senior_id,
		})
	}
}
