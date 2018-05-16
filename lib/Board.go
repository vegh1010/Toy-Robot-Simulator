// package toy_robot_lib is a library package that contains data struct, reference functions
// and unit test.
package toy_robot_lib

//Board's struct to define size of the board
type Board struct {
	Height int
	Width  int
}

func (self *Board) Init() {
	self.Height = 5
	self.Width = 5
}