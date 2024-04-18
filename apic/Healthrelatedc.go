package apic

import (
	// "fmt"
	// "net/http"
	_ "time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Healthrelated struct {
	Id             int64  `json:"id"`               //id
	Senior_id      int64  `json:"senior_id"`        //老者id
	Assessors_id   int64  `json:"assessors_id"`     //所属评估员id
	Number_id      int64  ` json:"number_id"`       //编号id
	Pressureinjury string `json:"pressureinjury"`   //压力性损伤
	Joint          string ` json:"Joint"`           //关节活动度
	Affectedarea   string ` json:"Affectedarea"`    //关节影响部位
	Woundcondition string `  json:"woundcondition"` //伤口情况
	Specialcare    string ` json:"specialcare"`     //特殊护理
	Painsensation  string `json:"painsensation"`    //疼痛感
	Toothloss      string ` json:"toothloss"`       //牙齿缺失
	Wearing        string `  json:"wearing"`        //义齿佩戴
	Swallow        string `  json:"swallow"`        //吞咽困难的情形和状况
	Malnutrition   string `  json:"malnutrition"`   //营养不良
	Cleaning       string `  json:"cleaning"`       //清理呼吸道无效
	Coma           string ` json:"coma"`            //昏迷
	Other          string ` json:"other"`           //其他
	Limit          int    `json:"limit"`
	Page           int    `json:"page"`
}

// 获取信息
func GetHealthrelatedlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Healthrelated
	c.ShouldBind(&searchdata)
	// //读取数据库

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
	}
	listdata := models.GetHealthrelatedList(10, 1, search)
	listnum := models.GetHealthrelatedtotal(search)
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

func GetHealthrelatedList1(c *gin.Context) {
	var searchdata Healthrelated
	c.ShouldBind(&searchdata)
	// fmt.Print(searchinfo.Id)
	// result := make(map[string]interface{})
	info, _ := models.SelectHealthrelatednumber_id(searchdata.Number_id)
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
