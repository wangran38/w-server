package apic

import (
	// "fmt"
	// "net/http"

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
	user := utils.GetLoginAssessorsc(token)
	//也可继承controller里的结构体
	var formdata Numberc
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Number)

	Intodata.Id = formdata.Id
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
