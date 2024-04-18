package controllers

import (
	"net/http"
	"time"
	"w-server/lib"
	"w-server/models"

	// "fmt"
	// "strconv"
	"github.com/gin-gonic/gin"
)

type Assessorsserch struct {
	Id       int    `json:"id"`
	Username string `json:"title"`
	Phone    string `json:"phone"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Order    string `json:"sort"`
}

//	type Assessorsform struct {
//		Username string `form:"username" binding:"required" json:"username"`
//		Nickname string `form:"nickname" binding:"required" json:"nickname"`
//		Phone string `form:"phone" binding:"required" json:"phone"`
//		Email string `form:"email" binding:"required" json:"email"`
//		Avatar string `form:"avatar" binding:"required" json:"avatar"`
//		Gid int64 `form:"gid" binding:"required" json:"gid"`
//		Password string `form:"password" binding:"required" json:"password"`
//	}
//
//	type Assessorsserch struct {
//		ID int `json:"id"`
//		Username string `json:"title"`
//		Limit int `json:"limit"`
//		Page int `json:"page"`
//		Order string `json:"sort"`
//	}
//
//	type AssessorsController struct {
//		BaseController //继承统一判断是否登录或者是否有权限类
//		// beego.Controller
//	}
func AddAssessors(c *gin.Context) {
	var Assessorsdata models.Assessorsjson
	if err := c.BindJSON(&Assessorsdata); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"msg":     "表单未完整！",
			"msgcode": "-1",
			// "data":    reqIP,
		})
		return
	}
	// //读取数据库
	Assessors := new(models.Assessors)
	Assessors.Username = Assessorsdata.Username
	Assessors.Nickname = Assessorsdata.Nickname
	Assessors.Phone = Assessorsdata.Phone
	Assessors.Avatar = Assessorsdata.Avatar
	Assessors.Email = Assessorsdata.Email
	info, _ := models.SelectUserByUserName(Assessors.Username) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "该用户已经存在！",
		})
		return
	}
	pwd, salt := lib.Password(4, Assessorsdata.Password) //截取四位随机盐+上这个做原始密码
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
		Authaccess := new(models.Authaccess)
		Authaccess.Uid = Assessors.Id
		Authaccess.Gid = Assessorsdata.Gid
		err := models.AddAuthaccess(Authaccess) //判断账号是否存在！
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "数据添加成功！",
				"data": "",
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
}

// 修改
func EditAssessors(c *gin.Context) {
	var formdata models.Assessorsjson
	c.BindJSON(&formdata)
	intodata := new(models.Assessors)
	intodata.Id = formdata.Id
	intodata.Username = formdata.Username
	intodata.Nickname = formdata.Nickname
	intodata.Phone = formdata.Phone
	intodata.Email = formdata.Email
	intodata.Avatar = formdata.Avatar
	if intodata.Password != "" {
		pwd, salt := lib.Password(4, intodata.Password) //截取四位随机盐+上这个做原始密码
		intodata.Password = pwd
		intodata.Salt = salt
	}

	if formdata.Id <= 0 {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res, err := models.UpAssessors(intodata) //判断账号是否存在！
		if err != nil {
			c.JSON(201, gin.H{
				"code": 201,
				"msg":  "修改数据出错！",
				"data": err,
			})
			return
		} else {
			// var formdata1 models.Assessorsjson
			// c.ShouldBind(&formdata1)
			updata := new(models.Authaccess)
			updata.Uid = formdata.Id
			updata.Gid = formdata.Gid
			_, err := models.UpAuthaccess(updata) //判断账号是否存在！
			if err == nil {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "数据修改成功！",
					"data": res,
				})
			}

		}
	}

}
func GetAssessorslist(c *gin.Context) {
	var searchdata Assessorsserch
	// c.BindJSON(&searchdata)
	c.ShouldBind(&searchdata)
	result := make(map[string]interface{})
	limit := searchdata.Limit
	page := searchdata.Page
	// username := searchdata.Username
	phone := searchdata.Phone
	order := searchdata.Order
	listdata := models.GetAssessorsList(limit, page, phone, order)
	listnum := models.GetAssessorstotal(phone)

	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取菜单失败1",
			"data":    "",
		})
		return
	} else {
		result["listdata"] = listdata
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功1",
			"data":    result,
		})
		return
	}
}

func DelAssessors(c *gin.Context) {
	var searchdata models.Assessorsjson
	c.BindJSON(&searchdata)
	delnum := models.DeleteAssessors(searchdata.Id)
	if delnum > 0 {
		del2 := models.DeleteAuthaccess(searchdata.Id)
		if del2 > 0 {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "删除成功！",
				"data":    del2,
			})
		} else {
			c.JSON(201, gin.H{
				"code":    201,
				"message": "删除用户关系失败！",
				"data":    del2,
			})
		}

	} else {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "删除失败！",
			"data":    delnum,
		})
	}

}
