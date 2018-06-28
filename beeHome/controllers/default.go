package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type UserController struct{
	beego.Controller
}
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"//在controlled里面将数据给了map给存储起来

	c.TplName = "index.tpl"
}
//这个时候我么使用来自user的方法来使用
func (u *UserController)  Get(){

	u.Data["name"] = "liliangtbin"
	u.Data["grade"] = "95"
	//u.Ctx.WriteString("hello") 直接输出的方式
	//使用模板的方式
	u.TplName = "user.tpl"
}