package main

import (
	"math"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot/v2/platforms/joystick"
)

type pair struct {
	x float64
	y float64
}

var (
	joystickAdaptor *joystick.Adaptor
	stick           *joystick.Driver

	leftX, leftY, rightX, rightY atomic.Value
)

const (
	offset = 32767.0
	max    = 30
)

func startJoystick() {
	joystickAdaptor = joystick.NewAdaptor("0")
	stick = joystick.NewDriver(joystickAdaptor, joystick.Dualshock3)

	joystickAdaptor.Connect()
	stick.Start()

	leftX.Store(float64(0.0))
	leftY.Store(float64(0.0))
	rightX.Store(float64(0.0))
	rightY.Store(float64(0.0))

	stick.On(joystick.TrianglePress, func(data interface{}) {
		robot.GetUp(5)
	})

	stick.On(joystick.LeftX, func(data interface{}) {
		val := float64(data.(int))
		leftX.Store(val)
	})

	stick.On(joystick.LeftY, func(data interface{}) {
		val := float64(data.(int))
		leftY.Store(val)
	})

	stick.On(joystick.RightX, func(data interface{}) {
		val := float64(data.(int))
		rightX.Store(val)
	})

	stick.On(joystick.RightY, func(data interface{}) {
		val := float64(data.(int))
		rightY.Store(val)
	})

	go handleLeftJoystick()
	go handleRightJoystick()
}

func handleRightJoystick() {
	var xs, ys, stopped bool

	for {
		rightStick := getRightStick()

		switch {
		case rightStick.y < -10:
			ys, stopped = false, false
			robot.DriveForward(ValidatePitch(rightStick.y, offset), 20)
		case rightStick.y > 10:
			ys, stopped = false, false
			robot.DriveBackward(ValidatePitch(rightStick.y, offset), 20)
		default:
			ys = true
		}

		switch {
		case rightStick.x > 10:
			xs, stopped = false, false
			robot.TurnRight(ValidatePitch(rightStick.x, offset), 5)
		case rightStick.x < -10:
			xs, stopped = false, false
			robot.TurnLeft(ValidatePitch(rightStick.x, offset), 5)
		default:
			xs = true
		}

		if ys && xs && !stopped {
			robot.Stop()
			stopped = true
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func handleLeftJoystick() {
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func getLeftStick() pair {
	s := pair{x: 0, y: 0}
	s.x = leftX.Load().(float64)
	s.y = leftY.Load().(float64)
	return s
}

func getRightStick() pair {
	s := pair{x: 0, y: 0}
	s.x = rightX.Load().(float64)
	s.y = rightY.Load().(float64)
	return s
}

// ValidatePitch helps validate pitch values such as those created by
// a joystick to values between 0-30
func ValidatePitch(data float64, offset float64) uint8 {
	value := math.Abs(data) / offset
	if value >= 0.1 {
		if value <= 1.0 {
			return uint8((float64(int(value*max)) / max) * max)
		}
		return 30
	}
	return 0
}
