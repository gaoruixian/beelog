package controllers

import (
	"beelog/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Get() {

}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid")
	err := models.AddReply(tid, this.Input().Get("nickname"), this.Input().Get("content"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}
func (this *ReplyController) Delete() {
	if !checkAccount(this.Ctx) {
		return
	}

	tid := this.Input().Get("tid")
	rid := this.Input().Get("rid")
	err := models.DeleteReply(rid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}
