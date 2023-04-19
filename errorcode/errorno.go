package errorcode

// 定义错误码
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}
