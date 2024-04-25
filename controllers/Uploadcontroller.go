package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func Chuan(c *gin.Context) {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// 设置请求表单最大内存限制,默认是 30MB
	// internal call the ParseMultipartForm method of the http request, which requires passing in a number of bytes,
	// to get the data of the MultipartForm field, you need to use the ParseMultipartForm() method to parse the form,
	// and when parsing, you need to specify the maximum number of bytes stored in memory, and the remaining bytes will be saved in the temporary disk file
	maxMultipartMemory := int64(8 << 20)
	log.Printf("The maximum number of bytes for parsing files to memory: %d", maxMultipartMemory)
	router.MaxMultipartMemory = maxMultipartMemory // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")                       //FormFile from the form to return the first matching file object (structure)
		log.Printf("Obtained file name: %s", file.Filename) //The file name must be safe and reliable, and only the file name can be retained by removing the path information
		// Upload the file to specific dst.
		// currentPath, _ := os.Getwd() //Get the current file path
		// 3. Create a directory for saving images
		today := time.Now().Format("2006_01_02") // Change to the format of year_month_day
		dir := path.Join("./static/upload", today)
		if err := os.MkdirAll(dir, 0666); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": "Directory creation failed",
			})
		}

		// dst := currentPath + "/" + file.Filename
		log.Printf("Saved file absolute path: %s", dir)
		c.SaveUploadedFile(file, dir)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}

// Get the current date
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// Simulate single file upload:
// curl -X POST http://localhost:8080/upload  -H "Content-Type: multipart/form-data" -F "file=@file name"
