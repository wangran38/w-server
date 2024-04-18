package controllers

import (
	"time"
	"w-server/models"

	"github.com/gin-gonic/gin"
)

type Bookinfoserch struct {
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

// 获取当前用户信息
func GetBookinfolist(c *gin.Context) {
	//从header中获取到token
	var searchdata Bookinfoserch
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// kname := searchdata.Kname
	order := searchdata.Order
	search := &models.Bookinfo{
		// Id:    searchdata.Id,
		Pid:         searchdata.Pid,
		Chaptername: searchdata.Chaptername,
		Level:       searchdata.Level,
		Name:        searchdata.Name,
		Content:     searchdata.Content,
	}
	listdata := models.GetbookinfoList(limit, page, search, order)
	listnum := models.GetbookinfoNum(search)

	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取数据为空",
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
func Addbookinfo(c *gin.Context) {
	var formdata models.Bookinfo
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Rulesdata := new(models.Bookinfo)

	Rulesdata.Pid = formdata.Pid
	Rulesdata.Chaptername = formdata.Chaptername
	Rulesdata.Content = formdata.Content
	Rulesdata.Level = formdata.Level
	Rulesdata.Name = formdata.Name
	// Rulesdata.Kcontent = formdata.Kcontent
	Rulesdata.Created = time.Now()
	// info, _ := models.SelectkpiByTitle(Rulesdata.Kname) //判断账号是否存在！
	// if info != nil {
	// 	c.JSON(200, gin.H{
	// 		"code": "201",
	// 		"msg":  "该指标已经存在！",
	// 	})
	// 	return
	// }
	err := models.Addbookinfo(Rulesdata) //判断账号是否存在！
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
func EditBookinfo(c *gin.Context) {
	var formdata models.Bookinfo
	c.ShouldBind(&formdata)
	intodata := new(models.Bookinfo)
	intodata.Id = formdata.Id
	intodata.Pid = formdata.Pid
	intodata.Chaptername = formdata.Chaptername
	intodata.Name = formdata.Name
	intodata.Content = formdata.Content
	// intodata.Weigh = formdata.Weigh
	// intodata.Kcontent = formdata.Kcontent
	if formdata.Id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.Upbookinfo(intodata) //判断
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

func DelBookinfo(c *gin.Context) {
	var searchdata models.Bookinfo
	c.BindJSON(&searchdata)
	delnum := models.Deletebookinfo(searchdata.Id)
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
