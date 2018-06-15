package routers

import (
	"myfirstweb/controllers"
	"net/http"

	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.Router("/", &controllers.BaseController{})

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
	//普通带参url取值参见：AutoController.GetInfo()方法

	//注解路由 只需要 Include 相应地 controller，然后在 controller 的 method 方法上面写上 router 注释（// @router）就可以
	beego.Include(&controllers.SelfController{})

	//请求数据处理
	beego.Router("/user/regist", &controllers.MainController{}, "post:Regist")

	//上传文件
	beego.Router("/file/upload", &controllers.MainController{}, "post:UploadFile")

	//数据绑定
	beego.Router("/data/bind", &controllers.MainController{}, "get:DataBind")

	//session控制
	beego.Router("/session/set", &controllers.MainController{}, "get:SetSessionFunc")
	beego.Router("/session/get", &controllers.MainController{}, "get:GetSessionFunc")

	//过滤器，还可以通过正则路由进行过滤，如果匹配参数就执行
	beego.InsertFilter("/session/get", beego.BeforeRouter, WebFilter)

	//表单数据验证
	beego.Router("/student/save", &controllers.ValidController{}, "get:SaveStudent")
	beego.Router("/teacher/save", &controllers.ValidController{}, "get:SaveTeacher")

	//错误处理
	beego.ErrorController(&controllers.ErrorController{})
	//	beego.ErrorHandler("404", err_404) // 这两个都可以，下面的最后注册，会默认执行

	//日志处理
	beego.SetLogger("file", "{\"filename\":\"logs/out.log\"}")
	//这个默认情况就会同时输出到两个地方，一个 console，一个 file，如果只想输出到文件，就需要调用删除操作：
	// beego.BeeLogger.DelLogger("console")
	// 日志的级别如上所示的代码这样分为八个级别：
	/*LevelEmergency
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
	级别依次降低，默认全部打印，但是一般我们在部署环境，可以通过设置级别设置日志级别：*/
	beego.SetLevel(beego.LevelDebug)
	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	beego.SetLogFuncCall(true)

	//ORM 配置了自动映射
	beego.AutoRouter(&controllers.OrmController{})
}

var WebFilter = func(ctx *context.Context) {
	u := ctx.Request.RequestURI //获取请求的链接 url
	i := ctx.Request.RemoteAddr //获取请求ip地址
	log.Println(u, i)
	uid, ok := ctx.Input.Session("uid").(int)           //获取不到session中的值
	id, status := ctx.Input.CruSession.Get("uid").(int) //可以获取到
	log.Println("uid: ", uid, id, status)
	//	log.Fatalln("uid: ", uid) //调用之后会调用os.exit(1) 接口退出程序
	log.Println("ok: ", ok)
	if ok {
		log.Println(ok)
	} else {
		log.Println(uid)
		//ctx.Redirect(302, "/index")
	}
}

func err_404(rw http.ResponseWriter, req *http.Request) {
	//	data := make(map[string]interface{})
	data := "page not fond"
	rw.Write([]byte(data))
}
