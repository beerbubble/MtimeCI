package utility

import (
	//"net/http"
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func RundeckRunJob(rundeckServerUrl string, token string, jobId string, args map[string]string) string {
	client := &http.Client{}

	var rundeckurl string

	rundeckurl = rundeckServerUrl + "job" + "/" + jobId + "/" + "run?authtoken=" + token + "&argString="

	var paramstring string

	for key, value := range args {
		paramstring += "-" + key + " " + value + " "
	}

	rundeckurl += url.QueryEscape(paramstring)

	reqest, _ := http.NewRequest("GET", rundeckurl, nil)

	fmt.Println(rundeckurl)

	//reqest, _ := http.NewRequest("GET", "http://10.10.130.221:4440/api/13/job/30eb1ae2-540b-4cff-bcad-976762bc33d2/run?authtoken=E4rNvVRV378knO9dp3d73O0cs1kd0kCd&argString=-url+http%3A%2F%2F192.168.0.25%3A6666%2FMtimeGoConfigWeb%2F20150805164616%2Fconfig-web.tar.gz", nil)
	//reqest, _ := http.NewRequest("GET", "http://10.10.130.221:4440/api/13/job/30eb1ae2-540b-4cff-bcad-976762bc33d2/run?authtoken=E4rNvVRV378knO9dp3d73O0cs1kd0kCd&argString=-url+"+url.QueryEscape(logs[0].Packagepath), nil)

	//reqest.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(reqest)

	//resp, err := http.Get("http://api.m.mtime.cn/PageSubArea/HotPlayMovies.api?locationId=290")
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	if err != nil {
		//handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(body))
	return string(body)
}

func RundeckExecutionInfo(rundeckServerUrl string, token string, executionId string) string {
	client := &http.Client{}

	var rundeckurl string

	rundeckurl = rundeckServerUrl + "execution" + "/" + executionId + "?authtoken=" + token

	reqest, _ := http.NewRequest("GET", rundeckurl, nil)

	fmt.Println(rundeckurl)

	//reqest, _ := http.NewRequest("GET", "http://10.10.130.221:4440/api/13/job/30eb1ae2-540b-4cff-bcad-976762bc33d2/run?authtoken=E4rNvVRV378knO9dp3d73O0cs1kd0kCd&argString=-url+http%3A%2F%2F192.168.0.25%3A6666%2FMtimeGoConfigWeb%2F20150805164616%2Fconfig-web.tar.gz", nil)
	//reqest, _ := http.NewRequest("GET", "http://10.10.130.221:4440/api/13/job/30eb1ae2-540b-4cff-bcad-976762bc33d2/run?authtoken=E4rNvVRV378knO9dp3d73O0cs1kd0kCd&argString=-url+"+url.QueryEscape(logs[0].Packagepath), nil)

	//reqest.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(reqest)

	//resp, err := http.Get("http://api.m.mtime.cn/PageSubArea/HotPlayMovies.api?locationId=290")
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	if err != nil {
		//handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(body))
	return string(body)
}
