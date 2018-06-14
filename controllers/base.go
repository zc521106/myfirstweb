package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。
func (this *BaseController) Prepare() {
	fmt.Println(" === prepare function === ")
}

func (this *BaseController) Get() {
	//转发到 /index
	this.Ctx.Redirect(302, "/index")
}

func (c *AutoController) error() {
	c.Ctx.WriteString("ERROR PAGE ")
}
