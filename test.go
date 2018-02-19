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
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		fmt.Println(err)
	}
	err = session.Url("http://www.dandinglong.site")
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(2 * time.Second)
	//el, _ :=session.FindElement(webdriver.ClassName,"search-field");
	key:="\ue00f"
	session.SendKeysOnActiveElement(key)
	value1, value2 := 4, 7
	script := "return arguments[0] + arguments[1]"
	res, err := session.ExecuteScript(script, []interface{}{value1, value2})
	fmt.Println(string(res))
	time.Sleep(4 * time.Second)
	session.Delete()
	chromeDriver.Stop()
}