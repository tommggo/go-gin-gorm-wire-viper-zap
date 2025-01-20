package timeutil

import "time"

const (
	// StandardMilliFormat 标准时间格式（精确到毫秒）
	StandardMilliFormat = "2006-01-02 15:04:05.000"

	// StandardFormat 标准时间格式
	StandardFormat = "2006-01-02 15:04:05"

	// DateFormat 日期格式
	DateFormat = "2006-01-02"

	// TimeFormat 时间格式
	TimeFormat = "15:04:05"

	// MillisecondsPerSecond 每秒的毫秒数
	MillisecondsPerSecond = 1000
)

// FromUnixMilli 将毫秒级时间戳转换为 time.Time
func FromUnixMilli(msTimestamp uint64) time.Time {
	return time.UnixMilli(int64(msTimestamp))
}

// ToUnixMilli 将 time.Time 转换为毫秒级时间戳
func ToUnixMilli(t time.Time) uint64 {
	return uint64(t.UnixMilli())
}

// FormatMilliTime 格式化时间，精确到毫秒
func FormatMilliTime(t time.Time) string {
	return t.Format(StandardMilliFormat)
}

// ParseMilliTime 解析精确到毫秒的时间字符串
func ParseMilliTime(s string) (time.Time, error) {
	return time.Parse(StandardMilliFormat, s)
}

// Format 使用指定格式格式化时间
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// Parse 使用指定格式解析时间字符串
func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

// Now 获取当前时间
func Now() time.Time {
	return time.Now()
}

// NowMilli 获取当前毫秒时间戳
func NowMilli() uint64 {
	return ToUnixMilli(Now())
}
