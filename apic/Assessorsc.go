package apic

import (
	// "fmt"
	// "net/http"
	"net/http"
	"strconv"
	"time"
	_ "time"
	"w-server/lib"
	"w-server/models"
	"w-server/utils"

	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Assessorlogin struct {
	Id    int64  `json:"id"`
	Phone string `json:"phone"`
	Psw   string `json:"psw"`
	Psw1  string `json:"psw1"`
}

func Rsassessors(c *gin.Context) {
	var logindata Assessorlogin
	c.ShouldBind(&logindata)
	// fmt.Print(searchinfo.Id)
	// result := make(map[string]interface{})
	if logindata.Phone == "" {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "手机号码不能为空",
			// "data": "",
		})
		return
	}
	if logindata.Psw == "" {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "密码不能为空",
			// "data": logindata.Phone,
		})
		return
	}
	info, _ := models.SelectAssessorsbyphone(logindata.Phone)
	if info != nil {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "对不起，当前手机号码已经被注册",
			"data": logindata.Phone,
		})
		return
	} else {

		if logindata.Psw != logindata.Psw1 {
			c.JSON(200, gin.H{
				"code": 201,
				"msg":  "对不起，两次输入的密码不正确",
				// "data": logindata.Phone,
			})
			return
		} else {
			// //读取数据库
			pwd, salt := lib.Password(4, logindata.Psw) //截取四位随机盐+上这个做原始密码
			Assessors := new(models.Assessors)
			Assessors.Phone = logindata.Phone
			Assessors.Password = pwd
			Assessors.Salt = salt
			Assessors.Created = time.Now()
			err := models.AddAssessors(Assessors) //判断账号是否存在！
			if err != nil {
				c.JSON(200, gin.H{
					"code": 201,
					"msg":  "添加数据出错！",
					"data": err,
				})
				return
			} else {
				Assessorsaccess := new(models.Assessorsaccess)
				Assessorsaccess.Uid = Assessors.Id
				Assessorsaccess.Gid = 1
				err := models.AddAssessorsaccess(Assessorsaccess) //添加关系表
				if err == nil {
					assessorsid := strconv.FormatInt(Assessors.Id, 10) //转换字符串
					token := utils.CreateAssessorscToken(assessorsid)
					c.JSON(http.StatusOK, gin.H{
						"code": 200,
						"msg":  "数据添加成功！",
						"data": token,
					})
					return
				} else {
					c.JSON(http.StatusOK, gin.H{
						"code": 201,
						"msg":  "数据添加失败！",
						"data": err,
					})

				}

			}
			//Assessors.Password = Assessors.psw
			// Assessors.Phone = Assessors.Phone
		}
		// c.JSON(200, gin.H{
		// 	"code": 201,
		// 	"msg":  "获取数据失败",
		// 	"data": logindata.Id,
		// })
	}
}
func Loginassessors(c *gin.Context) {
	var logindata Assessorlogin
	c.ShouldBind(&logindata)
	// fmt.Print(searchinfo.Id)
	// result := make(map[string]interface{})
	if logindata.Phone == "" {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "手机号码不能为空",
			"data": logindata.Phone,
		})
		return
	}
	info, _ := models.SelectAssessorsbyphone(logindata.Phone)
	if info == nil {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "对不起，该用户名无效",
			"data": logindata.Phone,
		})
		return
	} else {
		pwd := lib.Md5([]byte(logindata.Psw + "988cj.com" + info.Salt))
		if pwd != info.Password {
			c.JSON(http.StatusOK, gin.H{
				"code":    "201",
				"message": "密码不正确！",
				// "pwd":  pwd,
			})
			return

		}
		id := strconv.FormatInt(info.Id, 10) //转换字符串
		token := utils.CreateAssessorscToken(id)
		result := make(map[string]interface{})
		result["token"] = token
		c.JSON(200, gin.H{
			"code":    200,
			"message": "登录成功！",
			// "token":   token,
			"data": result,
		})

	}
}
