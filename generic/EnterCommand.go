// package toy_robot_generic is a generic package that contains support functions.
package toy_robot_generic

import (
	"os"
	"bufio"
	"fmt"
)

//get command from console input
func EnterCommand() (input string) {
	fmt.Print("Enter Command: ")

	//initialize scanner
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		//get input text entered
		input = scanner.Text()
	}
	return
}
