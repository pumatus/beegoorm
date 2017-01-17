package controllers

import (
	"api-gateway/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"strconv"
)

// Operations about service
type RegisterController struct {
	beego.Controller
	err  error
	data interface{}
}

func (c *RegisterController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
}

// @router /register/:access_key [get]
func (c *RegisterController) Get() {
	my := "hhd"
	c.Data["json"] = "Hello World - " + my
	c.ServeJSON()
}

// @router /register/:access_key [post]
func (c *RegisterController) RegistryAuth() {
	name := c.Input().Get("name")
	server_dns := c.Input().Get("server_dns")
	ip := c.Input().Get("i_p")
	post_id := c.Input().Get("id")

	id, _ := strconv.ParseUint(post_id, 10, 64)
	valid := validation.Validation{}
	auth := models.Auth{Name: name, ServerDns: server_dns, IP: ip, Id: id}
	b, _ := valid.Valid(&auth)
	if !b {
		for _, err := range valid.Errors {
			logs.Info(err.Key, " | ", err.Message)
		}
	} else {
		models.InsertorUpdateAuth(auth)
		c.Data["json"] = map[string]string{
			"errorCode": "200",
			"data":      "success",
		}
		c.ServeJSON()
	}
}

// @router /register/account [post]
func (this *RegisterController) RegistryAccount() {
	accessKey := this.Input().Get("accessKey")
	secretKey := this.Input().Get("secretKey")
	organizationName := this.Input().Get("organizationName")
	post_id := this.Input().Get("id")

	id, _ := strconv.ParseUint(post_id, 10, 64)
	valid := validation.Validation{}
	account := models.Account{AccessKey: accessKey, SecretKey: secretKey, OrganizationName: organizationName, Id: id}
	b, _ := valid.Valid(&account)
	if !b {
		for _, err := range valid.Errors {
			logs.Info(err.Key, " | ", err.Message)
		}
	} else {
		models.InsertAccount(account)
		this.Data["json"] = map[string]string{
			"errorCode": "200",
			"data":      "success",
		}
		this.ServeJSON()
	}
}

// @router /register/edge [post]
func (this *RegisterController) SelectEdge() {
	var edgetype models.EdgeType
	if types, _ := this.GetInt("type"); types == 0 {
		edgetype = models.ACCOUNT2ROLE
	} else {
		edgetype = models.ROLE2AUTH
	}
	u := this.Input().Get("u")
	if u == "" {
		return
	} else {

		edge_u, _ := strconv.ParseUint(u, 10, 64)
		models.SelectRoleV(edgetype, edge_u)

		this.Data["json"] = map[string]string{
			"errorCode": "200",
			"data":      "success",
		}
		this.ServeJSON()
	}
}

//函数结束时,组装成json结果返回
func (this *RegisterController) Finish() {
	this.Data["json"] = map[string]string{
		"errorCode": "201",
		"data":      "fail",
	}
	this.ServeJSON()
	this.StopRun()
}