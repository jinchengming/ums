package result

const (
	SUCCESS = 200
	ERROR   = 500
)

var codeMsg = map[int]string{
	SUCCESS:               "成功",
	ERROR:                 "未知错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
