package main

import (
	browser "github.com/EDDYCJY/fake-useragent"
	"log"
	"time"
)

func main()  {
	// 推荐使用
	random := browser.Random()
	log.Printf("Random: %s", random)

	chrome := browser.Chrome()
	log.Printf("IE: %s", chrome)

	internetExplorer := browser.InternetExplorer()
	log.Printf("IE: %s", internetExplorer)

	firefox := browser.Firefox()
	log.Printf("Firefox: %s", firefox)

	safari := browser.Safari()
	log.Printf("Safari: %s", safari)

	android := browser.Android()
	log.Printf("Android: %s", android)

	macOSX := browser.MacOSX()
	log.Printf("MacOSX: %s", macOSX)

	ios := browser.IOS()
	log.Printf("IOS: %s", ios)

	linux := browser.Linux()
	log.Printf("Linux: %s", linux)

	iphone := browser.IPhone()
	log.Printf("IPhone: %s", iphone)

	ipad := browser.IPad()
	log.Printf("IPad: %s", ipad)

	computer := browser.Computer()
	log.Printf("Computer: %s", computer)

	mobile := browser.Mobile()
	log.Printf("Mobile: %s", mobile)

	// 定制
	// 可以调整抓取数据源的最大页数、时间间隔以及最大超时时间

	client := browser.Client{
		MaxPage: 3,
		Delay:   200 * time.Millisecond,
		Timeout: 10 * time.Second,
	}
	cache := browser.Cache{}
	b := browser.NewBrowser(client, cache)

	log.Printf("Client: %s", b)
}

