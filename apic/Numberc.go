package apic

import (
	// "fmt"
	// "net/http"

	"time"
	_ "time"
	"w-server/models"
	"w-server/utils"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Numberc struct {
	Id       int64  `json:"id"`
	Code     string `json:"code"`
	Codetime string `json:"codetime"`
	Reasons  string `json:"reasons"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}

// //添加用户组
func AddNumberc(c *gin.Context) {
	//如若是models则为models结构体的名称
	token := c.Request.Header.Get("Authorization")
	if token == "" || len(token) == 0 {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "你没有权限,去远处玩！",
			"data":    "",
			// "permissions": menu,
			// "roles":       role,
		})
		return
	}
	// tokennum := utils.CheckRedisExits(token)
	// if tokennum < 1 {
	// 	c.JSON(201, gin.H{
	// 		"code":    201,
	// 		"message": "登录已经过期！",
	// 		"data":    "",
	// 		// "permissions": menu,
	// 		// "roles":       role,
	// 	})
	// 	return
	// }
	user, _ := utils.GetLoginAssessorsc(token)
	if user.Id < 0 {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "你没有权限,去远处玩！",
			"data":    "",
			// "permissions": menu,
			// "roles":       role,
		})
		return
	}
	//也可继承controller里的结构体
	var formdata Numberc
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Number)

	// Intodata.Id = formdata.Id
	Intodata.Code = formdata.Code
	Intodata.Codetime = formdata.Codetime
	Intodata.Reasons = formdata.Reasons
	Intodata.Assessorsid = user.Id
	//Intodata.Created = time.Now()
	info, _ := models.SelectNumberCode(Intodata.Code) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "该编号已经存在！",
		})
		return
	}
	err := models.AddNumber(Intodata) //判断账号是否存在！
	if err != nil {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "添加数据出错！",
			"data": "",
		})
		return
	} else {
		// result := make(map[string]interface{})
		// result["id"] = Rulestable.Id //返回当前总数
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "数据添加成功！",
			"data": Intodata.Id,
		})

	}

}

// 我的编号
func MyNumberc(c *gin.Context) {
	//如若是models则为models结构体的名称
	token := c.Request.Header.Get("Authorization")
	if token == "" || len(token) == 0 {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "你没有权限,去远处玩！",
			"data":    "",
			// "permissions": menu,
			// "roles":       role,
		})
		return
	}

	var searchdata Numberc
	c.BindJSON(&searchdata)
	user, finderr := utils.GetLoginAssessorsc(token)
	if finderr != nil {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "登录已经过期！111",
			"data":    "",
			// "permissions": menu,
			// "roles":       role,
		})
		return
	}
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	search := &models.Number{
		Id:          searchdata.Id,
		Assessorsid: user.Id,
		Code:        searchdata.Code,
	}
	listdata := models.GetnumberList(limit, page, search)
	listnum := models.Getnumbertotal(search)

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

// 我的编号
func EditNumberc(c *gin.Context) {
	//如若是models则为models结构体的名称
	token := c.Request.Header.Get("Authorization")
	if token == "" || len(token) == 0 {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "你没有权限,去远处玩！",
			"data":    "",
			// "permissions": menu,
			// "roles":       role,
		})
		return
	}

	user, _ := utils.GetLoginAssessorsc(token)
	if user.Id < 0 {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "你没有权限,去远处玩！",
			"data":    "",
			// "permissions": menu,
			// "roles":       role,
		})
		return
	}
	//也可继承controller里的结构体
	var formdata Numberc
	c.ShouldBind(&formdata)
	Intodata := new(models.Number)
	Intodata.Id = formdata.Id
	Intodata.Code = formdata.Code
	Intodata.Codetime = formdata.Codetime
	Intodata.Reasons = formdata.Reasons
	Intodata.Assessorsid = user.Id
	Intodata.Updated = time.Now()

	_, err := models.UpNumber(Intodata) //判断账号是否存在！
	if err != nil {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "添加数据出错！",
			"data": "",
		})
		return
	} else {
		// result := make(map[string]interface{})
		// result["id"] = Rulestable.Id //返回当前总数
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "修改编号成功！",
			"data": "",
		})

	}

}
