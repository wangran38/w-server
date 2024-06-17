package routers

import (
	"fmt"
	"net/http"
	"strings"
	"w-server/apic"
	"w-server/controllers"
	_ "w-server/models"
	"w-server/utils"

	"github.com/gin-gonic/gin"
	//"github.com/unrolled/secure"
)

func init() {

	router := gin.Default()
	router.Use(Cors())
	// 版本v1
	router.POST("admin/login", controllers.LoginController) //登录
	admin := router.Group("/admin").Use(Loginhead())
	{
		admin.POST("/logout", controllers.Loginout) //登录
		admin.POST("/add", controllers.AddAdmin)
		admin.POST("/del", controllers.Deladmin)
		admin.POST("/edit", controllers.EditAdmin)
		admin.POST("/getinfo", controllers.GetAdminInfo)
		admin.POST("/getrule", controllers.GetAdminRule)
		admin.POST("/getadminlist", controllers.GetAdminlist)
		//用户组别接口
		admin.POST("/getgrouplist", controllers.Getgrouplist)
		admin.POST("/getgrouptree", controllers.TreeGroup)
		admin.POST("/addgroup", controllers.AddGroup)
		admin.POST("/editgroup", controllers.EditGroup)
		admin.POST("/delgroup", controllers.Delgroup)
		//菜单接口
		admin.POST("/getruleslist", controllers.Getruleslist)
		admin.POST("/delrules", controllers.DelRules)
		admin.POST("/addrules", controllers.AddRules)
		admin.POST("/editrules", controllers.EditRules)
		admin.POST("/getallrule", controllers.GetAllRule)
		//城市接口
		admin.POST("/citylist", controllers.Getcitylist)

		//分类接口
		admin.POST("/categorylist", controllers.Getcategorylist)
		admin.POST("/addcategory", controllers.AddCategory)
		admin.POST("/editcategory", controllers.EditCategory)
		admin.POST("/delcategory", controllers.DelCategory)
		admin.POST("/treecategory", controllers.TreeCategory)
		//指标接口
		admin.POST("/kpilist", controllers.GetKpilist)
		admin.POST("/addkpi", controllers.AddKpi)
		admin.POST("/editkpi", controllers.EditKpi)
		admin.POST("/delkpi", controllers.DelKpi)
		admin.POST("/treekpi", controllers.TreeKpi)
		//指标题库接口
		admin.POST("/kpiinfolist", controllers.GetKpiinfolist)
		admin.POST("/addkpiinfo", controllers.AddKpiinfo)
		admin.POST("/editkpiinfo", controllers.EditKpiinfo)
		admin.POST("/delkpiinfo", controllers.DelKpiinfo)
		// admin.POST("/treekpi", controllers.TreeKpi)
		// //开启tcp服务
		//指标接口
		admin.POST("/assessorsclist", controllers.GetAssessorslist)
		admin.POST("/addassessorsc", controllers.AddAssessors)
		// admin.POST("/ediassessorsc", controllers.EditKpi)
		admin.POST("/delassessorsc", controllers.DelAssessors)
		// admin.POST("/treekpi", controllers.TreeKpi)
		// admin.POST("/opentcp", controllers.Opentcp)
		//新闻文章接口
		admin.POST("/newslist", controllers.Getnewslist)
		admin.POST("/addnews", controllers.AddNews)
		admin.POST("/editnews", controllers.EditNews)
		admin.POST("/delnews", controllers.DelNews)

		//查找企业资源数字化文章接口
		admin.POST("/getbooklist", controllers.GetBooklist)
		//添加企业资源数字化文章接口
		admin.POST("/addbook", controllers.Addbook)
		//修改企业资源数字化文章接口
		admin.POST("/editbook", controllers.EditBook)
		//删除企业资源数字化文章接口
		admin.POST("/delbook", controllers.DelBook)

		//查找企业资源数字化文章tree接口 （前为定义给前端访问的接口路径）（controlloers.调的是方法名）
		admin.POST("/getbookinfolist", controllers.GetBookinfolist)
		//添加企业资源数字化文章tree接口
		admin.POST("/addbookinfo", controllers.Addbookinfo)
		//修改企业资源数字化文章tree接口
		admin.POST("/editbookinfo", controllers.EditBookinfo)
		//删除企业资源数字化文章tree接口
		admin.POST("/delbookinfo", controllers.DelBookinfo)
		//删除企业资源数字化文章tree接口
		admin.POST("/treebook", controllers.TreeBook)

		//字典接口
		admin.POST("/dictionary", controllers.GetDictionarylist) //查询
		admin.POST("/adddictionary", controllers.AddDictionary)  //添加
		//admin.POST("/updictionary", controllers.Updedictionary)  //修改
		admin.POST("/deldictionary", controllers.DelDictionary) //删除

		//文件上传接口
		admin.POST("/mp4_upload", controllers.Uploadmp4) //提交文件上传
		//护理动作文章视频接口
		admin.POST("/getbookfilelist", controllers.GetBookfilelist)
		//添加护理动作文章视频接口
		admin.POST("/addbookfile", controllers.AddBookfile)
		//修改护理动作文章视频接口
		admin.POST("/editbookfile", controllers.EditBookfile)
		//删除护理动作文章视频接口
		admin.POST("/delbookfile", controllers.DelBookfile)

		//添加护理动作接口
		admin.POST("/addnursing", controllers.AddNursing)
		//修改护理动作接口
		admin.POST("/upnursing", controllers.UpNursing)
		//删除护理动作接口
		admin.POST("/deletenursing", controllers.DeleteNursing)

	}
	api := router.Group("/api")
	{
		api.POST("/getkpi", apic.Getkpilist)                //Kpi指标的API前端接口
		api.POST("/getdictionary", apic.GetDictionaryclist) //居住情况的API前端接口
		api.POST("/cglist", apic.Getcategorylist)
		// api.POST("/getinformation", apic.GetInformationlist) //登录


		api.POST("/getbook", apic.Getbooklist)  //企业资源数字化的API前端接口
		api.POST("/citytree", apic.Treecity)    //
		api.POST("/newslist", apic.GetNewslist) //
		api.POST("/newsinfo", apic.GetNewsinfo) //
		//评估员注册登录
		api.POST("/assessors_rg", apic.Rsassessors) //
		api.POST("/assessors_login", apic.Loginassessors)

		api.POST("/selectnursing", apic.SelectNursinglist) //护理动作接口
	}
	api.Use(Loginassessorschead())
	{
		//评估员操作
		api.POST("/add_numbercode", apic.AddNumberc)        //添加编号
		api.POST("/edit_numbercode", apic.EditNumberc)      //添加编号
		api.POST("/my_number", apic.MyNumberc)              //添加编号
		api.POST("/my_numberinfo", apic.MyNumberinfo)              //添加编号
		
		api.POST("/add_seniorc", apic.Addseniorc)           //添加老人信息
		api.POST("/my_seniorc", apic.MySeniorc)             //添加老人信息
		api.POST("/up_seniorc", apic.Upseniorc)             //修改老人信息
		//添加选择的量表记录
		api.POST("/add_score_record", apic.AddScore_record) //添加老人信息
		api.POST("/my_score_record", apic.Myscorerecord) //添加老人信息
		api.POST("/my_kpi_record", apic.Mykpirecord) //添加老人信息
		//api.POST("/gethealth", apic.GetHealthlist)               //健康

		api.POST("/addinformation", apic.AddInformation) //信息提供者及联系人接口
		api.POST("/myinformation", apic.MyInformation)   //信息提供者及联系人编号接口
		api.POST("/upinformation", apic.UpInformation)   //修改信息提供者及联系人编号接口

		api.POST("/addhealth", apic.AddHealth) //添加疾病接口
		api.POST("/myhealth", apic.MyHealth)   //疾病编号接口
		api.POST("/uphealth", apic.UpHealth)   //修改疾病接口

		api.POST("/addhealthrelated", apic.AddHealthrelated) //健康相关信息
		api.POST("/myhealthrelated", apic.MyHealthrelated)   //健康相关问题接口
		api.POST("/uphealthrelated", apic.UpHealthrelated)   //修改健康相关问题接口

	}

	// //开启TCP服务结束
	router.Run(":8189")
}

// //// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json,multipart/form-data,application/x-www-form-urlencoded")                                                                                                        // 设置返回格式是json
		}

		//放行所有OPTIONS方法

		//放行所有OPTIONS方法,防止vue握手2次的问题
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next() //  处理请求
	}
}

func Loginhead() gin.HandlerFunc { //是否登录权限
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token != "" || len(token) != 0 {
			//   fmt.Println(token)
			user := utils.GetLoginUser(token)
			if user.Id <= 0 {
				c.JSON(201, gin.H{
					"code":    201,
					"message": "你没有权限！",
					"data":    "",
					// "permissions": menu,
					// "roles":       role,
				})
				c.Abort()
				return
			}

		} else {
			c.JSON(201, gin.H{
				"code":    201,
				"message": "你没有登录！",
				"data":    "",
				// "permissions": menu,
				// "roles":       role,
			})
			c.Abort()
			return
		}
	}
}
func Loginassessorschead() gin.HandlerFunc { //是否登录权限
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token != "" || len(token) != 0 {
			_, finderr := utils.GetLoginAssessorsc(token)
			if finderr != nil {
				c.JSON(201, gin.H{
					"code":    201,
					"message": "你没有权限！",
					"data":    "",
					// "permissions": menu,
					// "roles":       role,
				})
				// c.Abort()
				return
			}

		} else {
			c.JSON(201, gin.H{
				"code":    201,
				"message": "你没有登录！",
				"data":    "",
				// "permissions": menu,
				// "roles":       role,
			})
			c.Abort()
			return
		}
	}
}
