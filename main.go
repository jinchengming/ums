package main

import (
	"ums/model"
	"ums/routers"
)

func main() {
	model.DB.AutoMigrate(&model.User{})
	routers.InitRouter()
}
