package mars

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// defines the properties needed for a rover to navigate mars
type mars struct {
	maxX, maxY int
}

// creates a new mars object
func newMars(x, y int) *mars {
	return &mars{
		maxX: x,
		maxY: y,
	}
}

// contains the final positions after controlling a rover
type Result struct {
	X, Y      int
	Direction Direction
}

// takes a set of instructions to moves rovers around mars, then returns their final positions
func Start(rawInstructions string) ([]Result, error) {
	// each section of instructions is on its own line
	instructions := strings.Split(rawInstructions, "\n")
	if len(instructions) < 3 {
		return nil, errors.New("atleast 3 instructions must be given")
	}

	// the first line of the instructions defines how large mars is
	marsX, marsY, err := parseMarsSize(instructions[0])
	if err != nil {
		return nil, fmt.Errorf("parsing mars size instruction: %w", err)
	}

	mars := newMars(marsX, marsY)

	// a rover requires two instructions - it's initial position and the instructions to control it
	instructionPairs := splitInstructionsIntoPairs(instructions)

	results := make([]Result, 0, len(instructionPairs))

	for _, instructions := range instructionPairs {
		if len(instructions) != 2 {
			return nil, errors.New("invalid rover instructions, must be 2 for each rover")
		}

		// the first line of a rovers instructions is it's position
		roverX, roverY, roverDirection, err := parseRoverPositionInstruction(instructions[0])
		if err != nil {
			return nil, fmt.Errorf("parsing rover position instruction: %w", err)
		}

		rover := mars.newRover(roverX, roverY, roverDirection)

		// the second line of a rovers instructions says how to control it
		if err := rover.instruct(instructions[1]); err != nil {
			return nil, fmt.Errorf("instructing the rover: %w", err)
		}

		results = append(results, Result{
			X:         rover.x,
			Y:         rover.y,
			Direction: rover.direction,
		})
	}

	return results, nil
}

// get the X and Y coordinates for the maximum size of mars given an input of "X Y"
func parseMarsSize(instruction string) (int, int, error) {
	x, err := strconv.Atoi(instruction[0:1])
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(instruction[2:3])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

// rovers require two instructions so we group them into pairs
func splitInstructionsIntoPairs(instructions []string) [][]string {
	output := make([][]string, 0)
	roverInstructions := make([]string, 0)

	for i, instruction := range instructions {
		if i == 0 {
			// this is the mars size instruction
			continue
		}

		if len(roverInstructions) == 2 {
			output = append(output, roverInstructions)
			roverInstructions = make([]string, 0)
		}

		roverInstructions = append(roverInstructions, instruction)
	}

	output = append(output, roverInstructions)
	return output
}

// get the X, Y coordinates as well as the direction for where the rover should start
// e.g. 3 1 N -> x: 3, y: 1 and the direction is North
func parseRoverPositionInstruction(instruction string) (x, y int, direction Direction, err error) {
	x, err = strconv.Atoi(instruction[0:1])
	if err != nil {
		return
	}

	y, err = strconv.Atoi(instruction[2:3])
	if err != nil {
		return
	}

	direction, err = GetDirectionFromSymbol(instruction[4:5])
	if err != nil {
		return
	}

	return
}
