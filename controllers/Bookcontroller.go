package controllers

import (
	// "fmt"
	"time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Bookserch struct {
	Id          int64     `json:"id"`          // ID 自增长
	Pid         int64     `json:"pid"`         // ID 自增长
	Chaptername string    `json:"chaptername"` // 章节名称
	Level       int       `json:"level"`       //级别
	Name        string    `json:"name"`        // 事件名称
	Content     string    `json:"content"`     // 哪本书
	Created     time.Time `json:"createtime"`  // 创建时间戳
	Updated     time.Time `json:"updatetime"`  // 更改时间戳
	Limit       int       `json:"limit"`
	Page        int       `json:"page"`
	Order       string    `json:"sort"`
}

// type Any interface{}
// 获取当前用户信息
func GetBooklist(c *gin.Context) {
	//从header中获取到token
	var searchdata Bookserch
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// kname := searchdata.Kname
	order := searchdata.Order
	search := &models.Book{
		// Id:    searchdata.Id,
		Pid:         searchdata.Pid,
		Chaptername: searchdata.Chaptername,
		Level:       searchdata.Level,
	}
	listdata := models.GetbookapiList(limit, page, search, order)

	listnum := models.Getbooktotal(search)

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

// 获取全部上下级
func TreeBook(c *gin.Context) {

	grouplist := models.GetBooktree(0)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "数据获取成功",
		"data":    grouplist,
	})

}

// //添加用户组
func Addbook(c *gin.Context) {
	var formdata models.Book
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Rulesdata := new(models.Book)

	Rulesdata.Pid = formdata.Pid
	Rulesdata.Chaptername = formdata.Chaptername
	Rulesdata.Level = formdata.Level
	Rulesdata.Name = formdata.Name
	Rulesdata.Content = formdata.Content
	Rulesdata.Created = time.Now()
	info, _ := models.SelectbookByTitle(Rulesdata.Chaptername) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "该指标已经存在！",
		})
		return
	}
	err := models.Addbook(Rulesdata) //判断账号是否存在！
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
func EditBook(c *gin.Context) {
	var formdata models.Book
	c.ShouldBind(&formdata)
	intodata := new(models.Book)
	intodata.Id = formdata.Id
	intodata.Pid = formdata.Pid
	intodata.Chaptername = formdata.Chaptername
	intodata.Level = formdata.Level
	intodata.Name = formdata.Name
	intodata.Content = formdata.Content
	if formdata.Id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.Upbook(intodata) //判断
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

func DelBook(c *gin.Context) {
	var searchdata models.Book
	c.BindJSON(&searchdata)
	delnum := models.Deletebook(searchdata.Id)
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
