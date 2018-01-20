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
	if name == "" {
		beego.Error("enroll error: don't have any user name")
		c.CustomAbort(500, NewErrMsg("enroll error: don't have any user name").String())
	}
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

func (c UserController) AddUserFact() {
	userID := c.GetString("userId")
	if userID == "" {
		beego.Error("add fact error: unknown user")
		c.CustomAbort(500, NewErrMsg("unknown user").String())
	}
	name := c.GetString("name")
	value := c.GetString("value")
	if name == "" || value == "" {
		beego.Error("add fact error: don't have fact's name [" + name + "] or value [" + value + "]")
		c.CustomAbort(500, NewErrMsg("don't have fact's name ["+name+"] or value ["+value+"]").String())
	}

	err := game.AddFact(userID, name, value)
	if err != nil {
		beego.Error("add fact error:", err)
		c.CustomAbort(500, err.Error())
	}

	if s := c.GetString("callback"); s == "" {
		c.Data["json"] = NewSuccessMsg("success")
		c.ServeJSON()
	} else {
		c.Data["jsonp"] = NewSuccessMsg("success")
		c.ServeJSONP()
	}
}
