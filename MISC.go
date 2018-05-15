package main

import (
	"os"
	"bufio"
)

type Coordinate struct {
	X int
	Y int
}

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

