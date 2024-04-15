package main

import "github.com/beego/beego/v2/server/web"

func main() {
	web.BConfig.Listen.EnableAdmin = true
	web.Run()
}
