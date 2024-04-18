package apic

import (
	// "fmt"
	// "net/http"

	"fmt"
	"time"
	_ "time"
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
	Kpiinfoid   int64 `json:"kpiinfo_id"`   //KPI的id
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
		fmt.Printf("Number: %d\n", item.Seniorid)
		Intodata.Created = time.Now()
		// sum++
		err := models.AddScore_record(Intodata) //判断账号是否存在！
		if err == nil {
			sum++
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
		"msg":  "添加数据总共！",
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
