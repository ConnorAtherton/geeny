package geeny

import (
	"testing"
)

var geenyTests = []struct {
	test             string
	input            []string
	expectedCommands commands
	expectedOptions  options
}{
	{
		"empty args",
		[]string{},
		[]string{},
		map[string]interface{}{},
	},
	{
		"only commands",
		[]string{"tool", "up"},
		[]string{"tool", "up"},
		map[string]interface{}{},
	},
	{
		"single flags",
		[]string{"tool", "-fGb"},
		[]string{"tool"},
		map[string]interface{}{
			"f": true,
			"G": true,
			"b": true,
		},
	},
	{
		"single flags with number",
		[]string{"tool", "-c3", "-n10"},
		[]string{"tool"},
		map[string]interface{}{
			"c": 3,
			"n": 10,
		},
	},
}

func TestGeeny(t *testing.T) {
	for _, tc := range geenyTests {
		t.Run(tc.test, func(t *testing.T) {
			res, err := Parse(tc.input)

			if err != nil {
				t.Error("Unexpected error")
			}

			if !commandsAreSame(res.Commands, tc.expectedCommands) {
				t.Errorf("Unexpected command list, have %v, expected %v", res.Commands, tc.expectedCommands)
			}

			if !optionsAreSame(res.Options, tc.expectedOptions) {
				t.Errorf("Unexpected options, have %v, expected %v", res.Options, tc.expectedOptions)
			}
		})
	}
}

func commandsAreSame(a commands, b commands) bool {
	set := map[string]bool{}

	if len(a) != len(b) {
		return false
	}

	for _, val := range a {
		set[val] = true
	}

	for _, val := range b {
		if _, ok := set[val]; !ok {
			return false
		}
	}

	return true
}

func optionsAreSame(a options, b options) bool {
	if len(a) != len(b) {
		return false
	}

	for key, val := range a {
		if res, ok := b[key]; !ok || (val != res) {
			return false
		}
	}

	return true
}
