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
	"encoding/xml"
	//"net/url"
	"time"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	/*
	   	var data = []byte(`<executions count='1'>
	     <execution id='1081' href='http://192.168.50.20:4440/execution/follow/1081' status='running' project='CI'>
	       <user>admin</user>
	       <date-started unixtime='1441462368695'>2015-09-05T14:12:48Z</date-started>
	       <job id='305afa2d-82eb-435d-88ee-2b1d12b353cb' averageDuration='24812'>
	         <name>Config-Web</name>
	         <group>Go</group>
	         <project>CI</project>
	         <description></description>
	       </job>
	       <description>#!/bin/bash
	   pwd
	   export GOROOT=/opt/go/go
	   export GOPATH=/home/mtimegit/go
	   export PATH=$PATH:/opt/go/go/bin:/home/mtimegit/go/bin

	   sh contract.sh
	   echo @option.JOB_NAME@
	   /home/mtimegit/go/bin/config-web/config-web -s quit
	   /home/mtimegit/go/bin/mtimewall-service/mtimewall-service -s quit
	   /home/mtimegit/go/bin/captcha-service/captcha-service -s quit
	   rm -rf /home/mtimegit/go/src/mtime.com/basis
	   git clone -b @option.Branch_NAME@ git@192.168.50.30:go/basis /home/mtimegit/go/src/mtime.com/basis
	   gobuild -v mtime.com/basis/config-web
	   status=$?
	   echo "gobuild command exit stats - $status"
	   if [ $status -gt 0 ]
	   then
	   	echo fail
	   	exit 1
	   fi
	   #输出git分支名和Hash值
	   cd /home/mtimegit/go/src/mtime.com/basis
	   git rev-parse --abbrev-ref HEAD &gt; /home/mtimegit/go/bin/config-web/GitBranchName
	   git rev-parse HEAD &gt; /home/mtimegit/go/bin/config-web/GitBranchHash
	   /home/mtimegit/go/bin/config-web/config-web

	   #将原始程序部署到公共目录
	   rm -rf /mnt/025/MtimeGoConfigWeb/@option.BUILD_NUMBER@
	   mkdir /mnt/025/MtimeGoConfigWeb/@option.BUILD_NUMBER@
	   cp -r /home/mtimegit/go/bin/config-web /mnt/025/MtimeGoConfigWeb/@option.BUILD_NUMBER@/

	   gobuild -v mtime.com/basis/mtimewall-service
	   if [ $? -gt 0 ]
	   then
	   	echo fail
	   	exit 1
	   fi
	   /home/mtimegit/go/bin/mtimewall-service/mtimewall-service

	   gobuild -v mtime.com/basis/captcha-service
	   if [ $? -gt 0 ]
	   then
	   	echo fail
	   	exit 1
	   fi
	   /home/mtimegit/go/bin/captcha-service/captcha-service</description>
	       <argstring />
	     </execution>
	   </executions>`)
	*/

	o := orm.NewOrm()

	var logs []*models.Executionlog
	num, err := o.QueryTable("Executionlog").Filter("id", 1).All(&logs)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	args := map[string]string{"BUILD_NUMBER": "10", "Branch_NAME": "develop"}

	response := utility.RundeckRunJob("http://192.168.50.20:4440/api/13/", "E4rNvVRV378knO9dp3d73O0cs1kd0kCd", "305afa2d-82eb-435d-88ee-2b1d12b353cb", args)

	r := models.RunJobExecutions{}
	xml_err := xml.Unmarshal([]byte(response), &r)

	if xml_err != nil {
		fmt.Printf("error: %v", xml_err)
		this.Data["json"] = xml_err //url.QueryEscape(logs[0].Packagepath)
		this.ServeJson()
	}

	if r.Exs != nil {
		fmt.Println(r.Exs[0].Id)
	}

	fmt.Println(r)

	i := 0
	for i < 1 {
		r2 := models.RunJobExecutions{}
		response_2 := utility.RundeckExecutionInfo("http://192.168.50.20:4440/api/13/", "E4rNvVRV378knO9dp3d73O0cs1kd0kCd", r.Exs[0].Id)

		xml_err2 := xml.Unmarshal([]byte(response_2), &r2)

		if xml_err2 != nil {
			fmt.Printf("error: %v", xml_err2)
			this.Data["json"] = xml_err2 //url.QueryEscape(logs[0].Packagepath)
			this.ServeJson()
		}

		if r2.Exs != nil {
			fmt.Println(r.Exs[0].Id)
		}

		fmt.Println(r2)

		if r2.Exs[0].Status == "succeeded" {
			i = 1
		}

		time.Sleep(time.Second * 3)
	}

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

	this.Data["json"] = r //url.QueryEscape(logs[0].Packagepath)
	this.ServeJson()

	//this.Data["Website"] = "beego.me" //+ string(body)
	//this.Data["Email"] = "astaxie@gmail.com"
	//this.TplNames = "index.tpl"
}
