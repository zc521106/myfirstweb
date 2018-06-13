package controllers

type AutoController struct {
	MainController
}

func (c *AutoController) Get() {
	c.Ctx.WriteString("welcome to auto")
}

func (c *AutoController) GetStr() {
	c.Ctx.WriteString("auto getstr")
}
