package main

import (
	"fmt"
	"os"
	"github.com/vegh1010/Toy-Robot-Simulator/lib"
	"github.com/vegh1010/Toy-Robot-Simulator/generic"
	"github.com/vegh1010/Toy-Robot-Simulator/gui"
)

func main() {
	//define reference functions
	var GUI = toy_robot_gui.Output
	var EnterCommand = toy_robot_generic.EnterCommand

	//initialize variable
	var board toy_robot_lib.Board
	board.Init()
	var wallE toy_robot_lib.Robot
	var commands []string

	//check if file path is provided
	if len(os.Args) == 2 {
		//read commands from specified file path
		commands, _ = toy_robot_generic.ReadCommands(os.Args[1])
	}

	//if no commands, User Input Mode
	if len(commands) == 0 {
		fmt.Println("Commands not found. Switching to User Input Mode.")

		//get input command
		var command = EnterCommand()
		//initialize first position if PLACE command entered
		wallE.Init(board, command)
		GUI(board, wallE)

		for {
			//get input command
			command = EnterCommand()
			wallE.Move(board, command)
			GUI(board, wallE)
		}
	} else {
		fmt.Println("Commands found. Processing commands...")
		var startCommand string
		if len(commands) > 0 {
			startCommand = commands[0]
			commands = commands[1:]
		}

		fmt.Println("Command:", startCommand)
		//initialize first position if specified else default
		wallE.Init(board, startCommand)
		GUI(board, wallE)

		//run commands from file
		for _, command := range commands {
			fmt.Println("Command:", command)
			//execute move function
			wallE.Move(board, command)
			GUI(board, wallE)
		}
	}
}
