package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// SysLog 系统日志
type SysLog struct {
	Id              int64  `xorm:"pk autoincr comment('ID') BIGINT"`
	TraceId         string `xorm:"comment('链路ID') VARCHAR(255)"`
	Description     string `xorm:"not null comment('日志描述') VARCHAR(255)"`
	Module          string `xorm:"not null comment('所属模块') index VARCHAR(50)"`
	RequestUrl      string `xorm:"not null comment('请求URL') VARCHAR(512)"`
	RequestMethod   string `xorm:"not null comment('请求方式') VARCHAR(10)"`
	RequestHeaders  string `xorm:"comment('请求头') TEXT(65535)"`
	RequestBody     string `xorm:"comment('请求体') TEXT(65535)"`
	StatusCode      int    `xorm:"not null comment('状态码') INT"`
	ResponseHeaders string `xorm:"comment('响应头') TEXT(65535)"`
	ResponseBody    string `xorm:"comment('响应体') MEDIUMTEXT(16777215)"`
	TimeTaken       int64  `xorm:"not null comment('耗时（ms）') BIGINT"`
	Ip              string `xorm:"comment('IP') index VARCHAR(100)"`
	Address         string `xorm:"comment('IP归属地') VARCHAR(255)"`
	Browser         string `xorm:"comment('浏览器') VARCHAR(100)"`
	Os              string `xorm:"comment('操作系统') VARCHAR(100)"`
	Status          uint   `xorm:"not null default 1 comment('状态（1：成功；2：失败）') UNSIGNED TINYINT"`
	ErrorMsg        string `xorm:"comment('错误信息') TEXT(65535)"`
	CreateUser      int64  `xorm:"comment('创建人') BIGINT"`
	base.BaseModel  `xorm:"extends"`
}
