package mars

import "fmt"

// an enum of all the possible directions a rover could turn
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// finds the direction to the left of the current direction e.g. South -> East
func (d Direction) Left() Direction {
	newDirection := int(d) - 1
	if newDirection < 0 {
		newDirection = 3
	}
	return Direction(newDirection)
}

// finds the direction to the right of the current direction e.g. East -> South
func (d Direction) Right() Direction {
	newDirection := int(d) + 1
	if newDirection > 3 {
		newDirection = 0
	}
	return Direction(newDirection)
}

// finds the Direction enum variant for a symbol or returns an error if the symbol is invalid
func GetDirectionFromSymbol(symbol string) (Direction, error) {
	switch symbol {
	case "N":
		return North, nil
	case "E":
		return East, nil
	case "S":
		return South, nil
	case "W":
		return West, nil
	default:
		return 0, fmt.Errorf("invalid direction: '%s'", symbol)
	}
}
