package runtime

import (
	"encoding/json"
	"github.com/weapon-team/weapon/pkg/configx"
)

var Setting Config // 配置(全局变量)

type Config struct {
	AppName    string
	Port       string
	DataSource DataSource
}

// DataSource 数据库
type DataSource struct {
	Driver string
	DNS    string
}

// InitConfig 初始化配置
func InitConfig(filename string) error {
	return configx.LoadConfigFromFile(filename, &Setting, false)
}

func (c Config) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}
