// package toy_robot_lib is a library package that contains data struct, reference functions
// and unit test.
package toy_robot_lib

import (
	"io/ioutil"
	"encoding/json"
	"strconv"
)

//Board's struct to define size of the board
type Board struct {
	Height int
	Width  int
}

//initialize board
func (self *Board) Init() {
	//width and height from board.json is used if found and loaded
	width, height := self.load()

	//if either one is zero, set default height and width
	if width == 0 || height == 0 {
		width = 5
		height = 5
	}
	self.Height = height
	self.Width = width
}

//load board data from board.json
func (self *Board) load() (width int, height int) {
	Configs := make(map[string]string)

	//get data from file path
	data, err := ioutil.ReadFile("./board.json")
	if err != nil {
		return
	} else {
		//parse json string into object
		err = json.Unmarshal(data, &Configs)
		if err != nil {
			return
		}

		//convert string to int
		width, _ = strconv.Atoi(Configs["Width"])
		height, _ = strconv.Atoi(Configs["Height"])
	}
	return
}