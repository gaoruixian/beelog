package controllers

import (
	"beelog/models"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"

}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	topics, err := models.GetAllTopics(false, "", "")
	if err != nil {
		beego.Error(err.Error)
	} else {
		this.Data["Topics"] = topics
	}

}
func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")
	category := this.Input().Get("category")
	label := this.Input().Get("label")

	//获取附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	if fh != nil {
		//保存附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}

	}
	if len(tid) == 0 {
		err = models.AddTopic(title, category, label, content, attachment)
	} else {
		err = models.ModifyTopic(tid, title, category, label, content, attachment)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)

}
func (this *TopicController) View() {
	this.TplName = "topic_view.html"

	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
	} else {
		this.Data["Topic"] = topic
		this.Data["Labels"] = strings.Split(topic.Labels, " ")

		//方便根据id修改对应的文章
		this.Data["Tid"] = this.Ctx.Input.Param("0")
		replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
		if err != nil {
			beego.Error(err)
			return
		}
		this.Data["Replies"] = replies
		this.Data["IsLogin"] = checkAccount(this.Ctx)
	}
}
func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"
	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}
func (this *TopicController) Delete() {
	//校验是否登录
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 302)
}
