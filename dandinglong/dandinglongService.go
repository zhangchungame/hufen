package dandinglong

import (
	"net/http"
	"fmt"
	"net/http/cookiejar"
	"strings"
	"io/ioutil"
	"encoding/json"
)

type DandinglongService struct {
	client http.Client
}

/**
登录
 */
func (d *DandinglongService)Login()  bool{
	fmt.Println("输入帐号")
	var account,password string
	fmt.Scan(&account)
	fmt.Println("输入密码")
	fmt.Scan(&password)
	mm := make(map[string]string)
	mm["account"] = account
	mm["password"] = password
	postData,_:=json.Marshal(mm)
	req, _ :=http.NewRequest("POST","http://localhost:8080/login",strings.NewReader(string(postData)))

	req.Header.Set("Content-Type", "application/json")

	resp, _ := d.client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(data))
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(string(data)), &m)
	if (err!=nil){
		fmt.Println(err)
		return false
	}
	if(m["code"]=="200"){
		return true
	}else{
		fmt.Println(m["msg"])
		return false
	}
}

/**
注册
 */
func (d *DandinglongService)Register()  {
	fmt.Println("输入帐号")
	var account,password,password2 string
	fmt.Scan(&account)
	fmt.Println("输入密码")
	fmt.Scan(&password)
	fmt.Println("确认密码")
	fmt.Scan(&password2)
	if(password!=password2){
		fmt.Println("两次密码不一致")
	}
	mm := make(map[string]string)
	mm["account"] = account
	mm["password"] = password
	postData,_:=json.Marshal(mm)
	req, _ :=http.NewRequest("POST","http://localhost:8080/login/register",strings.NewReader(string(postData)))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := d.client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(data))
}

func (d *DandinglongService)SaveAccount(account string) bool  {
	mm := make(map[string]string)
	mm["toutiaoAccount"] = account
	postData,_:=json.Marshal(mm)
	req, _ :=http.NewRequest("POST","http://localhost:8080/guanzhu/accountSave",strings.NewReader(string(postData)))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := d.client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(data))
	return true
}


func NewDandinglongService() *DandinglongService {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}
	return &DandinglongService{client: client}
}