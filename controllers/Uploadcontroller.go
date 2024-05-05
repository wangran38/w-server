package controllers

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Uploadmp4(c *gin.Context) {
	// 单文件上传

	// 解析 multipart 表单
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取文件上传出错",
			"data":    "",
		})
		return
	}
	// 2.获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
		".JPG":  true,
		".PNG":  true,
		".GIF":  true,
		".JPEG": true,
		".mp4":  true,
	}

	_, ok := allowExtMap[extName]
	if !ok {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "文件不合法",
			"data":    extName,
		})
		return
	}

	// 3.创建图片保存目录
	day := GetDay()

	// 文件保存目录
	dir := path.Join("./static/upload", day)

	if err := os.MkdirAll(dir, 0666); err != nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "文件目录创建失败",
			"data":    "",
		})
		return
	}

	// 4.生成保存目录，以及文件名
	// 生成文件名称，防止重名的覆盖
	fileName := strconv.FormatInt(GetUnix(), 10) + extName

	dst := path.Join(dir, fileName)

	// 保存图片
	c.SaveUploadedFile(file, dst)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "文件上传成功！",
		"data":    dst,
	})
	// return
}

// 获取当前的日期

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}
