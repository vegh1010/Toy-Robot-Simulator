
Brainstorming:-
- create two dimension board 5 x 5
    - x, y coordinates
    (4,0)(4,1)(4,2)(4,3)(4,4)
    (3,0)(3,1)(3,2)(3,3)(3,4)
    (2,0)(2,1)(2,2)(2,3)(2,4)
    (1,0)(1,1)(1,2)(1,3)(1,4)
    (0,0)(0,1)(0,2)(0,3)(0,4)

- create toy robot
    - store x, y location and face position
    - rotate LEFT and RIGHT based on NORTH, SOUTH, EAST or WEST.

          NORTH
            |
    WEST ---|--- EAST
            |
          SOUTH

    - MOVE command move forward based on face position
        - validate move based on board
        - ignore all commands that cause toy robot to fail
        - robot that is not on the board will ignore commands except REPORT
    - REPORT command output x, y coordinates and face direction
    - if PLACE wasn't define during the start of the simulation, toy robot will
    be place at (0,0) facing NORTH
    - ignore any invalid commands
    - commands not case sensitive

- create unit testing
- read commands from file
- get input from user