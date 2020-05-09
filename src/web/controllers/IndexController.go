package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"gocloud/datamodels"
)

type IndexController struct {
	Ctx iris.Context
}

func (this *IndexController) Get() mvc.View{
	sess:=sessions.Get(this.Ctx)
	if auth,err:=sess.GetBoolean("authenticated");!auth||err!=nil{
		return mvc.View{
			Layout: "shared/layout.fw.html",
			Name: "login/login.html",
		}
	}
	this.Ctx.Application().Logger().Debug("index get")

	user ,ok:=(sess.Get("user")).(*datamodels.UserModel)
	if !ok{
		this.Ctx.Application().Logger().Error("get user err by sesssion")
	}

	return mvc.View{
		Name: "index/index.html",
		Data: iris.Map{
			"username":user.Username,
		},
	}
}

func (this *IndexController) GetIndex() {
	//test
	a :=[]int{1,2,3}
	this.Ctx.JSON(a)
}

