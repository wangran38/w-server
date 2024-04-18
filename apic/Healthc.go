package apic

import (
	// "fmt"
	// "net/http"
	_ "time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Health struct {
	Id           int64  `json:"id"`           //id
	Senior_id    int64  `json:"senior_id"`    //老者id
	Assessors_id int64  `json:"assessors_id"` //所属评估员id
	Number_id    int64  ` json:"number_id"`   //编号id
	Disease      string `json:"disease"`      //疾病诊断
	Drugname     string ` json:"drugname"`    //药物名称
	Medication   string ` json:"medication"`  //服药方法
	Dosage       string `json:"dosage"`       //用药剂量
	Frequency    string `json:"frequency"`    //用药频率
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
}

// 获取展会信息
func GetHealthlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Health
	c.ShouldBind(&searchdata)
	// //读取数据库

	search := &models.Health{
		Id:           searchdata.Id,
		Senior_id:    searchdata.Senior_id,
		Assessors_id: searchdata.Assessors_id,
		Number_id:    searchdata.Number_id,
		Disease:      searchdata.Disease,
		Drugname:     searchdata.Drugname,
		Medication:   searchdata.Medication,
		Dosage:       searchdata.Dosage,
		Frequency:    searchdata.Frequency,
	}
	listdata := models.GetHealthList(10, 1, search)
	listnum := models.GetHealthtotal(search)
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

func GetHealthList1(c *gin.Context) {
	var searchdata Health
	c.ShouldBind(&searchdata)
	// fmt.Print(searchinfo.Id)
	// result := make(map[string]interface{})
	info, _ := models.SelectHealthnumber_id(searchdata.Number_id)
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
