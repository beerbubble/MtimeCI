package utility

import (
	//"net/http"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/url"
)

func ViewLogin(c *context.Context) {
	cookie := c.GetCookie("MtimeCIUserId")
	if len(cookie) <= 0 {
		c.Redirect(302, "/login?url="+url.QueryEscape(c.Input.Uri()))
	}
	beego.Informational(cookie)
}
