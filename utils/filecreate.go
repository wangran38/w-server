package utils

import (
	_"crypto/md5"
	"fmt"
	"math/rand"
	"time"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
    "strings"
	"os"
)

// 下载图片信息
func DownLoad(base string, url string) string {
    pic := base
    idx := strings.LastIndex(url, "/")
    if idx < 0 {
        pic += "/" + url
        //fmt.Printf(pic)
    } else {
        pic += "/" + url[idx+1:]
        fmt.Printf(pic)
        fmt.Printf("-----------------")
    }
    v, err := http.Get(url)
    if err != nil {
        fmt.Printf("Http get [%v] failed! %v", url, err)
        return "not ok"
    }
    defer v.Body.Close()
    content, err := ioutil.ReadAll(v.Body)
    if err != nil {
        fmt.Printf("Read http response failed! %v", err)
        return "not ok"
    }
    fmt.Printf(pic)
    err = ioutil.WriteFile(pic, content, 0666)
    if err != nil {
        fmt.Printf("Save to file failed! %v", err)
        return "not ok"
    }
    fmt.Printf("ok")
    return "ok"
}

//保存文件操作（如果程序被强行结束，下载的某些图片可能不完整）
func saveFile(finalUrl string, savePath string){
	//判断文件是否存在，如果存在，说明已经下载了，直接return，下载下一张图片
	_, err := os.Stat(savePath)
	if err == nil {
		fmt.Printf("图片 [ %s ] 已下载\n", finalUrl)
		return
	}

	//创建文件：如果文件已存在，会将文件清空
	file, err := os.Create(savePath)
	if err != nil {
		fmt.Println("createFile error:", err.Error())
		return
	}
	defer file.Close()

	//读取url的信息，存入到文件
	resp, err := http.Get(finalUrl)
	if err != nil {
		fmt.Println("imageGet error:", err.Error())
		return
	}
	defer file.Close()

	buf := make([]byte, 4096)
	for {
		res, err2 := resp.Body.Read(buf)
		if res == 0{
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//写入文件
		file.Write(buf[:res])
	}

	//循环结束，我们认为图片已经下载成功，控制台输出提示
	fmt.Printf("图片 [ %s ] 下载 [ %s ] 成功\n", finalUrl, savePath)
}

//生成随机字符串
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
//下载网络图片到本地
func Gethttpimg(openimgurl string, staticurl string) (ms string) {
	randstr := GetRandomString(4)
    timename := time.Now().Unix()
    filename := randstr + strconv.FormatInt(timename,10)
    resp, _ := http.Get(openimgurl)//打开链接
	body, _ := ioutil.ReadAll(resp.Body) //读取文件流
	out, _ := os.Create(staticurl + filename +".jpg") //生成新的图片文件
	io.Copy(out, bytes.NewReader(body))
	imgurl := staticurl + filename +".jpg"
	return string(imgurl)
}