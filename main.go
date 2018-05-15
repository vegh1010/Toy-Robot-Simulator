package main

import (
	"fmt"
)

//TODO: add description in README.md
//TODO: add comments on unit testing
//TODO: set file path as argument when run program (./binary <file path>)
//TODO: option if file path no specified or not found change to User Input Mode
//TODO: add more unit testing based on PROBLEM.md Example Input and Output
//TODO: research on suitable terminal based graphic
func main() {
	//read commands from specified file path
	commands, err := ReadCommands("./commands.txt")
	check(err)
	for _, command := range commands {
		fmt.Println(command)
	}

	//initialize board size
	var board Board
	board.Height = 5
	board.Width = 5

	var startCommand string

	//initialize robot wallE
	var wallE Robot
	if len(commands) > 0 {
		startCommand = commands[0]
		commands = commands[1:]
	}
	//initialize first position if specified else default
	wallE.Init(board, startCommand)

	//run commands from file
	for _, command := range commands {
		//execute move function
		wallE.Move(board, command)
	}
}

//check if err, panic err message
func check(err error) {
	if err != nil {
		panic(err)
	}
}
