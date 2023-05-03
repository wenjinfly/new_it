package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// 自定义时间
type LocalTime time.Time

// UnmarshalJSON 1. 为 LocalTime 重写 UnmarshalJSON 方法，在此方法中实现自定义格式的转换；
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime(t1)
	return err
}

// MarshalJSON 2. 为 LocalTime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value 3. 为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (t LocalTime) Value() (driver.Value, error) {
	// LocalTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime, nil
	//return tTime.Format("2006-01-02 15:04:05"), nil
}

// Scan 4. 为 Time 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (t *LocalTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = LocalTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *LocalTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
