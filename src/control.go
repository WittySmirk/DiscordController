package main

import (
	"github.com/go-vgo/robotgo"
)

func shoot()  {
	robotgo.MouseClick("left", false)
}

func aim(){
	robotgo.MouseClick("right", false)
}

func walk(){
	robotgo.KeyToggle("w", "down", "shift")
}

func sprint(){
	robotgo.KeyToggle("w", "down")
}

func jump(){
	robotgo.KeyTap("space")
}

func defuse(){
	robotgo.KeyToggle("4", "down")
}

func walkBack(){
	robotgo.KeyToggle("s", "down", "shift")
}

func a1(){
	robotgo.KeyTap("q")
}

func a2(){
	robotgo.KeyTap("e")
}

func a3(){
	robotgo.KeyTap("c")
}

func ult(){
	robotgo.KeyTap("x")
}

func lL(){
	x, y := robotgo.GetMousePos()
	robotgo.MoveMouseSmooth(x + 200, y)
}

func lR(){
	x, y := robotgo.GetMousePos()
	robotgo.MoveMouseSmooth(x - 200, y)
}

func lU(){
	x, y := robotgo.GetMousePos()
	robotgo.MoveMouseSmooth(x, y - 200)
}

func lD(){
	x, y := robotgo.GetMousePos()
	robotgo.MoveMouseSmooth(x, y + 200)
}

func reload(){
	robotgo.KeyTap("r")
}

func stopWalking(){
	robotgo.KeyToggle("w", "up", "shift")
}

func stopSprinting(){
	robotgo.KeyToggle("w", "up")
}

func stopWalkingB(){
	robotgo.KeyToggle("s", "up", "shift")
}

func stopDefusing(){
	robotgo.KeyToggle("4", "up")
}