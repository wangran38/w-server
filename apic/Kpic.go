package apic

import (
	// "fmt"
	_ "time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type KpiJserch struct {
	Id    int64  `json:"id"`
	Pid   int64  `json:"pid"`
	Kname string `json:"kname"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}
type Kpiinfojson struct {
	Id       int64  `json:"id"`
	Pid      int64  `json:"pid"`      //ID自增长
	Kname    string `json:"kname"`    //指标名称
	Level    int    `json:"level"`    //指标等级
	Kcontent string `json:"kcontent"` //指标的描述内容
	// Created  time.Time `json:"createtime"` //创建时间戳
	// Updated  time.Time `json:"updatetime"` //更改时间戳
	Weigh    int `json:"weigh"`  //排序，倒序，
	Status   int `json:"status"` //状态
	Children []*models.Kpiinfo
}

// type Any interface{}
// func Getcategorytree(c *gin.Context) {
// 	listdata := models.Getcategorytree(0)
// 	result := make(map[string]interface{})
// 	if listdata == nil {
// 		c.JSON(200, gin.H{
// 			"code":    201,
// 			"message": "获取分类失败",
// 			"data":    "",
// 		})
// 		return
// 	} else {
// 		result["listdata"] = listdata
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"message": "数据获取成功",
// 			"data":    result,
// 		})
// 		return
// 	}
// }

// 获取当前用户信息
func Getkpilist(c *gin.Context) {
	//从header中获取到token
	var searchdata KpiJserch
	c.ShouldBind(&searchdata)
	// //读取数据库

	// name:=""
	limit := searchdata.Limit
	if limit < 1 {
		limit = 10
	}
	page := searchdata.Page
	if page < 1 {
		page = 1
	}
	// title := searchdata.Title
	order := searchdata.Order
	search := &models.Kpi{
		Id:    searchdata.Id,
		Pid:   searchdata.Pid,
		Kname: searchdata.Kname,
	}
	listdata := models.GetkpiapiList(limit, page, search, order)
	treelist := []*Kpiinfojson{}
	for _, v := range listdata { //循环range切片
		search1 := &models.Kpiinfo{
			// Id:    searchdata.Id,
			Kpiid: v.Id,
		}

		infolistdata := models.GetkpiinfoList(10, 1, search1, "")
		node := &Kpiinfojson{
			Id:       v.Id,
			Pid:      v.Pid,
			Kname:    v.Kname,
			Level:    v.Level,
			Kcontent: v.Kcontent,
			// Level:  v.Level,
			Weigh: v.Weigh,
		}
		node.Children = infolistdata
		treelist = append(treelist, node)
		// fmt.Printf("%s -> %s\n", k, v.Kname)
	}
	listnum := models.Getkpitotal(search)
	result := make(map[string]interface{})
	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "当前数据库为空",
			"data":    "",
		})
		return
	} else {
		result["listdata"] = treelist
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data":    result,
		})
		return
	}
}

//获取全部上下级
