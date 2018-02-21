package main

import (
	"fmt"
	"github.com/fedesog/webdriver"
	"time"
)

func main() {
	chromeDriver := webdriver.NewChromeDriver("./chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		fmt.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	sessiondanding, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		fmt.Println(err)
	}
	err = sessiondanding.Url("http://www.dandinglong.site")
	if err != nil {
		fmt.Println(err)
	}
	session, err := chromeDriver.NewSession(desired, required)
	var i int
	for i=0;i<1;i++{
		if err != nil {
			fmt.Println(err)
		}
		err = session.Url("https://www.toutiao.com/i6522432980709802510/")
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(2 * time.Second)
		el, _ :=session.FindElement(webdriver.TagName,"body");
		bodysize, _ :=el.Size()
		fmt.Println(bodysize)
		cookie,_:=session.GetCookies()
		fmt.Println(cookie)
		key:="\ue00f"
		session.SendKeysOnActiveElement(key)
		time.Sleep(1 * time.Second)
		session.SendKeysOnActiveElement(key)
		time.Sleep(4 * time.Second)
	}

}