package controllers

import (
	// "fmt"
	"time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

// type NewsController struct{}
type Nursing struct {
	Id          int64     `json:"id"` //id
	Shift       string    `json:"shift" xorm:"comment('班次')"`
	Timeperiod  string    `json:"timeperiod" xorm:"comment('时间段')"`
	Nursingid   int64     `json:"nursingid" xorm:"comment('护理动作id')"`
	Nursingname string    `json:"nursingname" xorm:"comment('护理动作名称')"`
	Created     time.Time `json:"createtime" xorm:"int"`
	Updated     time.Time `json:"updatetime" xorm:"int"`
	Limit       int       `json:"limit"`
	Page        int       `json:"page"`
}

// type Any interface{}
// 获取当前用户信息
func GetNursinglist(c *gin.Context) {
	//从header中获取到token
	var searchdata Nursing
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	search := &models.Nursing{
		Id:          searchdata.Id,
		Shift:       searchdata.Shift,
		Timeperiod:  searchdata.Timeperiod,
		Nursingid:   searchdata.Nursingid,
		Nursingname: searchdata.Nursingname,
		Created:     searchdata.Created,
		Updated:     searchdata.Updated,
	}
	listdata := models.SelectNursinglist(limit, page, search)
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
func AddNursing(c *gin.Context) {
	var formdata models.Nursing
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Nursing)
	Intodata.Id = formdata.Id
	Intodata.Shift = formdata.Shift
	Intodata.Timeperiod = formdata.Timeperiod
	Intodata.Nursingid = formdata.Nursingid
	Intodata.Nursingname = formdata.Nursingname
	Intodata.Created = time.Now()
	info, _ := models.SelectNursingid(Intodata.Id) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "已经存在！",
		})
		return
	}
	err := models.AddNursing(Intodata) //判断账号是否存在！
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
func UpNursing(c *gin.Context) {
	var formdata models.Nursing
	c.ShouldBind(&formdata)
	updata := new(models.Nursing)
	updata.Id = formdata.Id
	updata.Shift = formdata.Shift
	updata.Timeperiod = formdata.Timeperiod
	updata.Nursingid = formdata.Nursingid
	updata.Nursingname = formdata.Nursingname
	updata.Updated = time.Now()
	if formdata.Id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.UpNursing(updata) //判断账号是否存在！
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
func DeleteNursing(c *gin.Context) {
	var searchdata models.Nursing
	c.BindJSON(&searchdata)
	delnum := models.DeleteNursing(searchdata.Id)
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
