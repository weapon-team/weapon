package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// SysStorage 系统存储表
type SysStorage struct {
	Id             int64  `xorm:"pk autoincr comment('ID') BIGINT"`
	Name           string `xorm:"not null comment('名称') VARCHAR(100)"`
	Code           string `xorm:"not null comment('编码') unique VARCHAR(30)"`
	Type           uint   `xorm:"not null default 1 comment('类型（1：兼容S3协议存储；2：本地存储）') UNSIGNED TINYINT"`
	AccessKey      string `xorm:"comment('Access Key（访问密钥）') VARCHAR(255)"`
	SecretKey      string `xorm:"comment('Secret Key（私有密钥）') VARCHAR(255)"`
	Endpoint       string `xorm:"comment('Endpoint（终端节点）') VARCHAR(255)"`
	BucketName     string `xorm:"comment('桶名称') VARCHAR(255)"`
	Domain         string `xorm:"not null default '' comment('自定义域名') VARCHAR(255)"`
	Description    string `xorm:"comment('描述') VARCHAR(200)"`
	IsDefault      int    `xorm:"not null default b'0' comment('是否为默认存储') BIT(1)"`
	Sort           int    `xorm:"not null default 999 comment('排序') INT"`
	Status         uint   `xorm:"not null default 1 comment('状态（1：启用；2：禁用）') UNSIGNED TINYINT"`
	CreateUser     int64  `xorm:"not null comment('创建人') index BIGINT"`
	UpdateUser     int64  `xorm:"comment('修改人') index BIGINT"`
	base.BaseModel `xorm:"extends"`
}
