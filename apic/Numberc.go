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
			"data": Intodata,
		})

	}

}
// 我的编号
func MyNumberinfo(c *gin.Context) {
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
	user, _ := utils.GetLoginAssessorsc(token)
	// if finderr != nil {
	// 	c.JSON(201, gin.H{
	// 		"code":    201,
	// 		"message": "登录已经过期！111",
	// 		"data":    "",
	// 		// "permissions": menu,
	// 		// "roles":       role,
	// 	})
	// 	return
	// }
	Number_id := searchdata.Id
	result := make(map[string]interface{})
	//获取对应评估信息的编号信息
	numberdata,err:= models.SelectNumberId(Number_id)
	if err!= nil {
		result["numberdata"]= ""
	} else {
		result["numberdata"]=numberdata
	}
		//获取对应老人的基本资料
	senior,err:= models.SelectSeniorbynumberid(Number_id)
	if err!=nil {
		result["senior"]=""
	} else {
		result["senior"]=senior
	}
	//获取信息和联系人的infomation基本资料
	infomation,err:= models.SelectInformationNumberId(Number_id)
	if err!=nil {
		result["infomation"]=""
	} else {
		result["infomation"]=infomation
	}
		//获取信息和联系人的health健康基本资料
	health,err:= models.SelectHealthnumberid(Number_id)
	if err!=nil {
		result["health"]=""
	} else {
		result["health"]=health
	}
	//获取信息和联系人的healthrelate健康相关问题
	healthrelate,err:= models.SelectHealthrelatednumberid(Number_id)
	if err!=nil {
		result["healthrelate"]=""
	} else {
		result["healthrelate"]=healthrelate
	}
	//获取获取量表分题目的答案
		search := &models.Score_record{
		Number_id:    Number_id,
		Seniorid:   senior.Id,
		Assessors_id: user.Id,
	}
	treelist := []*Score_recordjson{}
	record := models.GetScore_recordList(1000,1,search)
	if record!=nil {
		var kname,kpiinfotitle string
		for _, v := range record {
			kpidata,_ := models.SelectkpiById(v.Kpi_id)
			if kpidata!=nil {
				kname= kpidata.Kname
			} else {
				kname = ""
			}
			kpiinfodata,_ := models.SelectkpiinfoById(v.Kpiinfo_id)
			if kpiinfodata!=nil {
				kpiinfotitle= kpiinfodata.Title
			} else {
				kpiinfotitle= ""
			}
			node := &Score_recordjson{
			Id:       v.Id,
			Seniorid:      v.Seniorid,
			Assessorsid:      v.Assessors_id,
			Numberid:      v.Number_id,
			Kpiid: v.Kpi_id,
			Kname:    kname,
			Kpiinfoid:    v.Kpiinfo_id,
			Kpiinfotitle:   kpiinfotitle,
			Score:   v.Score,
		}
		// node.Children = infolistdata
		treelist = append(treelist, node)

		}
		result["kpiarr"]=treelist
	}
			c.JSON(200, gin.H{
			"code": 200,
			"msg":  "数据返回成功！",
			"data": result,
		})

}