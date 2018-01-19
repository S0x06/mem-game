package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wangxianzhuo/mem-game/server/game"
)

type UserController struct {
	beego.Controller
}

// TODO:
func (c UserController) Enroll() {
	name := c.GetString("name")
	beego.Debug("user name: ", name)
	result, err := game.Enroll(name)
	if err != nil {
		beego.Error(err)
		c.CustomAbort(500, err.Error())
	}

	if s := c.GetString("callback"); s == "" {
		c.Data["json"] = result
		c.ServeJSON()
	} else {
		c.Data["jsonp"] = result
		c.ServeJSONP()
	}
}
