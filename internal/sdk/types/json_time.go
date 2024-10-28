package types

import "time"

// JsonTime 统一时间类型
// desc: 返回给前端时自动转为 2006-01-02 15:04:05 格式
type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(time.DateTime) + `"`), nil
}
