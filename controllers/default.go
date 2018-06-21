package controllers

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/httplib"
)

type MainController struct {
	BaseController
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

//参数绑定到对象中去
func (c *MainController) Regist() {
	u := user{}
	if err := c.ParseForm(&u); err == nil {
		fmt.Println("succeed " + u.Username)
	} else {
		fmt.Println("error " + u.Username)
	}
	c.Data["json"] = u
	c.ServeJSON()
}

//对象属性首字母需要大写才能绑定成功
type user struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Phone    int    `form:"phonenum"`
}

/**
 *	Beego 提供了两个很方便的方法来处理文件上传：
 *	 	GetFile(key string) (multipart.File, *multipart.FileHeader, error)
 *			该方法主要用于用户读取表单中的文件名 the_file，然后返回相应的信息，用户根据这些变量来处理文件上传：过滤、保存文件等。
 *		SaveToFile(fromfile, tofile string) error
 *			该方法是在 GetFile 的基础上实现了快速保存的功能
 *			fromfile 是提交时候的 html 表单中的 name
 */
func (c *MainController) UploadFile() {
	file, header, erro := c.GetFile("uploadfilename")
	if erro != nil {
		log.Fatalln("get file error", erro)
	}
	defer file.Close()
	err := c.SaveToFile("uploadfilename", "static/upload/"+header.Filename)
	if err != nil {
		log.Fatalln("save file error", err)
	} else {
		c.Ctx.WriteString("ok")
	}
}

//数据绑定
func (c *MainController) DataBind() {
	var m, n string
	c.Ctx.Input.Bind(&m, "name")
	c.Ctx.Input.Bind(&n, "psw")
	c.Ctx.WriteString("m: " + m + "  n: " + n)
}

//session
func (c *MainController) SetSessionFunc() {
	k := c.GetString("key")
	v := c.GetString("value")
	c.SetSession(k, v)
	c.Ctx.WriteString("ok")
}

func (c *MainController) GetSessionFunc() {
	k := c.GetString("key")
	if v := c.GetSession(k); v != nil {
		c.Ctx.WriteString(v.(string))
	} else {
		c.Ctx.WriteString("nil")
	}
}

func (c *MainController) GetBlogInfo() {
	url := c.GetString("url")
	// 获取请求的方式：get、post等
	method := c.GetString("method")
	var req *httplib.BeegoHTTPRequest
	// 设置超时时间
	req.SetTimeout(10, 10)
	// 设置参数：参数可以直接在链接地址中拼接也可以通过这种方式设置
	//req.Param("key", "value")
	//req.Param("key", "value")
	if method == "get" {
		req = httplib.Get(url)
	} else if method == "post" {
		req = httplib.Post(url)
	} else if method == "put" {
		req = httplib.Put(url)
	} else if method == "delete" {
		req = httplib.Delete(url)
	} else {
		c.Ctx.WriteString("this method is not support!")
		return
	}
	res, err := req.String()
	if err == nil {
		c.Ctx.WriteString(res)
	} else {
		log.Println(res)
		log.Println(err)
		c.Ctx.WriteString("error")
	}
}
