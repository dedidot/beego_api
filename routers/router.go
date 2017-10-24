// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"skripsih/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/kv", &controllers.KvController{}, "get:GetAll")
	beego.Router("/kv/:id", &controllers.KvController{}, "get:GetOne;post:Post;put:Put;delete:Delete")
	
	beego.Router("agenda", &controllers.AgendaController{}, "get:GetAgenda")
}
