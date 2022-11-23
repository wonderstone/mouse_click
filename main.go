// for annoying online training courses that require you to click the mouse
package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.MouseSleep = 300
	sleeptime := 60*3 + 1
	robotgo.Move(200, 250)
	robotgo.Click("left", true)
	// sleep for 5 minutes
	// robotgo.Sleep(sleeptime)
	robotgo.Sleep(1)
	// iter n times
	for i := 0; i < 10; i++ {
		robotgo.MoveRelative(0, 61)
		robotgo.Click("left", true)
		robotgo.Sleep(sleeptime)
	}

}

func CurrentMousePos() {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
}
