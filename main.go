package main

import (
	"fmt"
	"os"
)

//TODO: add description in README.md
//TODO: option if file path no specified or not found change to User Input Mode
//TODO: research on suitable terminal based graphic

func main() {
	//check if file path is provided
	if len(os.Args) != 2 {
		//panic if not provided
		panic("File Path required")
	}

	//read commands from specified file path
	commands, err := ReadCommands(os.Args[1])
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
