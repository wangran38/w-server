package controllers

import (
	// "fmt"
	"time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Kpiserch struct {
	Kname string `json:"kname"`
	Id    int64  `json:"id"`
	Pid   int64  `json:"pid"`
	Level int    `json:"level"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}

// type Any interface{}
// 获取当前用户信息
func GetKpilist(c *gin.Context) {
	//从header中获取到token
	var searchdata Kpiserch
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// kname := searchdata.Kname
	order := searchdata.Order
	search := &models.Kpi{
		// Id:    searchdata.Id,
		Pid:   searchdata.Pid,
		Kname: searchdata.Kname,
		Level: searchdata.Level,
	}
	listdata := models.GetkpiList(limit, page, search, order)

	listnum := models.Getkpitotal(search)

	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取菜单失败",
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
func TreeKpi(c *gin.Context) {

	grouplist := models.GetKpitree(0)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "数据获取成功",
		"data":    grouplist,
	})

}

// //添加用户组
func AddKpi(c *gin.Context) {
	var formdata models.Kpi
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Rulesdata := new(models.Kpi)

	Rulesdata.Pid = formdata.Pid
	Rulesdata.Kname = formdata.Kname
	Rulesdata.Level = formdata.Level
	Rulesdata.Kcontent = formdata.Kcontent
	Rulesdata.Created = time.Now()
	info, _ := models.SelectkpiByTitle(Rulesdata.Kname) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "该指标已经存在！",
		})
		return
	}
	err := models.Addkpi(Rulesdata) //判断账号是否存在！
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
func EditKpi(c *gin.Context) {
	var formdata models.Kpi
	c.ShouldBind(&formdata)
	intodata := new(models.Kpi)
	intodata.Id = formdata.Id
	intodata.Pid = formdata.Pid
	intodata.Kname = formdata.Kname
	intodata.Level = formdata.Level
	intodata.Weigh = formdata.Weigh
	intodata.Kcontent = formdata.Kcontent
	if formdata.Id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.Upkpi(intodata) //判断
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

func DelKpi(c *gin.Context) {
	var searchdata models.Kpi
	c.BindJSON(&searchdata)
	delnum := models.Deletekpi(searchdata.Id)
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
