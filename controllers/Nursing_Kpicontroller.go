package controllers

import (
	// "fmt"
	"time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

// type NewsController struct{}
type Nursing_Kpi struct {
	Id         int64     `json:"id"`
	Nursing_id string    `json:"nursing_id"`
	Kpi_id     int       `json:"kpi_id"`
	Created    time.Time `json:"createtime"`
	Updated    time.Time `json:"updatetime"`
	P_id       int64     `json:"p_id"`
	Limit      int       `json:"limit"`
	Page       int       `json:"page"`
}

// type Any interface{}
// 获取当前用户信息
func GetNursing_Kpilist(c *gin.Context) {
	//从header中获取到token
	var searchdata Nursing_Kpi
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	search := &models.Nursing_Kpi{
		Id:         searchdata.Id,
		Nursing_id: searchdata.Nursing_id,
		Kpi_id:     searchdata.Nursing_id,
		P_id:       searchdata.P_id,
		Created:    searchdata.Created,
		Updated:    searchdata.Updated,
	}
	listdata := models.SelectNursing_Kpilist(limit, page, search)
	// listnum := models.GetHealthList(search)
	result["page"] = page
	// result["totalnum"] = listnum
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

// //添加用户组
func AddNursing_Kpi(c *gin.Context) {
	var formdata models.Nursing_Kpi
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Nursing_Kpi)
	Intodata.Id = formdata.Id
	Intodata.Nursing_id = formdata.Nursing_id
	Intodata.Kpi_id = formdata.Kpi_id
	Intodata.P_id = formdata.P_id
	Intodata.Created = time.Now()
	info, _ := models.SelectNursing_Kpiid(Intodata.Id) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "已经存在！",
		})
		return
	}
	err := models.AddNursing_Kpi(Intodata) //判断账号是否存在！
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
func UpNursing_Kpi(c *gin.Context) {
	var formdata models.Nursing_Kpi
	c.ShouldBind(&formdata)
	updata := new(models.Nursing_Kpi)
	updata.Id = formdata.Id
	updata.Nursing_id = formdata.Nursing_id
	updata.Kpi_id = formdata.Kpi_id
	updata.P_id = formdata.P_id
	updata.Updated = time.Now()
	if formdata.Id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.UpNursing_Kpi(updata) //判断账号是否存在！
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

// 删除功能
func DeleteNursing_Kpi(c *gin.Context) {
	var searchdata models.Nursing_Kpi
	c.BindJSON(&searchdata)
	delnum := models.DeleteNursing_Kpi(searchdata.Id)
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
