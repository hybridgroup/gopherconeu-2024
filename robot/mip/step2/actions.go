package main

import (
	"image/color"
	"time"
)

var (
	red   = color.RGBA{R: 255, G: 0, B: 0}
	green = color.RGBA{R: 0, G: 255, B: 0}
	blue  = color.RGBA{R: 0, G: 0, B: 255}
)

func performActions() {
	println("chest led red")
	if err := robot.FlashChestLED(red, 10, 10); err != nil {
		println(err)
		return
	}

	time.Sleep(3 * time.Second)

	println("chest led green")
	if err := robot.FlashChestLED(green, 10, 10); err != nil {
		println(err)
		return
	}

	time.Sleep(3 * time.Second)

	println("chest led blue")
	if err := robot.FlashChestLED(blue, 10, 10); err != nil {
		println(err)
		return
	}

	time.Sleep(3 * time.Second)
}
