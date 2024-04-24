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

// AddInformation 用于添加用户信息
// func AddBook(c *gin.Context) {
// 	// 从请求头中获取令牌
// 	token := c.Request.Header.Get("Authorization")
// 	if token == "" || len(token) == 0 {
// 		c.JSON(201, gin.H{
// 			"code":    201,
// 			"message": "你没有权限,去远处玩！",
// 			"data":    "",
// 			// "permissions": menu,
// 			// "roles":       role,
// 		})
// 		return
// 	}
// 	user, tokenerr := utils.GetLoginAssessorsc(token)
// 	if tokenerr != nil {
// 		c.JSON(201, gin.H{
// 			"code":    201,
// 			"message": "登录失效，请重新登录！",
// 			"data":    "",
// 			// "permissions": menu,
// 			// "roles":       role,
// 		})
// 		return
// 	}
// 	var formdata models.Book
// 	c.ShouldBind(&formdata)
// 	Intodata := new(models.Book)
// 	// Intodata.Id = formdata.Id
// 	Intodata.Pid = formdata.Pid
// 	Intodata.Chaptername = formdata.Chaptername
// 	Intodata.Level = formdata.Level
// 	Intodata.Name = formdata.Name
// 	Intodata.Content = formdata.Content
// 	Intodata.Isdel = formdata.Isdel
// 	err := models.AddBook(Intodata) // 判断账号是否存在！
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
