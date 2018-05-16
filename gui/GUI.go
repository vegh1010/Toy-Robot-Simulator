// package toy_robot_gui is a gui package that manages gui output on console.
package toy_robot_gui

import (
	"github.com/vegh1010/Toy-Robot-Simulator/lib"
	"fmt"
)

//Variables to manage toy robot's movement
var (
	Box = "o"
	//facing id
	FACING = map[int]string{
		1: "^",
		2: ">",
		3: "v",
		4: "<",
	}
)

func Output(board toy_robot_lib.Board, wallE toy_robot_lib.Robot) {
	for y := board.Height-1; y >= 0; y-- {
		var row []string
		for x := 0; x < board.Height; x++ {
			if wallE.X == x && wallE.Y == y {
				row = append(row, FACING[wallE.Face])
			} else {
				row = append(row, Box)
			}
		}
		fmt.Println(row)
	}
	fmt.Println()
}
