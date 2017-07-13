package main

import (
	//"github.com/asofdate/apps/mas/ca"
	//"github.com/asofdate/apps/mas/common"
	//"github.com/asofdate/apps/mas/ftp"
	"github.com/asofdate/hauth/core/service"
)

func main() {
	//service.AppRegister("ca", ca.Register)
	//service.AppRegister("ftp", ftp.Register)
	//service.AppRegister("common", common.Register)
	service.Bootstrap()
}
