package main

import (
	_ "golang_test/beeHome/routers"
	"github.com/astaxie/beego"
)

func main() {
	//StaticDir["/static"] = "static"
	beego.SetStaticPath("/down1", "download1")//静态资源加载位置
	beego.Run()
}

