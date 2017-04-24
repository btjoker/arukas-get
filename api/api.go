package api

import (
	"encoding/json"
	"errors"
	"log"
)

var ErrResult = errors.New("结构体为空")

// Containers 容器
type Containers struct {
	ID            string       `json:"id"`
	Type          string       `json:"type"`
	Attributes    Attribute    `json:"attributes"`
	Relationships Relationship `json:"relationships"`
}

// Attribute app配置信息
type Attribute struct {
	AppID        string       `json:"app_id"`
	ImageName    string       `json:"image_name"`
	Cmd          string       `json:"cmd"`
	IsRunning    bool         `json:"is_running"`
	Instances    int          `json:"instances"`
	Mem          int          `json:"mem"`
	Envs         Envs         `json:"envs"`
	Ports        Ports        `json:"ports"`
	PortMappings PortMappings `json:"port_mappings"`
	CreatedAt    string       `json:"created_at"`
	UpdatedAd    string       `json:"updated_ad"`
	StatusText   string       `json:"status_text"`
	ArukasDomain string       `json:"arukas_domain"`
	EndPoint     string       `json:"end_point"`
}

// Env 环境变量
type Env struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Envs 环境变量数组
type Envs []Env

// Port 端口
type Port struct {
	Number   int    `json:"number"`
	Protocol string `json:"protocol"`
}

// Ports 端口数组
type Ports []Port

// PortMapping 端口地图
type PortMapping struct {
	ContainerPort int    `json:"container_port"`
	ServicePort   int    `json:"service_port"`
	Host          string `json:"host"`
}

// TaskPorts 端口地图数组
type TaskPorts []*PortMapping

// PortMappings 端口地图包裹
type PortMappings []TaskPorts

// Relationship 关联
type Relationship struct {
	App Apps `json:"app"`
}

// Apps .
type Apps struct {
	Data Datas `json:"data"`
}

// Datas app数据
type Datas struct {
	ID   string `json:"ID"`
	Type string `json:"type"`
}

// Application 数据
type Application struct {
	Data []*Containers `json:"data"`
}

// Config ss 配置文件
type Config struct {
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	Password   string `json:"password"`
	Method     string `json:"method"`
	Remarks    string `json:"remarks"`
}

// GuiConfig ss客户端配置文件
type GuiConfig struct {
	Configs      []Config `json:"configs"`
	Strategy     *string  `json:"strategy"`
	Index        int      `json:"index"`
	Global       bool     `json:"global"`
	Enabled      bool     `json:"enabled"`
	ShareOverLan bool     `json:"shareOverLan"`
	IsDefault    bool     `json:"isDefault"`
	LocalPort    int      `json:"localPort"`
	PacURL       *string  `json:"pacUrl"`
	UseOnlinePac bool     `json:"useOnlinePac"`
}

// Add 添加一个ss配置
func (G *GuiConfig) Add(server, password string, port int) {
	G.Configs = []Config{
		Config{
			Server:     server,
			ServerPort: port,
			Password:   password,
			Method:     "aes-256-cfb",
			Remarks:    "Arukas",
		},
	}
	G.Enabled = true
	G.LocalPort = 1080
}

// Result 将结构体转为一个 []byte
func (G *GuiConfig) Result() ([]byte, error) {
	if G.Configs == nil {
		return []byte(nil), ErrResult
	}
	data, err := json.Marshal(G)
	if err != nil {
		log.Fatalln(err)
	}
	return data, nil
}
