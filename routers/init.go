package routers

import (
	"fmt"
	"net/http"
	"strings"
	"changxiaoyang/utils"
	"changxiaoyang/controllers"
	"changxiaoyang/apic"
	"changxiaoyang/htmlc"
	_ "changxiaoyang/models"

	"github.com/gin-gonic/gin"
	//"github.com/unrolled/secure"
)

func init() {

	router := gin.Default()
	    //加载模版
    //router.LoadHTMLGlob("./view/*")
	router.LoadHTMLGlob("./view/**/*")
	router.StaticFS("/static", http.Dir("./static"))
	//router.StaticFile("/MP_verify_4RLm31WaN7toWx0O.txt", "./MP_verify_4RLm31WaN7toWx0O.txt") //固定根目录微信文件可以访问
	// router.GET("/news_list/:page/:limit/:cagegoryid", apic.GetNewslist)
	//router.GET("/news_list/:page/:limit/:cagegoryid", htmlc.Newslist)
	//router.GET("/wx/login", htmlc.Pcindex)
	router.GET("/ex_list", htmlc.Pczhanhui)
	router.Use(Cors())
	// 版本v1
	router.POST("admin/login", controllers.LoginController) //登录
	admin := router.Group("/admin").Use(Loginhead())
	{		
		admin.POST("/logout", controllers.Loginout)       //登录
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
		//海外城市接口
		admin.POST("/countrylist", controllers.Getcountrylist)
		admin.POST("/addcountry", controllers.Addcountry)
		admin.POST("/editcountry", controllers.Editcountry)
		admin.POST("/delcountry", controllers.Delcountry)
		admin.POST("/treecountry", controllers.Treecountry)
		//分类接口
		admin.POST("/categorylist", controllers.Getcategorylist)
		admin.POST("/addcategory", controllers.AddCategory)
		admin.POST("/editcategory", controllers.EditCategory)
		admin.POST("/delcategory", controllers.DelCategory)
		admin.POST("/treecategory", controllers.TreeCategory)
		//展会接口
		// admin.POST("/marketlist", controllers.GetMarketlist)
		// admin.POST("/addmarket", controllers.AddMarket)
		// admin.POST("/editmarket", controllers.EditExhibition)

		// //开启tcp服务
		// admin.POST("/opentcp", controllers.Opentcp)
		//新闻文章接口
		admin.POST("/newslist", controllers.Getnewslist)
		admin.POST("/addnews", controllers.AddNews)
		admin.POST("/editnews", controllers.EditNews)
		admin.POST("/delnews", controllers.DelNews)
		//支付接口

	}
		api := router.Group("/api")
	{
api.POST("/cglist", apic.Getcategorylist) //登录
api.POST("/cgtree", apic.Getcategorytree) //登录
// api.POST("/marketlist", apic.Getmarketlist) //登录
// api.POST("/marketinfo", apic.Getmarketinfo) //登录
// api.POST("/exlist", apic.GetExlist) //登录
// api.POST("/exinfo", apic.GetExinfo) //登录
// api.POST("/eximginfo", apic.GetEximginfo) //登录
api.POST("/countrylist", apic.Getcountrylist) //登录
api.POST("/countryall", controllers.Treecountry) //登录
api.POST("/citylist", apic.Getcitylist) //登录
api.POST("/citytree", apic.Treecity) //登录
api.POST("/newslist", apic.GetNewslist) //登录
api.POST("/newsinfo", apic.GetNewsinfo) //登录
// api.GET("/wx/ckwx", apic.CkSign) //微信基础token
// api.GET("/wx/getopenid", apic.Getopenid) //微信基础token
// //支付接口
// api.POST("pay/codeurl", apic.Jsppay) //微信基础token
// api.POST("pay/get/cert", apic.V3cert) //微信基础token

//头部加密
//api.GET("/wx/get", apic.Getopenid) //微信基础token
	}
// 		wx := router.Group("/wx")
// 	{
// wx.POST("/huidiao", apic.Huidiao) //微信基础token
// 	}
    //router.Use(TlsHandler())

    //router.RunTLS(":8088", "./9963857_server.szhfair.com_nginx/9963857_server.szhfair.com.pem", "./9963857_server.szhfair.com_nginx/9963857_server.szhfair.com.key")

	//开启TCP服务啦
	// listen, err := net.Listen("tcp", "127.0.0.1:20000")
	// if err != nil {
	// 	fmt.Println("listen failed, err:", err)
	// 	return
	// }
	// for {
	// 	conn, err := listen.Accept() // 建立连接
	// 	if err != nil {
	// 		fmt.Println("accept failed, err:", err)
	// 		continue
	// 	}
	// 	go process(conn) // 启动一个goroutine处理连接
	// }
	// //开启TCP服务结束
	router.Run(":8088")
}
// func TlsHandler() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         secureMiddleware := secure.New(secure.Options{
//             SSLRedirect: true,
//             SSLHost:     "47.97.11.25:8088",
//         })
//         err := secureMiddleware.Process(c.Writer, c.Request)

//         // If there was an error, do not continue.
//         if err != nil {
//             return
//         }

//         c.Next()
//     }
// }

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
		// if method == "OPTIONS" {
		// 	c.JSON(http.StatusOK, "Options Request!")
		// }
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
		if user.Id<=0 {
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