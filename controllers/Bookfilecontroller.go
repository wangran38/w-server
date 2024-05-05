package controllers

import (
	"time"
	"w-server/models"

	"github.com/gin-gonic/gin"
)

type Bookfilec struct {
	Id      int64  `json:"id"`                                            // ID 自增长
	Bookid  int64  `xorm:"'book_id' comment('所属章节')" json:"book_id"`      // ID 自增长
	Fileurl string `xorm:"varchar(250) comment('图片视频地址')" json:"fileurl"` // 章节名称
	Title   string `xorm:"varchar(150) comment('文件名称')" json:"title"`     // 事件名称
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Order   string `json:"sort"`
}

// 获取当前用户信息
func GetBookfilelist(c *gin.Context) {
	//从header中获取到token
	var searchdata Bookfilec
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// kname := searchdata.Kname
	order := searchdata.Order
	search := &models.Bookfile{
		// Id:    searchdata.Id,
		Id:      searchdata.Id,
		Bookid:  searchdata.Bookid,
		Fileurl: searchdata.Fileurl,
		Title:   searchdata.Title,
	}
	listdata := models.GetbookfileList(limit, page, search, order)
	listnum := models.GetbookfileNum(search)

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

// 添加用户组
func AddBookfile(c *gin.Context) {
	var formdata models.Bookfile
	c.ShouldBind(&formdata)
	Rulesdata := new(models.Bookfile)

	Rulesdata.Bookid = formdata.Bookid
	Rulesdata.Title = formdata.Title
	Rulesdata.Fileurl = formdata.Fileurl
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
	err := models.Addbookfile(Rulesdata) //判断账号是否存在！
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
func EditBookfile(c *gin.Context) {
	var formdata models.Bookfile
	c.ShouldBind(&formdata)
	intodata := new(models.Bookfile)
	intodata.Id = formdata.Id
	intodata.Bookid = formdata.Bookid
	intodata.Fileurl = formdata.Fileurl
	intodata.Title = formdata.Title
	intodata.Created = time.Now()
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
		res, err := models.Upbookfile(intodata) //判断
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

func DelBookfile(c *gin.Context) {
	var searchdata models.Bookfile
	c.BindJSON(&searchdata)
	delnum := models.Deletebookfile(searchdata.Id)
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
