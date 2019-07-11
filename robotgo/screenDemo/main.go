package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	time.Sleep(5*time.Second)
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)
	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color----", color)
}
