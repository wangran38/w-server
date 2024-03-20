package controllers
import (
	// "time"
	// "strconv"
	// "strings"
    "encoding/json"
	"fmt"
	"log"
	// "linfeng/models"
	// "linfeng/utils"
    "io/ioutil"
	"net/http"
	// "regexp"
	// "ginstudy/utils"
	"github.com/gin-gonic/gin"
)
const (
    akkey = "HuGhPC4aBKeErz7BYfwM5GU1pkwxHmrV"
)
type Request_config struct {
	Status               int        `json:"status"`
	Result              [] Result `json:"result"`
}
type Result struct {
    Level string       `json:"level"`
    Lc              []Location `json:"location"`
}
type Location struct {
	Lat               int        `json:"lat"`
	Lng              int `json:"lng"`
}
func Allplace(c *gin.Context) {
    openurl := "https://api.map.baidu.com/geocoding/v3/?address=%E5%A4%A9%E6%B4%A5%E5%B8%82%E6%B4%A5%E5%8D%97%E5%8C%BA%E5%9B%BD%E5%B1%95%E5%A4%A7%E9%81%93888%E5%8F%B7&output=json&ak="+akkey
    resp, _ := http.Get(openurl)
	// resp, _ := http.Get("https://www.onezh.com/news/23201.html")
    defer resp.Body.Close() //go的特殊语法，main函数执行结束前会执行 resp.Body.Close()
        if resp.StatusCode == http.StatusOK { //如果响应码为200
        body, err := ioutil.ReadAll(resp.Body) //把响应的body读出
        if err != nil {                        //如果有异常
            fmt.Println(err) //把异常打印
            log.Fatal(err)   //日志
}
// fmt.Println(string(body))
// jsondata := []byte(str) //序列化json
// var stu Actor
// json.Unmarshal(abc,&stu) //传入结构体
	var config Request_config
    //config := map[string]interface{}{}
	jsonerr := json.Unmarshal([]byte(body), &config)

	if jsonerr != nil {
		fmt.Println(jsonerr)
	}
    
	//location_info := config["result"]
    // for _, v := range config.Result {
    //     fmt.Println(v)
    // }
 //fmt.Println("%+v\n",config.Result)
	// log.Info(config)
	// log.Info("%T %v", config, config)

	//fmt.Println(config["result"])
    fmt.Println(config.Result)

}
}