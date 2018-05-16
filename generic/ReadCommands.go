package toy_robot_generic

import (
	"os"
	"bufio"
)

//read commands from specified file path
func ReadCommands(filePath string) (list []string, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return
}
