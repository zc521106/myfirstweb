package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["myfirstweb/controllers:SelfController"] = append(beego.GlobalControllerRouter["myfirstweb/controllers:SelfController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getSelf`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myfirstweb/controllers:SelfController"] = append(beego.GlobalControllerRouter["myfirstweb/controllers:SelfController"],
		beego.ControllerComments{
			Method: "GetStr",
			Router: `/getSelfStr`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
