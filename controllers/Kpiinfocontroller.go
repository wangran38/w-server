package controllers

import (
	"changxiaoyang/models"
	"time"

	"github.com/gin-gonic/gin"
)

type Kpiinfoserch struct {
	Id    int64  `json:"id"`
	Kpiid int64  `json:"kpi_id"`
	Title string `json:"title"`
	Score int    `json:"score"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}

// 获取当前用户信息
func GetKpiinfolist(c *gin.Context) {
	//从header中获取到token
	var searchdata Kpiinfoserch
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// kname := searchdata.Kname
	order := searchdata.Order
	search := &models.Kpiinfo{
		// Id:    searchdata.Id,
		Kpiid: searchdata.Kpiid,
		Title: searchdata.Title,
		Score: searchdata.Score,
	}
	listdata := models.GetkpiinfoList(limit, page, search, order)
	listnum := models.GetkpiinfoNum(search)

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
func AddKpiinfo(c *gin.Context) {
	var formdata models.Kpiinfo
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Rulesdata := new(models.Kpiinfo)

	Rulesdata.Kpiid = formdata.Kpiid
	Rulesdata.Title = formdata.Title
	Rulesdata.Score = formdata.Score
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
	err := models.Addkpiinfo(Rulesdata) //判断账号是否存在！
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
func EditKpiinfo(c *gin.Context) {
	var formdata models.Kpiinfo
	c.ShouldBind(&formdata)
	intodata := new(models.Kpiinfo)
	intodata.Id = formdata.Id
	intodata.Kpiid = formdata.Kpiid
	intodata.Title = formdata.Title
	intodata.Score = formdata.Score
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
		res, err := models.Upkpiinfo(intodata) //判断
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

func DelKpiinfo(c *gin.Context) {
	var searchdata models.Kpiinfo
	c.BindJSON(&searchdata)
	delnum := models.Deletekpiinfo(searchdata.Id)
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
