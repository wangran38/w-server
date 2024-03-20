package controllers

import (
	// "fmt"
	"changxiaoyang/models"
	"time"
	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Placeserch struct {
	Id      int64     `json:"id"`
    Cityid     int64       `json:"city_id"`
	Countryid     int64       `json:"country_id"`
	Title    string    `json:"title" xorm:"varchar(200)"`
	Image   string    `json:"image" xorm:"TEXT "`
	Keywords   string  `json:"keywords" xorm:"TEXT "`
	Description   string  `json:"description" xorm:" TEXT "`
    Oldlink    string    `json:"oldlink" xorm:"varchar(200)"`
	Address    string    `json:"address" xorm:"varchar(200)"`
    Lng int `json:"lng"`  //地址经度
    Lat int `json:"lat"`  //地址维度
	Content   string  `json:"content" xorm:"LONGTEXT "`
	Isshow     int       `json:"isshow" xorm:"not null default 1 comment('是否启用 默认1 是 0 无') TINYINT"`
	Created time.Time `json:"createtime" xorm:"int"`
	Updated time.Time `json:"updatetime" xorm:"int"`
    Iscountry     int       `json:"iscountry" xorm:"not null default 1 comment('是否启用 默认2 是 1 无') TINYINT"`
	Weigh   int  `json:"weigh"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Order     string `json:"order"`
}
// type Any interface{}
//获取当前用户信息
func Getplacelist(c *gin.Context) {
	//从header中获取到token
	var searchdata Placeserch
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	psearch := &models.Places{
		Title:     searchdata.Title,
	}
	order := searchdata.Order
	listdata := models.GetPlacesList(limit, page, psearch, order)
	listnum := models.GetPlacestotal(psearch)

	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取数据为空或失败",
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
//获取全部上下级
// func TreePlace(c *gin.Context) {

//         grouplist := models.GetPlacetree(0)
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"message": "数据获取成功",
// 			"data": grouplist,
// 		})

// }


// //添加用户组
// func AddPlace(c *gin.Context) {
// 	var formdata models.Place
// 	c.ShouldBind(&formdata)
// 		// 	c.JSON(200, gin.H{
// 		// 	"code": "201",
// 		// 	"msg":  "添加数据出错！",
// 		// 	"data": formdata,
// 		// })
// 	Rulesdata := new(models.Place)
	
// 	Rulesdata.Pid = formdata.Pid
// 	Rulesdata.Title = formdata.Title
// 	Rulesdata.Image = formdata.Image
// 	Rulesdata.Keywords = formdata.Keywords
// 	Rulesdata.Description = formdata.Description
// 	Rulesdata.Content = formdata.Content
// 	Rulesdata.Created = time.Now()
// 	info, _ := models.SelectPlaceByTitle(Rulesdata.Title) //判断账号是否存在！
// 	if info != nil {
// 		c.JSON(200, gin.H{
// 			"code": "201",
// 			"msg":  "该分类已经存在！",
// 		})
// 		return
// 	}
// 	err := models.AddPlace(Rulesdata) //判断账号是否存在！
// 		if err != nil {
// 		c.JSON(201, gin.H{
// 			"code": 201,
// 			"msg":  "添加数据出错！",
// 			"data": err,
// 		})
// 		return
// 	} else {
// 		// result := make(map[string]interface{})
// 		// result["id"] = Rulestable.Id //返回当前总数
// 		c.JSON(200, gin.H{
// 			"code": 200,
// 			"msg":  "数据添加成功！",
// 			"data": "",
// 		})

// 	}
	
// }

// // //修改用户组
// func EditPlace(c *gin.Context) {
// 	var formdata models.Place
// 	c.ShouldBind(&formdata)
// 	intodata := new(models.Place)
// 	intodata.Id = formdata.Id
// 	intodata.Pid = formdata.Pid
// 	intodata.Title = formdata.Title
// 	intodata.Isshow = formdata.Isshow
// 	intodata.Image = formdata.Image
// 	intodata.Keywords = formdata.Keywords
// 	intodata.Description = formdata.Description
// 	intodata.Content = formdata.Content
// 	if(formdata.Id<=0) {
// 	c.JSON(201, gin.H{
// 			"code": 201,
// 			"msg":  "修改选择的ID出错！",
// 			"data": "",
// 		})
// 		return
// 	} else {
// 		res,err := models.UpPlace(intodata) //判断账号是否存在！
// 		if err != nil {
// 		c.JSON(201, gin.H{
// 			"code": 201,
// 			"msg":  "修改数据出错！",
// 			"data": err,
// 		})
// 		return
// 	} else {
		
// 		c.JSON(200, gin.H{
// 			"code": 200,
// 			"msg":  "数据修改成功！",
// 			"data": res,
// 		})

// 	}
// 	}

// }

// func DelPlace(c *gin.Context) {
// 	var searchdata models.Place
// 	c.BindJSON(&searchdata)
// 	delnum := models.DeletePlace(searchdata.Id)
// 	if delnum > 0 {
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"message": "删除成功！",
// 			"data":    delnum,
// 		})
// 	} else {
// 		c.JSON(200, gin.H{
// 			"code": 200,
// 			"msg":  "操作失败！",
// 		})

// 	}

// }