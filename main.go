package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/skip2/go-qrcode"
)

type config struct {
	ListenPort int
}

//请求二维码服务的处理方法
func server(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		bytes, err := getFileBytes("index.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		w.Write(bytes)
	} else {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		url := r.Form.Get("url")
		if len(url) == 0 {
			fmt.Fprintf(w, "未传参数url，无法生成二维码!")
			return
		}
		images, err := qrcode.Encode(url, qrcode.Medium, 256)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(images)
	}

}

func getFileBytes(fileName string) ([]byte, error) {
	bytes := make([]byte, 1024*1024)
	file, err := os.Open(fileName)
	if err != nil {
		return bytes[0:0], err
	}
	defer file.Close()
	n, err := file.Read(bytes)
	bytes = bytes[0:n]
	return bytes, err
}

func getConfig() (config, error) {
	var c config
	bytes, err := getFileBytes("config.json")
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(bytes, &c)
	return c, err

}
func main() {
	c, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}
	http.HandleFunc("/", server)
	log.Printf("服务已经启动，当前监听端口:%d", c.ListenPort)
	http.ListenAndServe(fmt.Sprintf(":%d", c.ListenPort), nil)

}
