package main

import (
	"fmt"
	"os"
)

//TODO: add description in README.md
//TODO: research on suitable terminal based graphic

func main() {
	//initialize board size
	var board Board
	board.Height = 5
	board.Width = 5

	//initialize variable
	var wallE Robot
	var commands []string

	//check if file path is provided
	if len(os.Args) == 2 {
		//read commands from specified file path
		commands, _ = ReadCommands(os.Args[1])
	}

	if len(commands) == 0 {
		fmt.Println("Commands not found. Switching to User Input Mode.")
		//initialize first position if PLACE command entered
		wallE.Init(board, EnterCommand())
		for {
			wallE.Move(board, EnterCommand())
		}
	} else {
		fmt.Println("Commands found. Processing commands...")
		var startCommand string
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
}

//check if err, panic err message
func check(err error) {
	if err != nil {
		panic(err)
	}
}
