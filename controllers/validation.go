package controllers

import (
	"log"
	"strings"

	"github.com/astaxie/beego/validation"
)

type ValidController struct {
	BaseController
}

type Student struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

// validation的直接使用
func (c *ValidController) SaveStudent() {
	s := Student{}
	if err := c.ParseForm(&s); err == nil {
		log.Println("in ", s.Name, s.Age)
		valid := validation.Validation{}
		valid.Required(s.Name, "name")
		valid.MaxSize(s.Name, 8, "name")
		valid.Range(s.Age, 0, 25, "age")
		//自定义错误校验和提示信息
		minAge := 18
		valid.Min(&s.Age, minAge, "age").Message("%d岁以上，少儿不宜", minAge)
		if valid.HasErrors() {
			for k, errMsg := range valid.Errors { // k是索引 可忽略
				log.Println(k, errMsg.Key, errMsg.Message)
			}
			c.Ctx.Output.Body([]byte("error 1"))
			return
		}
		c.Ctx.Output.Body([]byte("ok"))
		return
	} else {
		log.Println("error")
	}
	c.Ctx.Output.Body([]byte("error 2"))
	return
}

// 通过 StructTag 使用
// 	验证函数写在 "valid" tag 的标签里
// 	各个函数之间用分号 ";" 分隔，分号后面可以有空格
// 	参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
// 	正则函数(Match)的匹配模式用两斜杠 "/" 括起来
// 	各个函数的结果的 key 值为字段名.验证函数名
type Teacher struct {
	Id    int
	Name  string `form:"name" valid:"Required;Match(/^Chao.*/)"` // Name不可为空并且必须以Chao开头
	Age   int    `form:"age" valid:"Range(0, 28)"`               // 0 <= Age <= 28
	Email string `form:"email" valid:"Email";`                   // Email需要符合邮箱格式
	Phone string `form:"phone" valid:"Mobile; MaxSize(14)"`      // Phone必须为正确的手机号，且长度不能大于14
	IP    string `form:"ip" valid:"IP"`                          // IP必须为一个正确的 IPv4 地址
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (t *Teacher) Valid(v *validation.Validation) {
	errKey := "admin"
	if strings.Index(t.Name, errKey) > -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名字中不可包含"+errKey)
	}
}

func (c *ValidController) SaveTeacher() {
	t := Teacher{}
	valid := validation.Validation{}
	if err := c.ParseForm(&t); err == nil {
		log.Println(t.Name, t.Phone)
		res, er := valid.Valid(&t)
		if er != nil {
			log.Println(res, er)
			for _, msg := range valid.Errors {
				log.Println(msg.Key, msg.Message)
			}
			c.Ctx.WriteString("error0 ")
		}
		if !res { // res=false说明valid未通过
			errs := make(map[string]interface{})
			for index, msg := range valid.Errors {
				log.Println(index, msg.Key, msg.Message)
				errs[msg.Key] = msg.Message
			}
			c.Data["json"] = errs
			c.ServeJSON()
			// c.Ctx.WriteString("error 1 " + msg.Key + " " + msg.Message)
			return
		}
		c.Ctx.Output.Body([]byte(t.Name))
	} else {
		c.Ctx.WriteString("error 2")
	}
}
