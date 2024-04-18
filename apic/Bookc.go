package apic

import (
	// "fmt"
	_ "time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type BookJserch struct {
	Id          int64  `json:"id"`          // ID 自增长
	Pid         int64  `json:"pid"`         // ID 自增长
	Chaptername string `json:"chaptername"` // 章节名称
	Level       int    `json:"level"`       //级别
	Name        string `json:"name"`        // 事件名称
	Content     string `json:"content"`     // 哪本书
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Order       string `json:"sort"`
}
type Bookinfojson struct {
	Id          int64  `json:"id"`          // ID 自增长
	Pid         int64  `json:"pid"`         // ID 自增长
	Chaptername string `json:"chaptername"` // 章节名称
	Level       int    `json:"level"`       //级别
	Name        string `json:"name"`        // 事件名称
	Content     string `json:"content"`     // 哪本书
	// Created  time.Time `json:"createtime"` //创建时间戳
	// Updated  time.Time `json:"updatetime"` //更改时间戳
	Weigh    int `json:"weigh"`  //排序，倒序，
	Status   int `json:"status"` //状态
	Children []*models.Bookinfo
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
func Getbooklist(c *gin.Context) {
	//从header中获取到token
	var searchdata BookJserch
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
	search := &models.Book{
		Id:          searchdata.Id,
		Pid:         searchdata.Pid,
		Chaptername: searchdata.Chaptername,
		Name:        searchdata.Name,
		Content:     searchdata.Content,
	}
	listdata := models.GetbookapiList(limit, page, search, order)
	//treelist := []*Bookinfojson{}
	// for _, v := range listdata { //循环range切片
	// 	// search1 := &models.Bookinfo{
	// 	// 	// Id:    searchdata.Id,
	// 	// 	Kpiid: v.Id,
	// 	// }

	// 	// infolistdata := models.GetbookinfoList(10, 1, search1, "")
	// 	// node := &Bookinfojson{
	// 	// 	Id:          v.Id,
	// 	// 	Pid:         v.Pid,
	// 	// 	Chaptername: v.Chaptername,
	// 	// 	Level:       v.Level,
	// 	// 	Name:        v.Name,
	// 	// 	Content:     Content,
	// 	// 	// Level:  v.Level,
	// 	// 	// Weigh: v.Weigh,
	// 	// }
	// 	// node.Children = infolistdata
	// 	// treelist = append(treelist, node)
	// 	// fmt.Printf("%s -> %s\n", k, v.Kname)
	// }
	// listnum := models.Getbooktotal(search)
	result := make(map[string]interface{})
	result["page"] = page
	// result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "当前数据库为空",
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
