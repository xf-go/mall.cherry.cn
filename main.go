package main

import "mall.cherry.cn/server/router"

func main() {
	r := router.SetupRouter()
	r.Run()
}
