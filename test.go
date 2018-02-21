package main

import (
	"net/http/cookiejar"
	"net/http"
	"io/ioutil"
	"encoding/json"
)


func main() {
	gCurCookieJar, _ := cookiejar.New(nil)
	client := http.Client{Jar: gCurCookieJar}
	header:=http.Header{}
	header.Add("Host","www.toutiao.com")
	header.Add("User-Agent","Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:58.0) Gecko/20100101 Firefox/58.0")
	header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	header.Add("Accept-Language","zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	header.Add("Upgrade-Insecure-Requests","1")
	//req, _ :=http.NewRequest("GET","http://crm.easyrong.com/",nil)
	req, _ :=http.NewRequest("GET","https://sso.toutiao.com/",nil)
	req.Header=header
	resp, _ := client.Do(req)
	resp.Body.Close()
	cookieStr,_:=json.Marshal(client.Jar.Cookies(req.URL))
	ioutil.WriteFile("./cookieJson.txt", cookieStr, 0666) //写入文件(字节数组)
}