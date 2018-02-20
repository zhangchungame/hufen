package main

import (
	"runtime"
	"hufen/dandinglong"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//toutiaoObj:=toutiao.NewToutiao()
	//userInfo:=toutiaoObj.Login()
	//fmt.Println(userInfo)
	//userId:=strconv.Itoa(userInfo.UserId)
	dandingObj:=dandinglong.DandinglongService{}
	//danding.Register()
	if dandingObj.Login(){
		dandingObj.SaveAccount("asd")
	}
}
