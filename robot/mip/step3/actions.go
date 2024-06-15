package main

import (
	"time"
)

func performActions() {
	println("turn left")
	if err := robot.TurnLeft(45, 10); err != nil {
		println(err)
		return
	}

	time.Sleep(3 * time.Second)

	println("turn right")
	if err := robot.TurnRight(135, 10); err != nil {
		println(err)
		return
	}

	time.Sleep(3 * time.Second)

	println("face forward")
	if err := robot.TurnLeft(0, 10); err != nil {
		println(err)
		return
	}

	time.Sleep(3 * time.Second)

	print("stop")
	if err := robot.Stop(); err != nil {
		println(err)
		return
	}
}
