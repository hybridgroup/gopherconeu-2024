package main

import (
	"time"
)

func performActions() {
	println("turn around")
	if err := robot.TurnLeft(45, 5); err != nil {
		println(err)
		return
	}

	time.Sleep(3 * time.Second)

	println("drive forward")
	if err := robot.DriveForward(20, 100); err != nil {
		println(err)
		return
	}

	time.Sleep(3 * time.Second)

	println("turn back around")
	if err := robot.TurnRight(45, 5); err != nil {
		println(err)
		return
	}

	time.Sleep(3 * time.Second)

	println("drive forward")
	if err := robot.DriveForward(20, 100); err != nil {
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
