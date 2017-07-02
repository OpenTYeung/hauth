package main

import (
	"github.com/astaxie/beego"
	//_ "github.com/hzwy23/asofdate/apps"
	"github.com/asofdate/hauth/core/service"
)

func main() {
	service.StartHauth()
	beego.Run()
}
