package main

import (
	"fmt"
	"time"

	"github.com/tpisntgod/Service_Calculation/Hw6_Week16/client"
)

func syncRequest() {
	startTime := time.Now()
	r1 := client.HTTPGetSync("love")
	r2 := client.HTTPGetSync("learning")
	r3 := client.HTTPGetSync("go")
	r4 := client.HTTPGetSync("language")
	r5 := client.HTTPGetSync("thank")
	fmt.Printf("%s\n%s%s%s%s%s", "Translate Result:", r1, r2, r3, r4, r5)
	fmt.Println("Time consumed:", time.Since(startTime))
}

func asyncRequest() {
	startTime := time.Now()
	ch := make(chan string)
	go client.HTTPGetAsync("love", ch)
	go client.HTTPGetAsync("learning", ch)
	go client.HTTPGetAsync("go", ch)
	go client.HTTPGetAsync("language", ch)
	go client.HTTPGetAsync("thank", ch)
	fmt.Println("translate result:\n", <-ch, <-ch, <-ch, <-ch, <-ch)
	fmt.Println("Time consumed:", time.Since(startTime))
}

func main() {
	fmt.Printf("%s\n\n", "分别使用同步和异步方法向百度翻译发送5个翻译请求")
	fmt.Printf("%s\n", "使用 go HTTPClient 实现图 6-2 的 Naive Approach")

	syncRequest()

	fmt.Printf("%s\n\n", "------------我是分割线------------")
	fmt.Printf("%s\n", "为每个 HTTP 请求设计一个 goroutine ，利用 Channel 搭建基于消息的异步机制，实现图 6-3")

	asyncRequest()
}
