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