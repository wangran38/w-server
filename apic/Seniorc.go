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

type Seniorc struct {
	Id              int64  `json:"id"` //ID自增长
	Assessor_id     int64  `json:"assessor_id"`
	Number_id       int64  `json:"number_id"`
	Senior_name     string `json:"senior_name"`     //姓名
	Senior_gender   int    `json:"senior_gender"`   //性别
	Senior_dob      string `json:"senior_dob"`      //出生日期
	Birthday        int    `json:"birthday"`        //出生日期时间戳
	Height          int    `json:"height"`          //身高
	Weight          int    `json:"weight"`          //体重
	Ethnic          string `json:"ethnic"`          //民族
	Religion        string `json:"religion"`        //宗教
	Idnumber        string `json:"idnumber"`        //身份证号
	Education_level string `json:"education_level"` //文化程度
	Live_way        string `json:"live_way"`        //居家方式
	Is_marriage     string `json:"is_marriage"`     //婚姻状况
	// Ssnumber      string `xorm:"varchar(64)" json:"ssnumber"` //社保号码
	Payment   string `json:"payment"`   //医疗费用支付方式
	Financial string `json:"financial"` //经济来源

}

// //添加用户组
func Addseniorc(c *gin.Context) {
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
	var formdata Seniorc
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Senior)
	Intodata.Id = formdata.Id
	Intodata.Assessor_id = user.Id
	Intodata.Senior_gender = formdata.Senior_gender
	Intodata.Senior_dob = formdata.Senior_dob
	Intodata.Birthday = formdata.Birthday
	Intodata.Height = formdata.Height
	Intodata.Weight = formdata.Weight
	Intodata.Ethnic = formdata.Ethnic
	Intodata.Religion = formdata.Religion
	Intodata.Idnumber = formdata.Idnumber
	Intodata.Education_level = formdata.Education_level
	Intodata.Live_way = formdata.Live_way
	Intodata.Is_marriage = formdata.Is_marriage
	Intodata.Payment = formdata.Payment
	Intodata.Financial = formdata.Financial
	Intodata.Created = time.Now()
	// info, _ := models.SelectSeniorbyidnum(Intodata.Idnumber) //判断账号是否存在！
	// if info != nil {
	// 	c.JSON(200, gin.H{
	// 		"code": "201",
	// 		"msg":  "该老人身份证已经存在！",
	// 	})
	// 	return
	// }
	err := models.AddSenior(Intodata) //判断账号是否存在！
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

// 我的编号
func MySeniorc(c *gin.Context) {
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
