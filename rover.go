package mars

import (
	"fmt"
	"strings"
)

// an enum for each valid instruction for the rover
type RoverInstruction string

const (
	TurnLeft    RoverInstruction = "L"
	TurnRight   RoverInstruction = "R"
	MoveForward RoverInstruction = "M"
)

// contains the data needed to control the rover
type rover struct {
	mars      *mars
	x, y      int
	direction Direction
}

// creates a new rover
func (m *mars) newRover(x, y int, direction Direction) *rover {
	r := &rover{
		mars:      m,
		direction: direction,
		x:         x,
		y:         y,
	}
	return r
}

// takes a set of movement instructions e.g. LMMRRM
func (r *rover) instruct(instructions string) error {
	for _, instruction := range splitRoverInstructions(instructions) {
		if err := r.handleInstruction(RoverInstruction(instruction)); err != nil {
			return fmt.Errorf("handling rover instructions: %w", err)
		}
	}

	return nil
}

// takes a single movement instruction e.g. L to turn left or M to move forward
func (r *rover) handleInstruction(instruction RoverInstruction) error {
	switch instruction {
	case TurnLeft:
		r.direction = r.direction.Left()
	case TurnRight:
		r.direction = r.direction.Right()
	case MoveForward:
		r.moveForward()
	default:
		return fmt.Errorf("invalid instruction")
	}

	return nil
}

// moves the rover forward based on it's current direction
func (r *rover) moveForward() {
	switch r.direction {
	case North:
		r.y = min(r.y+1, r.mars.maxY)
	case East:
		r.x = min(r.x+1, r.mars.maxX)
	case South:
		r.y = max(r.y-1, 0)
	case West:
		r.x = max(r.x-1, 0)
	}
}

// splits a set of instructions into each individual instruction e.g. LMMR -> L, M, M, R
func splitRoverInstructions(instructions string) []string {
	output := make([]string, 0)
	for _, instruction := range strings.Split(instructions, "") {
		if instruction != "" {
			output = append(output, instruction)
		}
	}
	return output
}
