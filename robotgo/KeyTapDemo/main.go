package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	time.Sleep(time.Second * 5)
	robotgo.TypeStr("Hello World")
	robotgo.TypeStr("测试")
	robotgo.TypeString("测试")

	robotgo.TypeStr("山达尔星新星军团, galaxy. こんにちは世界.")
	robotgo.Sleep(1)

	// ustr := uint32(robotgo.CharCodeAt("测试", 0))
	// robotgo.UnicodeType(ustr)

	robotgo.KeyTap("enter")
	// robotgo.TypeString("en")
	robotgo.KeyTap("i", "alt", "command")

	arr := []string{"alt", "command"}
	robotgo.KeyTap("i", arr)

	robotgo.WriteAll("测试")
	text, err := robotgo.ReadAll()
	if err == nil {
		fmt.Println(text)
	}
}
