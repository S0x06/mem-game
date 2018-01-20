package main

import (
	"github.com/astaxie/beego"
	"github.com/wangxianzhuo/mem-game/server/controllers"
)

func main() {
	beego.Router("/game/map", &controllers.GameController{}, "get:GetMap")
	beego.Router("/user/enroll", &controllers.UserController{}, "get:Enroll")
	beego.Router("/user/fact", &controllers.UserController{}, "get,post:AddUserFact")
	beego.Run()
}
