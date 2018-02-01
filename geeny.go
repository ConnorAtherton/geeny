package geeny

import (
	"regexp"
	"strconv"
	"unicode"
)

type commands []string
type options map[string]interface{}

type Args struct {
	Commands commands
	Options  options
}

var (
	command            = regexp.MustCompile(`^[A-Za-z]`)
	singleDash         = regexp.MustCompile(`^-([A-Za-z])+`)
	doubleDash         = regexp.MustCompile(`^--([A-Za-z]*-?[A-Za-z]*)=?([A-Za-z0-9]*)`)
	doubleDashNegation = regexp.MustCompile(`^--no-([A-Za-z]*-?[A-Za-z]*)`)
)

func Parse(args []string) (*Args, error) {
	if len(args) == 0 {
		return &Args{}, nil
	}

	cmds := []string{}
	opts := map[string]interface{}{}
	shouldSkip := false

	for _, val := range args {
		if shouldSkip {
			shouldSkip = false
			continue
		}

		if command.MatchString(val) {
			cmds = append(cmds, val)
		}

		if singleDash.MatchString(val) {
			chars := []rune(val[1:])
			length := len(chars)

			for i, letter := range chars {
				if i+1 < length && unicode.IsDigit(chars[i+1]) {
					opts[string(letter)] = int(chars[i+1]) - '0'
					shouldSkip = true
					break
				} else {
					opts[string(letter)] = true
				}
			}

			continue
		}

		if doubleDashNegation.MatchString(val) {
			match := doubleDashNegation.FindStringSubmatch(val)

			opts[match[1]] = false

			continue
		}

		if doubleDash.MatchString(val) {
			match := doubleDash.FindStringSubmatch(val)

			if len(match) == 3 && match[2] == "" {
				opts[match[1]] = true
			} else {
				if intVal, err := strconv.Atoi(match[2]); err == nil {
					opts[match[1]] = intVal
				} else {
					opts[match[1]] = match[2]
				}
			}
		}
	}

	return &Args{Commands: cmds, Options: opts}, nil
}
