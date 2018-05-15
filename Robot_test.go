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