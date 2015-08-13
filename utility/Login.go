package utility

import (
	//"net/http"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func ViewLogin(c *context.Context) {
	cookie := c.GetCookie("MtimeCIUserId")
	if len(cookie) <= 0 {
		c.Redirect(302, "/login")
	}
	beego.Informational(cookie)
}
