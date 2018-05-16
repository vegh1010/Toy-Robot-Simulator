# Toy-Robot-Simulator
A toy robot simulator that accepts commands from a text file and move on a 5 x 5 board.

Tools Used:
   1) go version 1.8.3
   2) bash

Assumptions:
   1) file path needs to be provided as an argument when executing this program. ie: ./program <file path>
   2) If PLACE command is not specified during the start, the toy robot will be place by default 0,0,NORTH.
   3) If a PLACE command is provided during the start by invalid for some reason, the toy robot will be place on
   the default coordinates.
   4) The only commands that the toy robot will accept are PLACE, MOVE, RIGHT, LEFT and REPORT. The rest will be
   ignore.
   5) The toy robot will ignore all commands that will cause it to fall.
   6) If the toy robot is not on the board can only accept PLACE command.
   7) If file not found or no commands provided, commands are switch to User Input Mode where users need
   to manual input commands.

Commands Available:
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

Step to Step:
    1) build the program into a binary file / exe file
    2) run the build. ie: ./build <file path optional>
    3) done. The toy robot will move based on the commands provided in the text file or from console input.

Code Structure Design:
    1) TODO: once code is workable, rearrange code for reusability and readability