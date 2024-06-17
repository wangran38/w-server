package apic

import (
	// "fmt"
	// "net/http"
	"time"
	_ "time"
	"w-server/models"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Nursing_Kpi struct {
	Id         int64     `json:"id"`          //ID自增长
	Nursing_id string    ` json:"nursing_id"` //护理动作编号
	Kpi_id     string    ` json:"kpi_id"`     //题目id
	Created    time.Time `json:"createtime"`
	Updated    time.Time `json:"updatetime"`
	P_id       int64     ` json:"p_id"` //标识符
}

// 获取信息
func SelectNursing_Kpilist(c *gin.Context) {
	//从header中获取到token
	var searchdata Nursing_Kpi
	c.ShouldBind(&searchdata)
	// //读取数据库

	search := &models.Nursing_Kpi{
		Id:         searchdata.Id,
		Nursing_id: searchdata.Nursing_id,
		Kpi_id:     searchdata.Kpi_id,
		P_id:       searchdata.P_id,
	}
	listdata := models.SelectNursing_Kpilist(10, 1, search)
	listnum := models.GetNursing_kpitotal(search)
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

func GetNursing_KpiList(c *gin.Context) {
	var searchdata Nursing_Kpi
	c.ShouldBind(&searchdata)
	// fmt.Print(searchinfo.Id)
	// result := make(map[string]interface{})
	info, _ := models.SelectNursingid(searchdata.Id)
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
// func MyNursing(c *gin.Context) {
// 	//如若是models则为models结构体的名称
// 	token := c.Request.Header.Get("Authorization")
// 	if token == "" || len(token) == 0 {
// 		c.JSON(201, gin.H{
// 			"code":    201,
// 			"message": "你没有权限,去远处玩！",
// 			"data":    "",
// 			// "permissions": menu,
// 			// "roles":       role,
// 		})
// 		return
// 	}

// 	var searchdata Nursing
// 	c.ShouldBind(&searchdata)
// 	user, finderr := utils.GetLoginAssessorsc(token)
// 	if finderr != nil {
// 		c.JSON(201, gin.H{
// 			"code":    201,
// 			"message": "登录已经过期！111",
// 			"data":    "",
// 		})
// 		return
// 	}
// 	result := make(map[string]interface{})
// 	// name:=""
// 	limit := searchdata.Limit
// 	page := searchdata.Page
// 	search := &models.Nursing{
// 		Id:          searchdata.Id,
// 		Shift:       searchdata.Shift,
// 		Timeperiod:  searchdata.Timeperiod,
// 		Nursingid:   searchdata.Nursingid,
// 		Nursingname: searchdata.Nursingname,
// 	}
// 	// var byorder string
// 	// byorder = "id DESC"
// 	listdata := models.SelectNursinglist(limit, page, search) //byorder
// 	listnum := models.GetNursingtotal(search)

// 	result["page"] = page
// 	result["totalnum"] = listnum
// 	result["limit"] = limit
// 	if listdata == nil {
// 		c.JSON(200, gin.H{
// 			"code":    201,
// 			"message": "获取数据为空",
// 			"data":    "",
// 		})
// 		return
// 	} else {
// 		result["listdata"] = listdata
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"message": "数据获取成功",
// 			"data":    result,
// 		})
// 		return
// 	}

// }

// // 修改疾病信息
// func UpNursing(c *gin.Context) {
// 	//如若是models则为models结构体的名称
// 	token := c.Request.Header.Get("Authorization")
// 	if token == "" || len(token) == 0 {
// 		c.JSON(201, gin.H{
// 			"code":    201,
// 			"message": "你没有权限,去远处玩！",
// 			"data":    "",
// 			// "permissions": menu,
// 			// "roles":       role,
// 		})
// 		return
// 	}
// 	user, tokenerr := utils.GetLoginAssessorsc(token)
// 	if tokenerr != nil {
// 		c.JSON(201, gin.H{
// 			"code":    201,
// 			"message": "登录失效，请重新登录！",
// 			"data":    "",
// 		})
// 		return
// 	}
// 	//也可继承controller里的结构体
// 	var formdata Nursing
// 	c.ShouldBind(&formdata)
// 	Intodata := new(models.Nursing)
// 	Intodata.Id = formdata.Id
// 	Intodata.Shift = formdata.Shift
// 	Intodata.Timeperiod = formdata.Timeperiod
// 	Intodata.Nursingid = formdata.Nursingid
// 	Intodata.Nursingname = formdata.Nursingname
// 	Intodata.Created = time.Now()
// 	// info, _ := models.SelectSeniorbyidnum(Intodata.Idnumber) //判断账号是否存在！
// 	// if info != nil {
// 	// 	c.JSON(200, gin.H{
// 	// 		"code": "201",
// 	// 		"msg":  "该老人身份证已经存在！",
// 	// 	})
// 	// 	return
// 	// }
// 	_, err := models.UpNursing(Intodata) //判断账号是否存在！
// 	if err != nil {
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
// 			"data": Intodata.Id,
// 		})

// 	}

// }