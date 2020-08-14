package mars_test

import (
	"github.com/ducc/mars-rover"
	"github.com/go-test/deep"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []mars.Result
		expectedErr string
	}{
		{
			name:  "challenge example",
			input: "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM",
			expected: []mars.Result{
				{
					X:         1,
					Y:         3,
					Direction: mars.North,
				},
				{
					X:         5,
					Y:         1,
					Direction: mars.East,
				},
			},
		},
		{
			name:  "drive one rover around the perimeter",
			input: "2 2\n0 0 E\nMMLMMLMMLM",
			expected: []mars.Result{
				{
					X:         0,
					Y:         1,
					Direction: mars.South,
				},
			},
		},
		{
			name:  "the rover cannot drive off the edge of mars",
			input: "2 2\n0 0 N\nMMMMMMMMMMMMM",
			expected: []mars.Result{
				{
					X:         0,
					Y:         2,
					Direction: mars.North,
				},
			},
		},
		{
			name:  "every tile is a rover",
			input: "1 1\n0 0 E\nL\n1 0 N\nL\n0 1 W\nL\n1 1 S\nL",
			expected: []mars.Result{
				{
					X:         0,
					Y:         0,
					Direction: mars.North,
				},
				{
					X:         1,
					Y:         0,
					Direction: mars.West,
				},
				{
					X:         0,
					Y:         1,
					Direction: mars.South,
				},
				{
					X:         1,
					Y:         1,
					Direction: mars.East,
				},
			},
		},
		{
			name:        "invalid direction",
			input:       "1 1\n0 0 G\nL",
			expectedErr: "parsing rover position instruction: invalid direction: 'G'",
		},
		{
			name:        "no instructions",
			input:       "",
			expectedErr: "atleast 3 instructions must be given",
		},
		{
			name:        "invalid mars x",
			input:       "A 0\n0 0 N\nL",
			expectedErr: "parsing mars size instruction: strconv.Atoi: parsing \"A\": invalid syntax",
		},
		{
			name:        "invalid mars y",
			input:       "0 B\n0 0 N\nL",
			expectedErr: "parsing mars size instruction: strconv.Atoi: parsing \"B\": invalid syntax",
		},
		{
			name:        "uneven rover instructions",
			input:       "1 1\n0 0 N\nL\n1 1 N",
			expectedErr: "invalid rover instructions, must be 2 for each rover",
		},
		{
			name:        "invalid rover x",
			input:       "1 1\nA 0 N\nL",
			expectedErr: "parsing rover position instruction: strconv.Atoi: parsing \"A\": invalid syntax",
		},
		{
			name:        "invalid rover y",
			input:       "1 1\n0 B N\nL",
			expectedErr: "parsing rover position instruction: strconv.Atoi: parsing \"B\": invalid syntax",
		},
		{
			name:        "invalid rover instruction",
			input:       "1 1\n0 0 N\nG",
			expectedErr: "instructing the rover: handling rover instructions: invalid instruction",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results, err := mars.Start(test.input)

			if err != nil {
				if test.expectedErr == "" || test.expectedErr != err.Error() {
					t.Fatalf("unexpected err: %s", err)
				}
				return
			}
			if test.expectedErr != "" {
				t.Fatal("expected error")
			}

			if diff := deep.Equal(test.expected, results); diff != nil {
				t.Fatal(strings.Join(diff, "\n"))
			}
		})
	}
}
