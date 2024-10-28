package model

import "github.com/weapon-team/weapon/internal/sdk/base"

// SysFile 文件
type SysFile struct {
	Id             int64  `xorm:"pk autoincr comment('ID') BIGINT"`
	Name           string `xorm:"not null comment('名称') VARCHAR(255)"`
	Size           int64  `xorm:"not null comment('大小（字节）') BIGINT"`
	Url            string `xorm:"not null comment('URL') index VARCHAR(512)"`
	Extension      string `xorm:"comment('扩展名') VARCHAR(100)"`
	Type           uint   `xorm:"not null default 1 comment('类型（1：其他；2：图片；3：文档；4：视频；5：音频）') index UNSIGNED TINYINT"`
	StorageId      int64  `xorm:"not null comment('存储库ID') BIGINT"`
	CreateUser     int64  `xorm:"not null comment('创建人') index BIGINT"`
	UpdateUser     int64  `xorm:"not null comment('修改人') index BIGINT"`
	base.TimeModel `xorm:"extends"`
}
