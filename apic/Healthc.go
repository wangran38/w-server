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

type Health struct {
	Id           int64  `json:"id"`           //id
	Senior_id    int64  `json:"senior_id"`    //老者id
	Assessors_id int64  `json:"assessors_id"` //所属评估员id
	Number_id    int64  ` json:"number_id"`   //编号id
	Disease      string `json:"disease"`      //疾病诊断
	Drugname     string ` json:"drugname"`    //药物名称
	Medication   string ` json:"medication"`  //服药方法
	Dosage       string `json:"dosage"`       //用药剂量
	Frequency    string `json:"frequency"`    //用药频率
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
}

// AddHealth 用于添加用户信息
func AddHealth(c *gin.Context) {
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
			// "permissions": menu,
			// "roles":       role,
		})
		return
	}
	var formdata Health
	c.ShouldBind(&formdata)
	Intodata := new(models.Health)
	// Intodata.Id = formdata.Id
	Intodata.Senior_id = formdata.Senior_id
	Intodata.Assessors_id = user.Id
	Intodata.Number_id = formdata.Number_id
	Intodata.Disease = formdata.Disease
	Intodata.Drugname = formdata.Drugname
	Intodata.Medication = formdata.Medication
	Intodata.Dosage = formdata.Dosage
	Intodata.Frequency = formdata.Frequency

	err := models.AddHealth(Intodata) // 判断账号是否存在！
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
func GetHealthlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Health
	c.ShouldBind(&searchdata)
	// //读取数据库

	search := &models.Health{
		Id:           searchdata.Id,
		Senior_id:    searchdata.Senior_id,
		Assessors_id: searchdata.Assessors_id,
		Number_id:    searchdata.Number_id,
		Disease:      searchdata.Disease,
		Drugname:     searchdata.Drugname,
		Medication:   searchdata.Medication,
		Dosage:       searchdata.Dosage,
		Frequency:    searchdata.Frequency,
	}
	listdata := models.GetHealthList(10, 1, search)
	listnum := models.GetHealthtotal(search)
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

func GetHealthList1(c *gin.Context) {
	var searchdata Health
	c.ShouldBind(&searchdata)
	// fmt.Print(searchinfo.Id)
	// result := make(map[string]interface{})
	info, _ := models.SelectHealthid(searchdata.Id)
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

// 疾病编号
func MyHealth(c *gin.Context) {
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

	var searchdata Health
	c.ShouldBind(&searchdata)
	user, finderr := utils.GetLoginAssessorsc(token)
	if finderr != nil {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "登录已经过期！111",
			"data":    "",
		})
		return
	}
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	search := &models.Health{
		Id:           searchdata.Id,
		Senior_id:    searchdata.Senior_id,
		Assessors_id: user.Id,
		Number_id:    searchdata.Number_id,
		Disease:      searchdata.Disease,
		Drugname:     searchdata.Drugname,
		Medication:   searchdata.Medication,
		Dosage:       searchdata.Dosage,
		Frequency:    searchdata.Frequency,
	}
	// var byorder string
	// byorder = "id DESC"
	listdata := models.GetHealthList(limit, page, search) //byorder
	listnum := models.GetHealthtotal(search)

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

// 修改疾病信息
func UpHealth(c *gin.Context) {
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
		})
		return
	}
	//也可继承controller里的结构体
	var formdata Health
	c.ShouldBind(&formdata)
	Intodata := new(models.Health)
	Intodata.Id = formdata.Id
	Intodata.Senior_id = formdata.Senior_id
	Intodata.Assessors_id = user.Id
	Intodata.Number_id = formdata.Number_id
	Intodata.Disease = formdata.Disease
	Intodata.Drugname = formdata.Drugname
	Intodata.Medication = formdata.Medication
	Intodata.Dosage = formdata.Dosage
	Intodata.Frequency = formdata.Frequency
	Intodata.Created = time.Now()
	// info, _ := models.SelectSeniorbyidnum(Intodata.Idnumber) //判断账号是否存在！
	// if info != nil {
	// 	c.JSON(200, gin.H{
	// 		"code": "201",
	// 		"msg":  "该老人身份证已经存在！",
	// 	})
	// 	return
	// }
	_, err := models.UpHealth(Intodata) //判断账号是否存在！
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
