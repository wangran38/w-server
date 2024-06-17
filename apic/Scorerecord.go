package apic

import (
	// "fmt"
	// "net/http"

	"fmt"
	"time"
	_ "time"
	"strings"
	"w-server/models"
	"w-server/utils"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Score_recordc struct {
	Id          int64 `json:"id"`           //id
	Seniorid    int64 `json:"senior_id"`    //老者id
	Assessorsid int64 `json:"assessors_id"` //所属评估员id
	Numberid    int64 `json:"number_id"`    //编号id
	Kpiid       int64 `json:"kpi_id"`       //KPI的id
	Kpiids      string `json:"Kpiids"`       //KPI的id
	Kpiinfoid   int64 `json:"kpiinfo_id"`   //KPI的id
	Score       int   `json:"score"`        //分数
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}
type Score_recordjson struct {
	Id          int64 `json:"id"`           //id
	Seniorid    int64 `json:"senior_id"`    //老者id
	Assessorsid int64 `json:"assessors_id"` //所属评估员id
	Numberid    int64 `json:"number_id"`    //编号id
	Kpiid       int64 `json:"kpi_id"`       //KPI的id
	Kname       string `json:"kpi_kname"`
	Kpiinfoid   int64 `json:"kpiinfo_id"`   //KPI的id
	Kpiinfotitle string `json:"kpiinfo_title"`
	Score       int   `json:"score"`        //药物名称
}
// //添加用户组
func AddScore_record(c *gin.Context) {
	//如若是models则为models结构体的名称
	token := c.Request.Header.Get("Authorization")
	user, tokenerr := utils.GetLoginAssessorsc(token)
	if tokenerr != nil {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "登录失效，请重新登录！",
			"data":    "",
		})
		return
	}

	var Score_arr []Score_recordc
	// var formdata Score_arr
	c.ShouldBind(&Score_arr)
	//var orm *xorm.EngineGroup
	// session := orm.NewSession()
	// defer session.Close()

	// add Begin() before any action
	// if err := session.Begin(); err != nil {
	// 	c.JSON(200, gin.H{
	// 		"code": 201,
	// 		"msg":  "添加失败，已回滚数据",
	// 		"data": err,
	// 	})
	// 	return
	// }
	sum := 0
	//Intodata := make([]*models.Score_record, len(Score_arr))

	for _, item := range Score_arr {
		Intodata := new(models.Score_record)
		Intodata.Seniorid = item.Seniorid
		Intodata.Assessors_id = user.Id
		Intodata.Number_id = item.Numberid
		Intodata.Kpi_id = item.Kpiid
		Intodata.Kpiinfo_id = item.Kpiinfoid
		Intodata.Score = item.Score
		// fmt.Printf("Number: %d\n", item.Seniorid)
		Intodata.Created = time.Now()
		// sum++
		err := models.AddScore_record(Intodata) //判断账号是否存在！
		if err == nil {
			sum+=item.Score
			// return
		}
		// else {
		// 	session.Rollback()
		// 	c.JSON(200, gin.H{
		// 		"code": 201,
		// 		"msg":  "添加失败，已回滚数据",
		// 		"data": sum,
		// 	})
		// 	return
		// }
	}
	// session.Commit()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "当前KPI分数值",
		"data": sum,
	})
	// c.JSON(201, gin.H{
	// 	"code":    201,
	// 	"message": "登录失效，请重新登录！",
	// 	"data":    Intodata,
	// 	// "permissions": menu,
	// 	// "roles":       role,
	// })

}
// 我的量表提交记录
func Myscorerecord(c *gin.Context) {
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

	var searchdata Score_recordc
	c.ShouldBind(&searchdata)
	user, finderr := utils.GetLoginAssessorsc(token)
	if finderr != nil {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "登录已经过期!111",
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
	search := &models.Score_record{
		Id:             searchdata.Id,
		Seniorid:      searchdata.Seniorid,
		Assessors_id:   user.Id,
		Number_id:      searchdata.Numberid,
		Kpi_id:      searchdata.Kpiid,
		Kpiinfo_id:      searchdata.Kpiinfoid,
	// Numberid    int64 `json:"number_id"`    //编号id
	// Kpiid       int64 `json:"kpi_id"`       //KPI的id
	// Kpiinfoid   int64 `json:"kpiinfo_id"`   //KPI的id
	// Score       int   `json:"score"`        //药物名称

	}
	// var byorder string
	// byorder = "id DESC"
	listdata := models.GetScore_recordList(limit, page, search) //byorder
	listnum := models.GetScore_recordtotal(search)

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

// 我的kpi量表提交记录
func Mykpirecord(c *gin.Context) {
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

	var searchdata Score_recordc
	c.ShouldBind(&searchdata)
	user, finderr := utils.GetLoginAssessorsc(token)
	if finderr != nil {
		c.JSON(201, gin.H{
			"code":    201,
			"message": "登录已经过期!111",
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

	kpisearch:= &models.Kpi {
		Pid: searchdata.Kpiid,
	}
	var order string 
	order="weigh DESC"
	kpidata:= models.GetkpiapiList(1000, 1, kpisearch, order)
	if len(kpidata)<1 {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "当前量表数据为空",
			"data":    0,
		})
		return
	}
	var str_arr = make([]string, len(kpidata))
	// strids := strings.Join(strs, " ")
	for k,v := range kpidata {
		str_arr[k] = fmt.Sprintf("%d", v.Id)
// str := fmt.Sprintf("%d", v.Id)
// 		ids := strings.Join(str, ",")
	}
	var ids = strings.Join(str_arr, ",")
	// var byorder string
	// byorder = "id DESC"
		search := &models.Score_record{
		Id:             searchdata.Id,
		Seniorid:      searchdata.Seniorid,
		Assessors_id:   user.Id,
		Number_id:      searchdata.Numberid,
		// Kpi_id:      searchdata.Kpiid,
		Kpiids:      ids,
		// Kpiinfo_id:      searchdata.Kpiinfoid,
	// Numberid    int64 `json:"number_id"`    //编号id
	// Kpiid       int64 `json:"kpi_id"`       //KPI的id
	// Kpiinfoid   int64 `json:"kpiinfo_id"`   //KPI的id
	// Score       int   `json:"score"`        //药物名称

	}
	listdata := models.GetScore_recordList(limit, page, search) //byorder
	listnum := models.GetScore_recordtotal(search)
	if listnum<=1 {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "当前量表数据为空",
			"data":    0,
		})
		return
	}
	var allscore int
	for _,v := range listdata {
		allscore+=v.Score
	}
	result["page"] = page
	result["totalnum"] = listnum
	result["allscore"] = allscore
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