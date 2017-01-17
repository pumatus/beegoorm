package dao

import (
	"api-gateway/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

func AccountService(){
	//使用QueryBuilder
	db := models.GetOrmer()
	qb, _ := orm.NewQueryBuilder("mysql")
	var maps []orm.Params

	//api_account  api_role
	qb.Select("*").From("api_account").InnerJoin("api_role").On("api_account.id = api_role.id")
	_, err1 := db.Raw(qb.String()).Values(&maps)
	if err1 == nil {
		for _, term := range maps {
			logs.Info(term["organization_name"], ":", term["name"], ":", term["access_key"])
		}
	}

	//api_role  api_auth
	qb2, _ := orm.NewQueryBuilder("mysql")
	qb2.Select("api_role.name rolename","api_auth.name authname").From("api_role").InnerJoin("api_auth").On("api_auth.id = api_role.id")
	_, err2 := db.Raw(qb2.String()).Values(&maps)
	if err2 == nil {
		for _, term := range maps{
			logs.Info(term["rolename"],":",term["authname"])
		}
	}

}