package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error401() {
	this.Data["json"] = map[string]interface{}{"code": 401, "message": "NO AUTHORIZATION", "data": nil, "extra": nil}
	this.ServeJSON()
}
