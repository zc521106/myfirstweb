package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type OrmController struct {
	BaseController
}

type User struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" orm:"size(64)" json:"name"`
}

type KeyValue struct {
	Key   string
	Value string
}

var o orm.Ormer

func init() {
	orm.RegisterDataBase("default", "mysql", "root:MysqlPsw1!@tcp(39.106.15.201:3306)/go_test?charset=utf8", 30)
	orm.RegisterModel(new(User))
	//	orm.RegisterModel(new(KeyValue)) //手动建表
	orm.RunSyncdb("default", false, true)
	o = orm.NewOrm()
	//开启orm调试，会输出sql语句
	orm.Debug = true
}

func (c *OrmController) SaveUser() {
	res := make(map[string]interface{})
	u := User{}
	if err := c.ParseForm(&u); err == nil {
		beego.Debug("succeed")
		res["errcode"] = -1
		//		o := orm.NewOrm()
		//insert
		id, errmsg := o.Insert(&u)
		// 同时插入多个
		//	users := []User{{Name: "slene"}, {Name: "astaxie"}, {Name: "unknown"}, ...}
		//	o.InsertMulti(100, uses)
		if errmsg == nil {
			res["errcode"] = 0
		}
		res["msg"] = id
		res["id"] = errmsg
	} else {
		beego.Debug("error")
		res["errcode"] = 1
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *OrmController) UpdateUser() {
	res := make(map[string]interface{})
	u := User{}
	if err := c.ParseForm(&u); err == nil {
		beego.Debug("succeed")
		res["errcode"] = -1
		//		o := orm.NewOrm()
		//update
		num, errmsg := o.Update(&u)
		// 只更新 Name
		// o.Update(&user, "Name")
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)
		if errmsg == nil {
			res["errcode"] = 0
		}
		res["msg"] = num
		res["num"] = errmsg
	} else {
		beego.Debug("error")
		res["errcode"] = 1
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *OrmController) QueryUser() {
	res := make(map[string]interface{})
	id := c.GetString("id")
	idi, _ := strconv.Atoi(id)
	u := User{Id: idi}
	if err := c.ParseForm(&u); err == nil {
		beego.Debug("succeed")
		res["errcode"] = -1
		//		o := orm.NewOrm()
		//read
		errmsg := o.Read(&u)
		// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
		// user := User{Name: "slene"}
		// err = o.Read(&user, "Name")

		// ReadOrCreate 尝试从数据库读取，不存在的话就创建一个，默认必须传入一个参数作为条件字段，同时也支持多个参数多个条件字段
		// user := User{Name: "slene"}
		// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
		// if created, id, err := o.ReadOrCreate(&user, "Name"); err == nil {
		//    if created {
		//        fmt.Println("New Insert an object. Id:", id)
		//   } else {
		//        fmt.Println("Get an object. Id:", id)
		//    }
		// }
		if errmsg == nil {
			res["errcode"] = 0
		}
		res["msg"] = id
		res["data"] = u
	} else {
		beego.Debug("error")
		res["errcode"] = 1
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *OrmController) DeleteUser() {
	res := make(map[string]interface{})
	id := c.GetString("id")
	idi, _ := strconv.Atoi(id)
	u := User{Id: idi}
	if err := c.ParseForm(&u); err == nil {
		beego.Debug("succeed")
		res["errcode"] = -1
		//		o := orm.NewOrm()
		//delete
		num, errmsg := o.Delete(&u)
		if errmsg == nil {
			res["errcode"] = 0
		}
		res["msg"] = errmsg
		res["num"] = num
	} else {
		beego.Debug("error")
		res["errcode"] = 1
	}
	c.Data["json"] = res
	c.ServeJSON()
}

//原生SQL 操作数据库
func (c *OrmController) QueryUsers() {
	var users []User
	ids := []int{1, 2, 3, 6}
	num, err := o.Raw("select id, name from user where id in (?, ?, ?, ?)", ids).QueryRows(&users)
	if err == nil {
		beego.Debug(num)
		c.Data["json"] = users
	} else {
		beego.Error(err)
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

func (c *OrmController) QueryUserById() {
	id := c.GetString("id")
	//	var user User
	user := User{}
	err := o.Raw("select id, name from user where id = ? ", id).QueryRow(&user)
	if err != nil {
		c.Ctx.WriteString("error")
	} else {
		c.Data["json"] = user
		c.ServeJSON()
	}
}

func (c *OrmController) UpdateUserNameById() {
	id := c.GetString("id")
	name := c.GetString("name")
	res, err := o.Raw("update user set name = ? where id = ?", name, id).Exec()
	/*
		SetArgs
		改变 Raw(sql, args…) 中的 args 参数，返回一个新的 RawSeter
		用于单条 sql 语句，重复利用，替换参数然后执行。
		res, err := r.SetArgs("arg1", "arg2").Exec()
		res, err := r.SetArgs("arg1", "arg2").Exec()
	*/
	if err == nil {
		num, errmsg := res.RowsAffected()
		if errmsg == nil {
			c.Ctx.WriteString(strconv.FormatInt(num, 10))
		} else {
			c.Ctx.WriteString("error 0")
		}
	} else {
		c.Ctx.WriteString("error 1")
	}
}

/*
	[
	  {
	    "id": "1",
	    "name": "zhaochao"
	  },
	  {
	    "id": "3",
	    "name": "王钢蛋"
	  },
	  {
	    "id": "4",
	    "name": "往上举"
	  }
	]
*/
func (c *OrmController) TestValues() {
	var maps []orm.Params
	id := c.GetString("id")
	num, err := o.Raw("select id, name from user where id >= ?", id).Values(&maps)
	if err == nil && num > 0 {
		beego.Debug(num)
		c.Data["json"] = maps
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

/*
	[
	  [
	    "1",
	    "zhaochao"
	  ],
	  [
	    "3",
	    "王钢蛋"
	  ],
	  [
	    "4",
	    "往上举"
	  ],
	  [
	    "5",
	    "往上举··`"
	  ]
	]
*/
func (c *OrmController) TestValuesList() {
	var lists []orm.ParamsList
	id := c.GetString("id")
	num, err := o.Raw("select id, name from user where id >= ?", id).ValuesList(&lists)
	if err == nil && num > 0 {
		beego.Debug(num)
		c.Data["json"] = lists
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

/*
	[
	  "1",
	  "zhaochao",
	  "3",
	  "王钢蛋",
	  "4",
	  "往上举"
	]
*/
func (c *OrmController) TestValuesFlat() {
	var list orm.ParamsList
	id := c.GetString("id")
	num, err := o.Raw("select id, name from user where id >= ?", id).ValuesFlat(&list)
	if err == nil && num > 0 {
		beego.Debug(num)
		c.Data["json"] = list
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

/*
	{
	  "zhaochao": "1",
	  "往上举": "4",
	  "往上举··`": "5",
	  "王钢蛋": "3",
	  "赵铁锤": "6",
	  "赵铁锤221": "7",
	  "铁锤妹妹3": "8"
	}
*/
//将 name值映射为map的key，id为value
func (c *OrmController) TestRowsToMap() {
	res := make(orm.Params)
	id := c.GetString("id")
	num, err := o.Raw("select id, name from user where id >= ?", id).RowsToMap(&res, "name", "id")
	if err == nil && num > 0 {
		beego.Debug(num)
		c.Data["json"] = res
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

// 注意：查询到的key和value字段值分别为对象的属性及相应的属性值
func (c *OrmController) TestRowsToStruct() {
	//	id := c.GetString("id")
	res := new(ResultBody)
	num, err := o.Raw("select key, value from key_value").RowsToStruct(res, "key", "value")
	beego.Debug(res.LogLevel)
	if err == nil && num > 0 {
		beego.Debug(num)
		beego.Debug(res)
		c.Data["json"] = res
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

/*
	id	  key		value
	1	loglevel	debug
	2	msg			nothing
*/
type ResultBody struct {
	LogLevel string
	Msg      string
}
