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
	Jwt        JwtCfg     `yaml:"Jwt"`
	Log        LogCfg     `yaml:"Log"`
}

// LogCfg 日志配置
type LogCfg struct {
	Path     string // 日志文件目录, 为空时打印到控制台
	Level    string // 日志级别debug、info、warn、error
	ShowSQL  bool   // 是否打印sql语句
	Interval int    // 日志切割时间间隔, 单位:h
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

type JwtCfg struct {
	Secret string `yaml:"Secret"`
	Expire int    `yaml:"Expire"` // 过期时间
}

// InitConfig 初始化配置
func InitConfig(filename string) error {
	return configx.LoadConfigFromFile(filename, &Setting, false)
}

func (c Config) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}
