package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetInfo() {
	c.Ctx.WriteString("main info")
}

func (c *MainController) GetCon() {
	s := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString(s)
}

func (c *MainController) GetInt() {
	s := c.Ctx.Input.Param(":num")
	c.Ctx.WriteString("数字：" + s)
}

func (c *MainController) File() {
	a := string(c.Ctx.Input.ParamsLen())
	b := c.Ctx.Input.Param("ext")
	c.Ctx.WriteString("文件路径：" + a + " 文件类型：" + b)
}
