package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/btjoker/arukas-get/api"
)

const (
	// Token Token
	Token = ""
	// Secret Secret
	Secret = ""
	// Port ss设置的端口
	Port = 8989
	// Password ss的密码,
	Password = ``
)

func main() {
	init := time.Now()
	if !resolve() {
		fmt.Println("无法写入, 生成失败!")
		return
	}
	fmt.Println("生成成功!")
	exec.Command("Shadowsocks.exe").Start()
	fmt.Println(time.Since(init))
}

func resolve() bool {
	var (
		JSON    api.Application
		configs api.GuiConfig
	)
	req, _ := http.NewRequest("GET", "https://app.arukas.io/api/containers", nil)
	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("User-Agent", "地外行星")
	// 基本认证
	req.SetBasicAuth(Token, Secret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &JSON); err != nil {
		log.Fatalln(err)
	}
	for _, v1 := range JSON.Data[0].Attributes.PortMappings {
		for _, v2 := range v1 {
			if v2.ContainerPort == Port {
				configs.Add(v2.Host, Password, v2.ServicePort)
			}
		}
	}
	file, err := os.Create("gui-config.json")
	if err != nil {
		log.Fatalln(err)
	}
	if data, err := configs.Result(); err == nil {
		file.Write(data)
		return true
	}
	defer file.Close()
	return false
}
