package controllers

import (
	"zhiwen2/zhiwen2/models"

	"github.com/astaxie/beego"
)

type PaperController struct {
	beego.Controller
}

func (c *PaperController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Number"] = id
	c.Data["Title"] = models.GetPaperTitleById(id)
	c.Data["Content"] = models.GetPaperContentById(id)

	c.TplName = "paper.tpl"
}
