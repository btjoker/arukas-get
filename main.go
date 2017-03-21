package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var Port = 8989.0
var ID, Token, Secret = apikey()
var file = `{
    "configs": [
        {
            "server": "%s",
            "server_port": %.0f,
            "password": "acjoker.tk",
            "method": "aes-256-cfb",
            "remarks": "arukas"
        }
    ],
    "strategy": null,
    "index": 0,
    "global": false,
    "enabled": true,
    "shareOverLan": false,
    "isDefault": false,
    "localPort": 1080,
    "pacUrl": null,
    "useOnlinePac": false
}`

func main() {
	init := time.Now()
	resolve(infoGet())
	run()
	fmt.Println(time.Since(init))
}

func infoGet() []byte {
	url := "https://app.arukas.io/api/containers"
	req, err := http.NewRequest("GET", url, nil)
	errCheck(err)
	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("User-Agent", "地外行星")
	req.SetBasicAuth(Token, Secret)
	resp, err := http.DefaultClient.Do(req)
	errCheck(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errCheck(err)
	return body
}

func errCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func resolve(data []byte) {
	var (
		dat     map[string]interface{}
		portMap []interface{}
		host    interface{}
		port    interface{}
	)
	if err := json.Unmarshal(data, &dat); err != nil {
		panic(err)
	}
	applist := dat["data"].([]interface{})
	for _, v := range applist {
		if v.(map[string]interface{})["id"] == ID {
			portMap = v.(map[string]interface{})["attributes"].(map[string]interface{})["port_mappings"].([]interface{})
		}
	}
	for _, v := range portMap {
		for _, i := range v.([]interface{}) {
			if i.(map[string]interface{})["container_port"] == Port {
				host = i.(map[string]interface{})["host"]
				port = i.(map[string]interface{})["service_port"]
			}
		}
	}
	centen := fmt.Sprintf(file, host, port)
	file, _ := os.Create("gui-config.json") // create创建文件时如果存在会清空文件，不会返回错误
	defer file.Close()
	file.WriteString(centen)
}

func apikey() (ID, Token, Secret string) {
	file, err := ioutil.ReadFile("./apikey.txt")
	errCheck(err)

	doc := strings.Split(string(file), "\r\n")
	ID = strings.TrimSpace(strings.Split(doc[0], ":")[1])
	Token = strings.TrimSpace(strings.Split(doc[1], ":")[1])
	Secret = strings.TrimSpace(strings.Split(doc[2], ":")[1])
	return ID, Token, Secret
}

func run() {
	err := exec.Command("Shadowsocks.exe").Start()
	if err != nil {
		log.Fatalln("找不到 Shadowsocks.exe 文件， 请将本程序移动到Shadowsocks.exe根目录！")
	}
	log.Println("状态：正常运行")
}
