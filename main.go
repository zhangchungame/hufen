package main

import (
	"hufen/toutiao"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	toutiao:=toutiao.NewToutiao()
	toutiao.Login()
}
