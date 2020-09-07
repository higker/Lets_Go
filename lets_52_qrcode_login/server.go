package main

import (
	"fmt"
	"github.com/go-redis/redis"
	qrcode "github.com/skip2/go-qrcode"
	"golang.org/x/net/websocket"
	"html/template"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	address       = "127.0.0.1:6379"
	password      = "admin"
	port          = ":8080"
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	qrcodeContent = "http://%s%s/verify?token=%s"
	redisProfix   = "qrcode_login_token:"
)

var (
	rdb *redis.Client
)

func init() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	if rdb.Ping().Err() != nil {
		return
	}
}

// 手机端请求处理 http://192.168.88.7:8080/verify?token=唯一的登录请求url&info=客户端存储的用户加密的数据
// 二维码登录服务器
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/verify", verify)
	http.HandleFunc("/qrcode", gqrcode)
	http.Handle("/ws/check_login", websocket.Handler(checkLogin))
	http.ListenAndServe(port, nil)
}

func gqrcode(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "image/png")
	uuid := request.URL.Query().Get("token")
	// 这里把token放到redis里面 然后下次这里面有值说明用户登录成功  10分钟自动超时
	if rdb.Set(redisProfix+uuid, "noInfo", 10*time.Second*60).Err() != nil {
		fmt.Println("REDIS SET ERROR")
	}
	png, err := qrcode.Encode(fmt.Sprintf(qrcodeContent, getIP(), port, uuid), qrcode.Medium, 256)
	if err != nil {
		return
	}
	writer.Write(png)
}

func checkLogin(ws *websocket.Conn) {

	var err error

	for {
		time.Sleep(2 * time.Second)
		token := ws.Request().URL.Query().Get("token")
		data, _ := rdb.Get(redisProfix + token).Result()

		//这里是发送消息
		if err = websocket.Message.Send(ws, func() string {
			if data == "admin" {
				return "succeed"
			}
			return ""
		}()); err != nil {

			fmt.Println("send failed:", err)

			break

		}

	}

}

func verify(writer http.ResponseWriter, request *http.Request) {
	token := request.URL.Query().Get("token")
	info := request.URL.Query().Get("info")
	fmt.Println(token)
	data, err := rdb.Get(redisProfix + token).Result()
	if err != nil || data != "noInfo" {
		writer.Write([]byte("你的登录二维码已经过期！！！或登录信息已失效！！"))
		return
	}
	// 模拟校验 这里的info可以是客户端登录保存的用户token也可以是用户信息加密的字符串 传输到服务器端进行校验
	if info == "admin" {
		val := rdb.TTL(redisProfix + token).Val()
		rdb.Set(redisProfix+token, info, val)
		writer.Write([]byte("登录成功正在通知PC端页面跳转！！"))
		return
	}
	writer.Write([]byte("你客户端信息存在异常！！请重新登录客户端再扫描二维码！！"))
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	files, err := template.ParseFiles("./index.html")
	if err != nil {
		writer.Write([]byte(err.Error()))
	}
	files.Execute(writer, randUUID())
}

func randUUID() string {
	bit := make([]byte, 32)
	for b := range bit {
		bit[b] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bit)
}

func getIP() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
