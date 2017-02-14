package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

var file = `
{
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

const (
	// ID appId
	ID = ""
	// Token Token 1
	Token = ""
	// Secret Secret 1
	Secret = ""
)

func main() {
	resolve(info(), ID)
	err := exec.Command("Shadowsocks.exe").Start()
	if err != nil {
		fmt.Println("找不到 Shadowsocks.exe 文件， 请将本程序移动到Shadowsocks.exe根目录！")
		return
	}
	fmt.Println("状态：正常运行")
}

func info() []byte {
	url := "https://app.arukas.io/api/containers"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("User-Agent", "地外行星")
	req.SetBasicAuth(Token, Secret)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func resolve(data []byte, ID string) {
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
			if i.(map[string]interface{})["container_port"] == 8989.0 { // json返回的数字大多是float64
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