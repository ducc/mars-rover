# Mars Rover
This repository contains the code to control theoretical robotic rovers to explore Mars.

Rovers are placed on mars at certain coordinates facing a direction, and can be moved one tile forward at a time.
The rovers can turn by left or right relative to their current direction.

The rover also cannot drive past the area to explore or off the edge of mars as it limits itself to the input
coordinates.

I have written this code in Go because it allows logic to be expressed clearly, has an excellent unit testing framework
and compiles quickly which helps to make quick code changes and see their impact almost instantly.

## Example input/output
See the `mars_test.go` file for examples of possible inputs into the rovers, as well as 100% coverage of all
statements including invalid inputs and errors.

## Continuous integration
Pushes to this repository are tested using Github Actions and `go test`.

## Usage
1. `git clone git@github.com:ducc/mars-rover.git`
1. `cd mars-rover`
1. `go test`

## Troubleshooting
This project uses go modules which are still in beta - if you have dependency issues try adding the following to your bash/zsh profile:
```bash
export GO111MODULE=on
export GOFLAGS=-mod=vendor
```

Tested using go 1.14.4 on macOS 10.15.5

