package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/wangxianzhuo/mem-game/server/game"
)

type GameController struct {
	beego.Controller
}

func (c GameController) GetMap() {
	picPath := beego.AppConfig.String("PicPath")
	if picPath == "" {
		beego.Error("can't find any pics' uri")
		c.CustomAbort(500, NewErrMsg("can't find any pics' uri").String())
	}
	level, err := c.GetInt("level")
	if err != nil {
		var err1 error
		level, err1 = beego.AppConfig.Int("DefaultLevel")
		if err1 != nil {
			beego.Error("get map error: ", err1)
			c.CustomAbort(500, NewErrMsg(err1.Error()).String())
		}
		beego.Info("can't get level, use default level:", level)
	}
	len, err := c.GetInt("len")
	if err != nil {
		var err1 error
		len, err1 = beego.AppConfig.Int(fmt.Sprintf("Level%vLen", level))
		if err1 != nil {
			beego.Error("get map error: ", err1)
			c.CustomAbort(500, NewErrMsg(err1.Error()).String())
		}
		beego.Info("can't get len, use default len:", len)
	}
	wid, err := c.GetInt("wid")
	if err != nil {
		var err1 error
		wid, err1 = beego.AppConfig.Int(fmt.Sprintf("Level%vWid", level))
		if err1 != nil {
			beego.Error("get map error: ", err1)
			c.CustomAbort(500, NewErrMsg(err1.Error()).String())
		}
		beego.Info("can't get wid, use default wid:", wid)
	}

	gameMap, err := game.NewGameMap(len, wid, picPath)
	if err != nil {
		beego.Error("get map error: ", err)
		c.CustomAbort(500, NewErrMsg(err.Error()).String())
	}

	if s := c.GetString("callback"); s == "" {
		c.Data["json"] = gameMap
		c.ServeJSON()
	} else {
		c.Data["jsonp"] = gameMap
		c.ServeJSONP()
	}
}
