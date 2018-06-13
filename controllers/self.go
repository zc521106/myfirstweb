package controllers

type SelfController struct {
	MainController
}

//func (c *CMSController) URLMapping() {
//	c.Mapping("GetSelf", c.Get)
//	c.Mapping("GetSelfStr", c.GetStr)
//}

// @router /getSelf [get]
func (c *SelfController) Get() {
	c.Ctx.WriteString("welcome to self")
}

// @router /getSelfStr
func (c *SelfController) GetStr() {
	c.Ctx.WriteString("auto getstr")
}
