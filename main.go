package main

import (
	_ "github.com/jicg/liteblog/routers"
	_ "github.com/jicg/liteblog/models"
	"github.com/astaxie/beego"
	"strings"
	"encoding/gob"
	"github.com/jicg/liteblog/models"
	"os"
	"github.com/astaxie/beego/logs"
)

func main() {
	initLog()
	initSession()
	initTemplate()
	beego.Run()
}
func initLog() {
	if err := os.MkdirAll("data/logs", 0777); err != nil {
		beego.Error(err)
		return
	}
	logs.SetLogger("file", `{"filename":"data/logs/lyblog.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.Async(1e3)
}

func initSession() {
	gob.Register(&models.User{})
	//https://beego.me/docs/mvc/controller/session.md
	beego.SetStaticPath("assert", "assert")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}
func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})

}
