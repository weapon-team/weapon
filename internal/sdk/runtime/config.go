package runtime

import (
	"encoding/json"

	"github.com/weapon-team/weapon/pkg/configx"
)

var Setting Config // 配置(全局变量)

type Config struct {
	AppName    string     `yaml:"AppName"`
	Port       string     `yaml:"Port"`
	DataSource DataSource `yaml:"DataSource"`
	Redis      RedisCfg   `yaml:"Redis"`
}

// DataSource 数据库
type DataSource struct {
	Driver string `yaml:"Driver"`
	DNS    string `yaml:"DNS"`
}

// RedisCfg Redis配置
type RedisCfg struct {
	Addr      string `yaml:"Addr"`
	MaxActive int    `yaml:"MaxActive"`
	Username  string `yaml:"Username"`
	Password  string `yaml:"Password"`
}

// InitConfig 初始化配置
func InitConfig(filename string) error {
	return configx.LoadConfigFromFile(filename, &Setting, false)
}

func (c Config) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}
