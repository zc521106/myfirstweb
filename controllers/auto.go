package controllers

import (
	"fmt"
)

type AutoController struct {
	BaseController
}

//先查找auto中的prepare()方法，有就执行，没有就寻找basecontroller中的prepare()方法并执行
func (c *AutoController) Prepare() {
	fmt.Println(" === prepare in auto === ")
	//提前终止运行，终止后 Finish()方法不会被调用
	if stop := c.GetString("stop"); stop == "true" {
		fmt.Println("stop")
		c.Data["json"] = map[string]interface{}{"reason": "stop=true", "msg": "用户终止操作！"}
		c.ServeJSON()
		c.StopRun()
	}
}

func (c *AutoController) Finish() {
	fmt.Println("执行finish清理工作！")
}

func (c *AutoController) Get() {
	c.Ctx.WriteString("welcome to auto")
}

func (c *AutoController) GetStr() {
	c.Ctx.WriteString("auto getstr")
}

//普通带参url取值
func (c *AutoController) GetInfo() {
	param1 := c.GetString("param1")
	param2 := c.GetString("param2")
	fmt.Println(param1, param2)
	c.Ctx.WriteString("param1: " + param1 + " param2: " + param2)
}

func (c *AutoController) Query() {
	c.Data["json"] = map[string]string{"reason": "stop=true", "msg": "用户终止操作！"}
	c.ServeJSON()
}
