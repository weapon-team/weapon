package types

import "time"

// Timestamp 秒级时间戳
type Timestamp int64

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}

// NowTimestamp 当前时间戳
func NowTimestamp() Timestamp {
	return Timestamp(time.Now().Unix())
}
