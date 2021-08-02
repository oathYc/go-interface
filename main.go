package main

import (
	"ihtest/boot"
	"ihtest/library/global"
	"log"
)

func main() {
	if err := boot.Boot(); nil != err {
		log.Fatalln(": ", err.Error())
	}
	global.Log.Info("初始化完成")
}
