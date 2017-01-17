package tests

import (
	"api-gateway/models"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/orm"
	"log"
	"path/filepath"
	"runtime"
	"testing"
	"github.com/astaxie/beego/orm"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestInitEngine(t *testing.T) {
	//orm := models.GetOrmer()
	//if orm == nil {
	//	t.FailNow()
	//}

	db := models.GetOrmer()
	var edge models.Edge
	err :=db.Raw("select * from api_edge where type = ? AND U = ?",0,1).QueryRow(&edge)
	if err == nil {
		logs.Informational("Edge.V   : ", edge.V)
	}
}

func TestInsert(t *testing.T) {
	orm := models.GetOrmer()
	account := new(models.Account)
	account.AccessKey = "12312eadqweqe12"
	account.OrganizationName = "kreditplus"
	account.SecretKey = "12312412312312"
	orm.Insert(account)
	if existed := orm.QueryTable("api_account").Filter("access_key", account.AccessKey).Exist(); existed == false {
		t.Fail()
	}
}

func TestSelectEdge(t *testing.T){
	//models.SelectRoleV(0,1)
	//使用QueryBuilder
	db := models.GetOrmer()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("api_account").InnerJoin("api_role").On("api_account.id = api_role.id")

	var maps []orm.Params
	_, err1 := db.Raw(qb.String()).Values(&maps)
	if err1 == nil {
		for _, term := range maps {
			logs.Info(term["organization_name"], ":", term["name"], ":", term["access_key"])
		}
	}

	qb2, _ := orm.NewQueryBuilder("mysql")
	qb2.Select("api_role.name rolename","api_auth.name authname").From("api_role").InnerJoin("api_auth").On("api_auth.id = api_role.id")
	_, err2 := db.Raw(qb2.String()).Values(&maps)
	if err2 == nil {
		for _, term := range maps{
			logs.Info(term["rolename"],":",term["authname"])
		}
	}
}

func TestDB(this *testing.T) {
	db := models.GetOrmer()

	//var edge models.Edge
	//err :=db.Raw("select * from api_edge where type = 0 AND U = ?",1).QueryRow(&edge)
	//if err == nil {
	//	fmt.Println("fdsfsdfsdfsdfsd1   :  ",edge.V)
	//}


	//qb, _ := orm.NewQueryBuilder("mysql")
	//qb.Select("api_account.organization_name","api_role.name").From("api_account").InnerJoin("api_role").On("api_account.id = api_role.id").Where("api_account.id = ?")
	//
	//var account models.Account
	//var role models.Role
	//err1 := db.Raw(qb.String(),1).QueryRow(&account)
	//if err1 == nil {
	//	fmt.Println("fdsfsdfsdfsdfsd2   :  ",account.OrganizationName)
	//}
	//err2 := db.Raw(qb.String(),1).QueryRow(&role)
	//if err2 == nil {
	//	fmt.Println("fdsfsdfsdfsdfsd3   :  ",role.Name)
	//}




	var maps []orm.Params
	num, err := db.Raw("select account.organization_name,role.name from api_account as account left join api_role as role on account.id = role.id where account.id = ? ",1).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["organization_name"])
		fmt.Println(maps[0]["name"])
		fmt.Printf("account :%v \n", maps[0]["organization_name"])
		fmt.Printf("role :%v", maps[0]["name"])
	}

	num1, err1 := db.Raw("select role.name,auth.server_dns from api_role as role left join api_auth as auth on role.id = auth.id where role.id = ? ",1).Values(&maps)
	if err1 == nil && num1 > 0 {
		fmt.Println(maps[0]["name"])
		fmt.Println(maps[0]["server_dns"])
		fmt.Printf("role :%v \n",maps[0]["name"])
		fmt.Printf("auth :%v", maps[0]["server_dns"])
	}


	//account := models.Service{Name: "2name", ServerDns: "12server_dns", IP: "2ip"}
	////insert
	//id, err := db.Insert(&account)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)

	////update
	//account.OrganizationName = "hhds"
	//num, err := db.Update(&account)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	////read one
	//a := models.Account{Id: account.Id}
	//err = db.Read(&a)
	//fmt.Printf("ERR: %v\n", err)
	//
	////delete
	//num, err = db.Delete(&a)
	//fmt.Printf("NUM: %D, ERR: %v\n", num, err)
}

func TestSelect(this *testing.T) {
	var organization_name = ``
	db, err := sql.Open("mysql", "work:LwTgrtAd1hqLBn2x@tcp(115.159.124.203:3306)/advance_api_gateway?charset=utf8")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	//defer db.Close()
	rows, err := db.Query("SELECT organization_name FROM api_account WHERE id = ? ", 13)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for rows.Next() {
		err := rows.Scan(&organization_name)
		if err != nil {
			log.Fatal(err)
		}
		logs.Info(" organization_name :" + organization_name)
		log.Println(" 23333 " + organization_name)
	}
	err = rows.Err()
	if err != nil {
		logs.Info("没有您要查的数据")
		log.Fatal(err) //打印完信息 直接退出  不再继续下面的逻辑操作
	}
}
