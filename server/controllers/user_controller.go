package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// TODO:
func (c UserController) Enroll() {
	id := c.GetString("id")
	beego.Debug("user id: ", id)
	if s := c.GetString("callback"); s == "" {
		c.Data["json"] = id
		c.ServeJSON()
	} else {
		c.Data["jsonp"] = id
		c.ServeJSONP()
	}
}
