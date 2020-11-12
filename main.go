package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	//if any MouseCursorKoordinate moves more than cursorExitDelta,
	//the programm will terminate
	cursorExitDelta      = 50
	exitConditionSleepMS = 100
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func exitCondition() {
	cursorCurrentPosX, cursorCurrentPosY := robotgo.GetMousePos()

	for {
		bufferX, bufferY := robotgo.GetMousePos()
		if (abs(bufferX-cursorCurrentPosX) > cursorExitDelta) ||
			(abs(bufferY-cursorCurrentPosY) > cursorExitDelta) {
			os.Exit(0)
		}
		cursorCurrentPosX = bufferX
		cursorCurrentPosY = bufferY
		time.Sleep(exitConditionSleepMS * time.Millisecond)
	}
}

func main() {
	screenX, screenY := robotgo.GetScreenSize()
	go exitCondition()
	for {
		//lets be safe and only move the mouse in between half the screen size
		robotgo.MoveMouseSmooth(rand.Intn(screenX/2), rand.Intn(screenY/2), 1.0, 100.0)
	}
}
