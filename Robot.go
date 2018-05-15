package main

//Variables to manage toy robot's movement
var (
	NORTH = Coordinate{1, 0}
	EAST  = Coordinate{0, 1}
	WEST  = Coordinate{0, -1}
	SOUTH = Coordinate{-1, 0}

	Facing = map[int]string{
		1: "NORTH",
		2: "EAST",
		3: "WEST",
		4: "SOUTH",
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
func (self *Robot) Init(command string) {
	//TODO: PLACE, if not provide, call move function. If PLACE command but error, default (0,0) NORTH

}

//Control robot's movement based on command provided
func (self *Robot) Move(board Board, command string) {
	//TODO: PLACE, MOVE, LEFT, RIGHT, REPORT
}
