package routers

import (
	"Api_Crud_Usuario/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSInclude("/documento",
			beego.NSInclude(
				&controllers.DocumentoController{},
			),),
		beego.

	)
	beego.AddNamespace(ns)
}
