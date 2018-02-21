package toutiao

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"net/http/cookiejar"
	"regexp"
	"encoding/base64"
	"time"
	"strconv"
	"wiki.ruokuai.com/ApiDemo_Go.ashx/rkdama"
	"os"
	"strings"
	"encoding/json"
)

type ToutiaoService struct {
	client http.Client
}

func (tt *ToutiaoService) Login() ToutiaoLoginReturn{
	header:=http.Header{}
	header.Add("Host","www.toutiao.com")
	header.Add("User-Agent","Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:58.0) Gecko/20100101 Firefox/58.0")
	header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	header.Add("Accept-Language","zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	header.Add("Upgrade-Insecure-Requests","1")
	//req, _ :=http.NewRequest("GET","http://crm.easyrong.com/",nil)
	req, _ :=http.NewRequest("GET","https://sso.toutiao.com/",nil)
	req.Header=header
	resp, _ := tt.client.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	aaa:=string(data)
	r, _ := regexp.Compile("captcha: '(.*)'")
	bbb := r.FindStringSubmatch(aaa)
	fileName := "./yanzhengma/" + saveImageFile(bbb[1])
	creatreurlsult, err := rkdama.RKCreate("dandinglong", "Qwert12345", "3040", "60", "1", "b40ffbee5c1cf4e38028c197eb2fc751", fileName)
	if err == nil {
		fmt.Println("答题结果:" + creatreurlsult.Result)
		fmt.Println("结果ID:" + creatreurlsult.Id)
	} else {
		fmt.Println(err)
	}

	//获取验证码
	req, _ =http.NewRequest("GET","https://sso.toutiao.com/send_activation_code/?mobile=13681736848&captcha="+creatreurlsult.Result+"&type=24",nil)
	req.Header=header
	resp, _ = tt.client.Do(req)
	data, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(data))
	var code string
	fmt.Println("输入code")
	fmt.Scan(&code)
	req, _ =http.NewRequest("POST","https://sso.toutiao.com/quick_login/",strings.NewReader("mobile=13681736848&code="+code+"&account=&password=&captcha="+creatreurlsult.Result+"&is_30_days_no_login=false&service=https://www.toutiao.com/"))

	req.Header=header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, _ = tt.client.Do(req)

	data, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var m  ToutiaoLoginReturn
	err = json.Unmarshal([]byte(string(data)), &m)
	if (err!=nil){
		fmt.Println(err)
	}
	return m
}
func (tt *ToutiaoService)Guanzhu(userId string) bool {
	req,_:=http.NewRequest("POST","https://www.toutiao.com/c/user/follow/",strings.NewReader("user_id="+userId))
	resp, _ := tt.client.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var m  =make(map[string]interface{})
	err := json.Unmarshal([]byte(string(data)), &m)
	if(err!=nil){
		fmt.Println("error",err)
		return false
	}
	fmt.Println(m)
	return true
}

func (tt *ToutiaoService)SaveCookies()  {
	req, _ :=http.NewRequest("GET","https://sso.toutiao.com/",nil)
	cookieStr,_:=json.Marshal(tt.client.Jar.Cookies(req.URL))
	ioutil.WriteFile("./cookieJson.txt", cookieStr, 0666) //写入文件(字节数组)
}

func (tt *ToutiaoService)LoadCookies()  {
	req, _ :=http.NewRequest("GET","https://sso.toutiao.com/",nil)
	res,_:=ioutil.ReadFile("./cookieJson.txt") //写入文件(字节数组)
	var tmpJar []*http.Cookie
	err:=json.Unmarshal(res,&tmpJar)
	if err!=nil{
		fmt.Println(err)
		return
	}
	tt.client.Jar.SetCookies(req.URL,tmpJar)
}

/**
保存文件返回文件名
 */
func saveImageFile(data string) string {
	t := time.Now()
	fileName := strconv.FormatInt(t.UnixNano(), 10) + ".gif"
	info, _ := base64.StdEncoding.DecodeString(data)
	mkdir("./yanzhengma")
	ioutil.WriteFile("./yanzhengma/"+fileName, info, 0666) //写入文件(字节数组)
	return fileName
}

func mkdir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
	}
}
func NewToutiao() *ToutiaoService {
	gCurCookieJar, _ := cookiejar.New(nil)
	client := http.Client{Jar: gCurCookieJar}
	return &ToutiaoService{client: client}
}
