package controllers

import (
	"beerbubble/MtimeCI/models"
	"beerbubble/MtimeCI/utility"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	//"io/ioutil"
	//"net/http"
	"net/url"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	o := orm.NewOrm()

	var logs []*models.Executionlog
	num, err := o.QueryTable("Executionlog").Filter("id", 1).All(&logs)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	args := map[string]string{"BUILD_NUMBER": "10", "Branch_NAME": "develop"}

	utility.RunRundeckJob("http://192.168.50.20:4440/api/13/", "305afa2d-82eb-435d-88ee-2b1d12b353cb", "E4rNvVRV378knO9dp3d73O0cs1kd0kCd", args)

	/*
		client := &http.Client{}

		//reqest, _ := http.NewRequest("GET", "http://10.10.130.221:4440/api/13/job/30eb1ae2-540b-4cff-bcad-976762bc33d2/run?authtoken=E4rNvVRV378knO9dp3d73O0cs1kd0kCd&argString=-url+http%3A%2F%2F192.168.0.25%3A6666%2FMtimeGoConfigWeb%2F20150805164616%2Fconfig-web.tar.gz", nil)
		reqest, _ := http.NewRequest("GET", "http://10.10.130.221:4440/api/13/job/30eb1ae2-540b-4cff-bcad-976762bc33d2/run?authtoken=E4rNvVRV378knO9dp3d73O0cs1kd0kCd&argString=-url+"+url.QueryEscape(logs[0].Packagepath), nil)

		//reqest.Header.Set("Connection", "keep-alive")

		resp, err := client.Do(reqest)

		//resp, err := http.Get("http://api.m.mtime.cn/PageSubArea/HotPlayMovies.api?locationId=290")
		//req.Header.Add("If-None-Match", `W/"wyzzy"`)
		if err != nil {
			//handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
	*/

	this.Data["json"] = url.QueryEscape(logs[0].Packagepath)
	this.ServeJson()

	this.Data["Website"] = "beego.me" //+ string(body)
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
