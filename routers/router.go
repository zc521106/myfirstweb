package routers

import (
	"myfirstweb/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.Router("/index", &controllers.MainController{})
	//	beego.Router("/auto1", &controllers.AutoController{})

	//自定义简单的映射
	beego.Get("/info", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello"))
	})

	//自动映射; /auto/functionName
	beego.AutoRouter(&controllers.AutoController{})

	//自定义方法及 RESTful 规则; 注意：controller中的方法名必须大写字母开头，否则编译时会报错无方法
	beego.Router("/info", &controllers.MainController{}, "*:GetInfo")
	beego.Router("/autoinfo", &controllers.AutoController{}, "*:GetStr")
	//正则路由    带参url必须有‘:’，否则收不到参数值 ‘?’表示可以匹配无参url ‘/getCon’，此url可匹配：/getCon、/getCon/12、/getCon/da 等
	beego.Router("/getCon/?:id", &controllers.MainController{}, "get:GetCon")
	//:int表示只匹配参数值为int类型的url，由于去掉了‘?’，所以不可匹配 ‘/getInt’
	beego.Router("/getInt/:num:int", &controllers.MainController{}, "get:GetInt")
	//download之后的链接是 文件路径  写了File方法之后会输出该方法中打印的内容到下载文件中去
	beego.Router("/download/*.*", &controllers.MainController{}, "get:File")

	//注解路由 只需要 Include 相应地 controller，然后在 controller 的 method 方法上面写上 router 注释（// @router）就可以
	beego.Include(&controllers.SelfController{})
}
