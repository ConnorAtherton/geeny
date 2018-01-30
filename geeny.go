package geeny

import (
	"regexp"
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
	singleDash         = regexp.MustCompile(`^-.+`)
	doubleDash         = regexp.MustCompile(`^--.+`)
	doubleDashNegation = regexp.MustCompile(`^--no-.+`)
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
			chars := []byte(val)
			length := len(chars)

			for i, letter := range chars[1:] {
				if i+1 < length && unicode.IsNumber(rune(chars[i+1])) {
					opts[string(letter)] = chars[i+1]
					shouldSkip = true
					break
				} else {
					opts[string(letter)] = true
				}
			}
		}
	}

	return &Args{Commands: cmds, Options: opts}, nil
}
