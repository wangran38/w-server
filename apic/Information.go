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

// Information 结构体表示用户信息
type Information struct {
	Id           int64  `json:"Id"`
	Senior_id    int64  `json:"senior_id"`
	Provider     string `json:"provider"`
	Assessors_id int64  `json:"assessors_id"`
	Number_id    int64  `json:"number_id"`
	Relationship string `json:"relationship"`
	Contactname  string `json:"contactname"`
	Phone        string `json:"phone"`
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
	// Order        string `json:"order"`
}

// AddInformation 用于添加用户信息
func AddInformation(c *gin.Context) {
	// 从请求头中获取令牌
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
	user, tokenerr := utils.GetLoginAssessorsc(token)
	if tokenerr != nil {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "登录失效，请重新登录！",
			"data":    "",
		})
		return
	}
	var formdata Information
	c.ShouldBind(&formdata)
	Intodata := new(models.Information)
	// Intodata.Id = formdata.Id
	Intodata.Senior_id = formdata.Senior_id
	Intodata.Assessors_id = user.Id
	Intodata.Provider = formdata.Provider
	Intodata.Number_id = formdata.Number_id
	Intodata.Relationship = formdata.Relationship
	Intodata.Contactname = formdata.Contactname
	Intodata.Phone = formdata.Phone
	err := models.AddInformation(Intodata) // 判断账号是否存在！
	if err != nil {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "添加数据出错！",
			"data": err,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "数据添加成功！",
			"data": Intodata.Id,
		})
	}
}

// 获取信息
func GetInformationlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Information
	c.ShouldBind(&searchdata)
	// //读取数据库

	search := &models.Information{
		Id:           searchdata.Id,
		Senior_id:    searchdata.Senior_id,
		Assessors_id: searchdata.Assessors_id,
		Number_id:    searchdata.Number_id,
		Provider:     searchdata.Provider,
		Relationship: searchdata.Relationship,
		Contactname:  searchdata.Contactname,
		Phone:        searchdata.Phone,
	}
	listdata := models.GetInformationList(10, 1, search)
	listnum := models.GetInformationtotal(search)
	result := make(map[string]interface{})
	result["totalnum"] = listnum
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取数据失败",
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

func GetInformationList1(c *gin.Context) {
	var searchdata Information
	c.ShouldBind(&searchdata)
	// fmt.Print(searchinfo.Id)
	// result := make(map[string]interface{})
	info, _ := models.SelectInformationId(searchdata.Id)
	if info != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "获取成功",
			"data": info,
		})

	} else {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "获取数据失败",
			"data": searchdata.Id,
		})
	}
}

// 信息提供者编号
func MyInformation(c *gin.Context) {
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

	var searchdata Information
	c.ShouldBind(&searchdata)
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
	search := &models.Information{
		Id:           searchdata.Id,
		Senior_id:    searchdata.Senior_id,
		Assessors_id: user.Id,
		Number_id:    searchdata.Number_id,
		Provider:     searchdata.Provider,
		Relationship: searchdata.Relationship,
		Contactname:  searchdata.Contactname,
		Phone:        searchdata.Phone,
	}
	//  var byorder string
	//  byorder = "id DESC"
	listdata := models.GetInformationList(limit, page, search)
	listnum := models.GetInformationtotal(search)

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

// 修改信息提供者及联系人接口
func UpInformation(c *gin.Context) {
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
	user, tokenerr := utils.GetLoginAssessorsc(token)
	if tokenerr != nil {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "登录失效，请重新登录！",
			"data":    "",
			// "permissions": menu,
			// "roles":       role,
		})
		return
	}
	//也可继承controller里的结构体
	var formdata Information
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Information)
	Intodata.Id = formdata.Id
	Intodata.Senior_id = formdata.Senior_id
	Intodata.Assessors_id = user.Id
	Intodata.Number_id = formdata.Number_id
	Intodata.Provider = formdata.Provider
	Intodata.Relationship = formdata.Relationship
	Intodata.Contactname = formdata.Contactname
	Intodata.Phone = formdata.Phone
	Intodata.Created = time.Now()
	// info, _ := models.SelectSeniorbyidnum(Intodata.Idnumber) //判断账号是否存在！
	// if info != nil {
	// 	c.JSON(200, gin.H{
	// 		"code": "201",
	// 		"msg":  "该老人身份证已经存在！",
	// 	})
	// 	return
	// }
	_, err := models.UpInformation(Intodata) //判断账号是否存在！
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
			"data": Intodata.Id,
		})

	}

}
