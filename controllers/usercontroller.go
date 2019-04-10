package controllers

import (
	"github.com/xiliangMa/restapi/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

// Users object api list
type UserController struct {
	beego.Controller
}

// @Title GetUser
// @Description get Users
// @Param mobile query string false "User mobile"
// @Param email query string false "User email"
// @Param page query int 1 false "page"
// @Param number query int 20 false "page"
// @Success 200 {object} models.Result
// @router / [post]
func (this *UserController) UserList() {
	mobile := this.GetString("mobile")
	email := this.GetString("email")
	number, _ := this.GetInt("number")
	page, _ := this.GetInt("page")
	this.Data["json"] = models.GetUserList(mobile, email, page, number)
	this.ServeJSON(false)

}

// @Title AddUser
// @Description dd User
// @Param User body models.User true "User object"
// @Success 200 {object} models.Result
// @router /addUser [post]
func (this *UserController) AddUser() {
	var h models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &h)
	this.Data["json"] = models.AddUser(&h)
	this.ServeJSON(false)

}



// @Title AddUser
// @Description dd User
// @Param id path int true "User id"
// @Success 200 {object} models.Result
// @router /:id [delete]
func (this *UserController) DeleteUser() {
	id, _ := this.GetInt(":id")
	this.Data["json"] = models.DeleteUser(id)
	this.ServeJSON(false)

}