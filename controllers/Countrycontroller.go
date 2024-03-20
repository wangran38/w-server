package controllers

import (
	// "fmt"
	"changxiaoyang/models"
	_"time"
	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Countryserch struct {
    Id  int64 `json:"id"`
    Pid  int `json:"pid"`
	Name  string `json:"name"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}
// type Any interface{}
//获取当前用户信息
func Getcountrylist(c *gin.Context) {
	//从header中获取到token
	var searchdata Countryserch
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	search := &models.Country{
		Id:        searchdata.Id,
		Pid:       searchdata.Pid,
		Name:     searchdata.Name,
	}
	order := searchdata.Order
	listdata := models.GetCountryList(limit, page, search, order)
	listnum := models.GetCountrytotal(search)

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
//获取全部上下级
func Treecountry(c *gin.Context) {

        grouplist := models.Getcountrytree(0)
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data": grouplist,
		})

}


// //添加用户组
func Addcountry(c *gin.Context) {
	var formdata models.Country
	c.ShouldBind(&formdata)
		// 	c.JSON(200, gin.H{
		// 	"code": "201",
		// 	"msg":  "添加数据出错！",
		// 	"data": formdata,
		// })
	Rulesdata := new(models.Country)
	
	Rulesdata.Pid = formdata.Pid
	Rulesdata.Name = formdata.Name
	Rulesdata.Shortname = formdata.Shortname
	Rulesdata.Level = formdata.Level
	Rulesdata.Pinyin = formdata.Pinyin
	Rulesdata.Code = formdata.Code
	// Rulesdata.Created = time.Now()
	info, _ := models.SelectCountryByTitle(Rulesdata.Name) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "该城市已经存在！",
		})
		return
	}
	err := models.AddCountry(Rulesdata) //判断账号是否存在！
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
func Editcountry(c *gin.Context) {
	var formdata models.Country
	c.ShouldBind(&formdata)
	intodata := new(models.Country)
	intodata.Id = formdata.Id
	intodata.Pid = formdata.Pid
	intodata.Name = formdata.Name
	intodata.Shortname = formdata.Shortname
	intodata.Level = formdata.Level
	intodata.Pinyin = formdata.Pinyin
	intodata.Code = formdata.Code
	// intodata.Content = formdata.Content
	if(formdata.Id<=0) {
	c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res,err := models.UpCountry(intodata) //判断账号是否存在！
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

func Delcountry(c *gin.Context) {
	var searchdata models.Country
	c.BindJSON(&searchdata)
	delnum := models.DeleteCountry(searchdata.Id)
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