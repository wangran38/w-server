package lib

import (
    "context"
	"log"
	"time"
    "fmt"
	"strconv"
    "io/ioutil"
     "crypto/rsa"
	 "errors"
	"crypto/x509"
	"encoding/pem"
	 "crypto"
    "crypto/rand"
	"crypto/sha256"
    // "errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
		"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	//"github.com/wechatpay-apiv3/wechatpay-go/utils"
)
func JsapiApiService_Prepay() {
	var (
		mchID                      string = "1642420641"                               // 商户号
		mchCertificateSerialNumber string = "49F994F3D4860DC59E27D68F25D2FA6888D7421E" // 商户证书序列号
		mchAPIv3Key                string = "zoJkIWrZSp9GKiAuWAiwrKA8p2xPl9j8"         // 商户APIv3密钥
	)
    //fmt.Printf("testwtewewrrw")
    // 	file := "./keypem/test.txt"
	// content,err := ioutil.ReadFile(file)
	// if err != nil{
	// 	fmt.Println("read file error:",err)
	// }
	// fmt.Println(string(content))
    // return
    file := "./keypem/apiclient_key.pem"
	privateKeyBytes, err := ioutil.ReadFile(file)
    fmt.Println(string(privateKeyBytes))
	if err != nil {
		fmt.Printf("打开文件出错！！")
        return
	}
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, werr := LoadPrivateKey(string(privateKeyBytes))
	if werr != nil {
		log.Print("加载key错误")
        //fmt.Printf("找不到文件没目录")
        //return
	}
fmt.Println(mchPrivateKey)
	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
    
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		jsapi.PrepayRequest{
			Appid:         core.String("wxceff67ea8eaa04fa"),
			Mchid:         core.String("1642420641"),
			Description:   core.String("Image形象店-深圳腾大-QQ公仔"),
			OutTradeNo:    core.String("12177525012014070332333680338"),
			TimeExpire:    core.Time(time.Now()),
			Attach:        core.String("自定义数据说明"),
			NotifyUrl:     core.String("https://www.weixin.qq.com/wxpay/pay.php"),
			GoodsTag:      core.String("WXG"),
			LimitPay:      []string{"LimitPay_example"},
			SupportFapiao: core.Bool(false),
			Amount: &jsapi.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(100),
			},
			Payer: &jsapi.Payer{
				Openid: core.String("oNbgG5-1NUNuS6nYDZWA8DLmR2LE"),
			},
			Detail: &jsapi.Detail{
				CostPrice: core.Int64(608800),
				GoodsDetail: []jsapi.GoodsDetail{jsapi.GoodsDetail{
					GoodsName:        core.String("iPhoneX 256G"),
					MerchantGoodsId:  core.String("ABC"),
					Quantity:         core.Int64(1),
					UnitPrice:        core.Int64(828800),
					WechatpayGoodsId: core.String("1001"),
				}},
				InvoiceId: core.String("wx123"),
			},
			SceneInfo: &jsapi.SceneInfo{
				DeviceId:      core.String("013467007045764"),
				PayerClientIp: core.String("14.23.150.211"),
				StoreInfo: &jsapi.StoreInfo{
					Address:  core.String("广东省深圳市南山区科技中一道10000号"),
					AreaCode: core.String("440305"),
					Id:       core.String("0001"),
					Name:     core.String("腾讯大厦分店"),
				},
			},
			SettleInfo: &jsapi.SettleInfo{
				ProfitSharing: core.Bool(false),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
// LoadPrivateKey 通过私钥的文本内容加载私钥
func LoadPrivateKey(privateKeyStr string) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, fmt.Errorf("decode private key err")
	}
	if block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("the kind of PEM should be PRVATE KEY")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key err:%s", err.Error())
	}
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not a RSA private key")
	}
	return privateKey, nil
}
//签名
func Signbymyself (jsondata []byte,url string,randstr string,timestr int64) string {
	s1 := strconv.FormatInt(timestr, 10)
	var str string
	if len(jsondata)>2 {
str = "POST"+"\n"+ url +"\n"+s1 +"\n"+randstr+"\n"+string(jsondata)+"\n"
	} else {
str = "GET"+"\n"+ url +"\n"+s1 +"\n"+randstr+"\n"
	}

    //打开key
	// fmt.Printf("======================================")
	// fmt.Printf(str)
	// fmt.Printf("======================================")
    file := "./keypem/apiclient_key.pem"
	privateKeyBytes, err := ioutil.ReadFile(file)
    //fmt.Println(string(privateKeyBytes))
	if err != nil {
		fmt.Printf("打开文件出错！！")
        //return
	}
	info := RsaSignWithSha256(str,string(privateKeyBytes))
	return string(info)

}
//签名
func Signbywx (jsondata []byte,url string,randstr string,timestr int64) string {
	s1 := strconv.FormatInt(timestr, 10)
	var str string
	if len(jsondata)>2 {
str = "POST"+"\n"+ url +"\n"+s1 +"\n"+randstr+"\n"+string(jsondata)+"\n"
	} else {
str = "GET"+"\n"+ url +"\n"+s1 +"\n"+randstr+"\n"
	}

    //打开key
	fmt.Printf("======================================")
	fmt.Printf(str)
	fmt.Printf("======================================")
    file := "./keypem/apiclient_key.pem"
	privateKeyBytes, err := ioutil.ReadFile(file)
    //fmt.Println(string(privateKeyBytes))
	if err != nil {
		fmt.Printf("打开文件出错！！")
        //return
	}
	info := RsaSignWithSha256(str,string(privateKeyBytes))
	return string(info)

}
func RsaSignWithSha256(neirong string, keyBytes string) []byte {
    h := sha256.New()
    h.Write([]byte(neirong))
    hashed := h.Sum(nil)
    block, _ := pem.Decode([]byte(keyBytes))
    if block == nil {
        panic(errors.New("private key error"))
    }
    pkcs8PrivateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
    if err != nil {
        fmt.Println("ParsePKCS8PrivateKey err", err)
        panic(err)
    }
privateKey := pkcs8PrivateKey.(*rsa.PrivateKey)
    signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
    if err != nil {
        fmt.Printf("Error from signing: %s\n", err)
        panic(err)
    }

    return signature
}
//256解密
func DecryptToString(apiv3Key, associatedData, nonce, ciphertext string) (certificate string, err error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher([]byte(apiv3Key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	certificateByte, err := gcm.Open(nil, []byte(nonce), decodedCiphertext, []byte(associatedData))
	if err != nil {
		return "", err
	}
	certificate = string(certificateByte)
	return certificate, nil
}

