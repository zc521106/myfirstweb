package controllers

type ErrorController struct {
	BaseController
}

func (e *ErrorController) Error404() {
	//	e.Ctx.WriteString("error 404 requestUrl not found")
	e.Data["errmsg"] = "request url not found"
	e.TplName = "404.tpl"
}
