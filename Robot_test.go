package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDefaultRobotCoordinate(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 1,2,NORTHs")

	assert.Equal(t, "0,0,NORTH", wallE.GetReport())
}

func TestInitRobotCoordinate(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 1,2,NORTH")

	assert.Equal(t, "1,2,NORTH", wallE.GetReport())
}

func TestRobotTurnLeft(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "LEFT")

	assert.Equal(t, "0,0,SOUTH", wallE.GetReport())
}

func TestRobotTurnRight(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "RIGHT")

	assert.Equal(t, "0,0,EAST", wallE.GetReport())
}

func TestRobotTurnLeftTwice(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "LEFT")
	wallE.Move(board, "LEFT")

	assert.Equal(t, "0,0,WEST", wallE.GetReport())
}

func TestRobotTurnRightTwice(t *testing.T) {
	var board Board
	board.Height = 5
	board.Width = 5

	var wallE Robot
	wallE.Init(board, "PLACE 0,0,NORTH")
	wallE.Move(board, "RIGHT")
	wallE.Move(board, "RIGHT")

	assert.Equal(t, "0,0,WEST", wallE.GetReport())
}

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

	assert.Equal(t, "1,5,SOUTH", wallE.GetReport())
}