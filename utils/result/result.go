package result

const (
	SUCCESS      = 200
	ERROR        = 500
	UsernameUsed = 10001
)

var codeMsg = map[int]string{
	SUCCESS:      "成功",
	ERROR:        "未知错误",
	UsernameUsed: "昵称已存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
