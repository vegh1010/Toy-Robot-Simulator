# Toy-Robot-Simulator
A toy robot simulator that accepts commands from a text file and move on a 5 x 5 board.

### Tools Used:
   1) go version 1.8.3
   2) bash

### Assumptions:
   1) Predefine commands can be execute from a file by including file path as an argument when executing this program. ie: ./program ./commands.txt
   2) If file not found or no commands provided, commands are switch to User Input Mode where users need to manually input commands.
   3) Commands are not case sensitive. The program will convert any commands to uppercase.
   4) If the toy robot is not on the board, it can only accept PLACE command.
   5) If PLACE command is not specified during the start, the toy robot will be place by default 0,0,NORTH.
   6) If a PLACE command that is provided during the start is invalid for some reason, the toy robot will be place on the default coordinates.
   7) The only commands that the toy robot will accept are PLACE, MOVE, RIGHT, LEFT and REPORT. The rest will be ignore.
   8) The toy robot will ignore all commands that will cause it to fall.

### Commands Available:
   1) PLACE X,Y,F
        - X: 0-4
        - Y: 0-4
        - F: NORTH, SOUTH, EAST, WEST
   2) MOVE
        - will move forward based on facing direction
   3) RIGHT
        - will turn right
   4) LEFT
        - will turn left
   5) REPORT
        - will output current coordinate and facing direction

### Step to Step:
   1) build the program into a binary / exe file
   2) run the build.
      - if you want to run with predefined commands. example: ./build commands.txt
      - if you want to manually input commands. example: ./build
   3) done. WallE will now try to move and report.

    Note: board size can be modified by changing the values in board.json.
    The program will try to read the data from board.json and used it if the json file and data defined is valid.
    To modify the data, go into board.json and change the value of Width or Height.
    Any invalid values will be ignored and the default values will be set to the board.
    Default values are 5 x 5.

### Code Structure Design:
   1) For this project, 3 packages have been created.
      - generic contains any generic functions to support toy robot simulator
      - lib contains data structs, initialize data and reference functions to support toy
      robot simulator's features
      - gui contains function to output basic gui on console
   2) In main.go contains basic conditions to run the toy robot simulator and output the gui.
   3) From my perspective, this structure design will allow for a more manageable, reusable and readable of code.
   4) Further explanation has been included in the code.