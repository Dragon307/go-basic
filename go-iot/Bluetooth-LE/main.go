package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"os"
	"time"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	battery := ble.NewBatteryDriver(bleAdaptor)

	work := func() {
		gobot.Every(5*time.Second, func() {
			fmt.Println("Battery level:", battery.GetBatteryLevel())
		})
	}

	robot := gobot.NewRobot("bleBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{battery},
		work,
		)

	robot.Start()
}

// 70:F0:87:2F:0A:E4
// 24-F6-77-11-A3-5B