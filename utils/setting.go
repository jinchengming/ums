package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	// 加载ini配置文件
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误：", err)
	}
	LoadDatabase(file)
}

// 加载数据库配置
func LoadDatabase(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("chengming123")
	DbName = file.Section("database").Key("DbName").MustString("go-blog")
}