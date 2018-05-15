package main

import (
	"fmt"
	"strings"
	"strconv"
)

//Variables to manage toy robot's movement
var (
	NORTH = Coordinate{1, 0}
	EAST  = Coordinate{0, 1}
	WEST  = Coordinate{0, -1}
	SOUTH = Coordinate{-1, 0}

	FACING = map[int]string{
		1: "NORTH",
		2: "EAST",
		3: "WEST",
		4: "SOUTH",
	}
	FACING_INVERT = map[string]int{
		"NORTH": 1,
		"EAST": 2,
		"WEST": 3,
		"SOUTH": 4,
	}
)

//Board's struct to define size of the board
type Board struct {
	Height int
	Width  int
}

//Toy Robot's struct where it's coordinate and commands will be processed
type Robot struct {
	X    int
	Y    int
	Face int
}

//Initialize robot's coordinate and face position if defined else
//default position and execute command provided
func (self *Robot) Init(board Board, command string) {
	command = strings.ToUpper(command)
	var defaultPlace = "PLACE 0,0,NORTH"

	//validate command
	if !self.Place(command) {
		self.Place(defaultPlace)
		self.Move(board, command)
	}
}

//Control robot's movement based on command provided
func (self *Robot) Move(board Board, command string) {
	//TODO: PLACE, MOVE, LEFT, RIGHT, REPORT
	command = strings.ToUpper(command)
	if strings.Contains(command, "PLACE") {
		self.Place(command)
	} else if command == "LEFT" {
		self.Face -= 1
		if self.Face < 1 {
			self.Face = 4
		}
	} else if command == "RIGHT" {
		self.Face += 1
		if self.Face > 4 {
			self.Face = 1
		}
	} else if command == "REPORT" {
		fmt.Println(self.GetReport())
	}
}

//manage robot's placement
func (self *Robot) Place(command string) (bool) {
	//check if PLACE is in command
	if strings.Contains(command, "PLACE") {
		commands := strings.Split(command, " ")
		if len(commands) == 2 {
			args := strings.Split(commands[1], ",")
			if len(args) == 3 {
				//validate arguments
				x, err :=strconv.Atoi(args[0])
				if err != nil {
					return false
				}
				y, err :=strconv.Atoi(args[1])
				if err != nil {
					return false
				}
				value, exist := FACING_INVERT[args[2]]
				if !exist {
					return false
				}
				//set robot's coordinate and face position
				self.X = x
				self.Y = y
				self.Face = value

				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
	return false
}

//get robot's coordinate and face position
func (self *Robot) GetReport() (string) {
	report := fmt.Sprint(self.X, ",", self.Y, ",", FACING[self.Face])

	return report
}
