// package toy_robot_lib is a library package that contains data struct, reference functions
// and unit test.
package toy_robot_lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//test robot's default placement if command is invalid
func TestDefaultRobotCoordinate(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 1,2,NORTHs")

	assert.Equal(t, "0,0,NORTH", wallE.GetReport())
}

//test set robot's placement
func TestInitRobotCoordinate(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 1,2,NORTH")

	assert.Equal(t, "1,2,NORTH", wallE.GetReport())
}

//test robot turn left
func TestRobotTurnLeft(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "LEFT")

	assert.Equal(t, "0,0,WEST", wallE.GetReport())
}

//test robot turn right
func TestRobotTurnRight(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "RIGHT")

	assert.Equal(t, "0,0,EAST", wallE.GetReport())
}

//test robot turn left twice
func TestRobotTurnLeftTwice(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "LEFT")
	wallE.Move(board, "LEFT")

	assert.Equal(t, "0,0,SOUTH", wallE.GetReport())
}

//test robot turn right twice
func TestRobotTurnRightTwice(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "RIGHT")
	wallE.Move(board, "RIGHT")

	assert.Equal(t, "0,0,SOUTH", wallE.GetReport())
}

//test robot turning and report
func TestRobotTurnAndReport(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 1,5,NORTH")
	wallE.Move(board, "RIGHT")
	wallE.Move(board, "RIGHT")
	wallE.Move(board, "RIGHT")
	wallE.Move(board, "REPORT")

	assert.Equal(t, "1,5,NORTH", wallE.GetReport())
}

//test a robot that is not on the board can only execute command PLACE and REPORT
func TestStuckRobotMoveAndReport(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 1,5,NORTH")
	wallE.Move(board, "RIGHT")
	wallE.Move(board, "MOVE")
	wallE.Move(board, "RIGHT")
	wallE.Move(board, "MOVE")
	wallE.Move(board, "REPORT")

	assert.Equal(t, "1,5,NORTH", wallE.GetReport())
}

//a robot that is not on the board can only execute command PLACE and REPORT
func TestRobotMoveAndReport(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 2,2,NORTH")
	wallE.Move(board, "RIGHT")
	wallE.Move(board, "MOVE")
	wallE.Move(board, "RIGHT")
	wallE.Move(board, "MOVE")
	wallE.Move(board, "REPORT")

	assert.Equal(t, "3,1,SOUTH", wallE.GetReport())
}

//PROBLEM.md Example A
func TestRobotExampleA(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "MOVE")
	wallE.Move(board, "REPORT")

	assert.Equal(t, "0,1,NORTH", wallE.GetReport())
}

//PROBLEM.md Example B
func TestRobotExampleB(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "LEFT")
	wallE.Move(board, "REPORT")

	assert.Equal(t, "0,0,WEST", wallE.GetReport())
}

//PROBLEM.md Example C
func TestRobotExampleC(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 1,2,EAST")
	wallE.Move(board, "MOVE")
	wallE.Move(board, "MOVE")
	wallE.Move(board, "LEFT")
	wallE.Move(board, "MOVE")
	wallE.Move(board, "REPORT")

	assert.Equal(t, "3,3,NORTH", wallE.GetReport())
}
