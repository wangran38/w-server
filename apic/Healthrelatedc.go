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

type Healthrelated struct {
	Id             int64     `json:"id"`               //id
	Senior_id      int64     `json:"senior_id"`        //老者id
	Assessors_id   int64     `json:"assessors_id"`     //所属评估员id
	Number_id      int64     ` json:"number_id"`       //编号id
	Pressureinjury string    `json:"pressureinjury"`   //压力性损伤
	Joint          string    ` json:"Joint"`           //关节活动度
	Affectedarea   string    ` json:"Affectedarea"`    //关节影响部位
	Woundcondition string    `  json:"woundcondition"` //伤口情况
	Specialcare    string    ` json:"specialcare"`     //特殊护理
	Painsensation  string    `json:"painsensation"`    //疼痛感
	Toothloss      string    ` json:"toothloss"`       //牙齿缺失
	Wearing        string    `  json:"wearing"`        //义齿佩戴
	Swallow        string    `  json:"swallow"`        //吞咽困难的情形和状况
	Malnutrition   string    `  json:"malnutrition"`   //营养不良
	Cleaning       string    `  json:"cleaning"`       //清理呼吸道无效
	Coma           string    ` json:"coma"`            //昏迷
	Other          string    ` json:"other"`           //其他
	Limit          int       `json:"limit"`
	Page           int       `json:"page"`
	Created        time.Time `xorm:"int" json:"createtime"`
	Updated        time.Time `xorm:"int" json:"updatetime"`
}

// 获取信息
func GetHealthrelatedlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Healthrelated
	c.ShouldBind(&searchdata)
	// //读取数据库

	search := &models.Healthrelated{
		Id:             searchdata.Id,
		Senior_id:      searchdata.Senior_id,
		Assessors_id:   searchdata.Assessors_id,
		Number_id:      searchdata.Number_id,
		Pressureinjury: searchdata.Pressureinjury,
		Joint:          searchdata.Joint,
		Affectedarea:   searchdata.Affectedarea,
		Woundcondition: searchdata.Woundcondition,
		Specialcare:    searchdata.Specialcare,
		Painsensation:  searchdata.Painsensation,
		Toothloss:      searchdata.Toothloss,
		Wearing:        searchdata.Wearing,
		Swallow:        searchdata.Swallow,
		Malnutrition:   searchdata.Malnutrition,
		Cleaning:       searchdata.Cleaning,
		Coma:           searchdata.Coma,
		Other:          searchdata.Other,
	}
	listdata := models.GetHealthrelatedList(10, 1, search)
	listnum := models.GetHealthrelatedtotal(search)
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

func GetHealthrelatedList1(c *gin.Context) {
	var searchdata Healthrelated
	c.ShouldBind(&searchdata)
	info, _ := models.SelectHealthrelatedid(searchdata.Id)
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

// AddHealth 用于添加用户信息
func AddHealthrelated(c *gin.Context) {
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
	var formdata Healthrelated
	c.ShouldBind(&formdata)
	Intodata := new(models.Healthrelated)
	Intodata.Senior_id = formdata.Senior_id
	Intodata.Assessors_id = user.Id
	Intodata.Number_id = formdata.Number_id
	Intodata.Pressureinjury = formdata.Pressureinjury
	Intodata.Joint = formdata.Joint
	Intodata.Affectedarea = formdata.Affectedarea
	Intodata.Woundcondition = formdata.Woundcondition
	Intodata.Specialcare = formdata.Specialcare
	Intodata.Painsensation = formdata.Painsensation
	Intodata.Toothloss = formdata.Toothloss
	Intodata.Wearing = formdata.Wearing
	Intodata.Swallow = formdata.Swallow
	Intodata.Malnutrition = formdata.Malnutrition
	Intodata.Cleaning = formdata.Cleaning
	Intodata.Coma = formdata.Coma
	Intodata.Other = formdata.Other

	err := models.AddHealthrelated(Intodata) // 判断账号是否存在！
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

// 疾病编号
func MyHealthrelated(c *gin.Context) {
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

	var searchdata Healthrelated
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
	search := &models.Healthrelated{
		Id:             searchdata.Id,
		Senior_id:      searchdata.Senior_id,
		Assessors_id:   user.Id,
		Number_id:      searchdata.Number_id,
		Pressureinjury: searchdata.Pressureinjury,
		Joint:          searchdata.Joint,
		Affectedarea:   searchdata.Affectedarea,
		Woundcondition: searchdata.Woundcondition,
		Specialcare:    searchdata.Specialcare,
		Painsensation:  searchdata.Painsensation,
		Toothloss:      searchdata.Toothloss,
		Wearing:        searchdata.Wearing,
		Swallow:        searchdata.Swallow,
		Malnutrition:   searchdata.Malnutrition,
		Cleaning:       searchdata.Cleaning,
		Coma:           searchdata.Coma,
		Other:          searchdata.Other,
	}
	// var byorder string
	// byorder = "id DESC"
	listdata := models.GetHealthrelatedList(limit, page, search) //byorder
	listnum := models.GetHealthrelatedtotal(search)

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

// 修改健康相关问题
func UpHealthrelated(c *gin.Context) {
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
	var formdata Healthrelated
	c.ShouldBind(&formdata)
	Intodata := new(models.Healthrelated)
	Intodata.Id = formdata.Id
	Intodata.Senior_id = formdata.Senior_id
	Intodata.Assessors_id = user.Id
	Intodata.Number_id = formdata.Number_id
	Intodata.Pressureinjury = formdata.Pressureinjury
	Intodata.Joint = formdata.Joint
	Intodata.Affectedarea = formdata.Affectedarea
	Intodata.Woundcondition = formdata.Woundcondition
	Intodata.Specialcare = formdata.Specialcare
	Intodata.Painsensation = formdata.Painsensation
	Intodata.Toothloss = formdata.Toothloss
	Intodata.Wearing = formdata.Wearing
	Intodata.Swallow = formdata.Swallow
	Intodata.Malnutrition = formdata.Malnutrition
	Intodata.Cleaning = formdata.Cleaning
	Intodata.Coma = formdata.Coma
	Intodata.Other = formdata.Other
	Intodata.Created = time.Now()
	_, err := models.UpHealthrelated(Intodata) //判断账号是否存在！
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
