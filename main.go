package main

import (
	"runtime"
	"hufen/toutiao"
	"fmt"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	toutiaoObj:=toutiao.NewToutiao()
	userInfo:=toutiaoObj.Login()
	fmt.Println(userInfo)
	userId:=strconv.Itoa(userInfo.UserId)
	fmt.Println("userId=",userId)
	toutiaoObj.Guanzhu("55153509430")


	//dandingObj:=dandinglong.NewDandinglongService()
	//
	////danding.Register()
	//if dandingObj.Login(){
	//	//userInfo:=toutiaoObj.Login()
	//	//fmt.Println(userInfo)
	//	//userId:=strconv.Itoa(userInfo.UserId)
	//	//dandingObj.SaveAccount(userId)
	//	//toutiaoObj.SaveCookies()
	//}
}
