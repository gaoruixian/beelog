package main

import (
	"beelog/controllers"
	"beelog/models"
	_ "beelog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)

func init() {
	models.RegisterDB()

}
func main() {
	//开启ORM调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)
	//注册beego路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.AutoRouter(&controllers.TopicController{})

	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	//作为静态文件
	beego.SetStaticPath("/attachment", "attachment")
	//作为单独一个控制器来处理
	beego.Router("/attachment/:all", &controllers.AttachController{})
	//启动beego
	beego.Run()
}
