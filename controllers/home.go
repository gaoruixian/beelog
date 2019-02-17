package controllers

import (
	"beelog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	category := this.Input().Get("cate")
	topics, err := models.GetAllTopics(true, category, this.Input().Get("label"))
	if err != nil {
		beego.Error(err.Error)
	}
	this.Data["Topics"] = topics
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories

	/*	c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"
		c.TplName = "home.html"

		c.Data["TrueCond"] = true
		c.Data["FalseCond"] = false

		type u struct {
			Name string
			Age  int
			Sex  string
		}
		user := &u{
			Name: "郜瑞仙",
			Age:  18,
			Sex:  "男",
		}
		c.Data["User"] = user
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		c.Data["Nums"] = nums
		c.Data["TplVar"] = "hey guys"

		c.Data["Html"] = "<div>hello beego</div>"
		c.Data["Pipe"] = "<div>hello beego</div>"*/

	//beego.Trace("trace test1")
	//beego.Info("info test1")
	//
	//beego.SetLevel(beego.LevelInformational)
	//beego.Trace("trace test1")
	//beego.Info("info test1")

}
