package uniq

import (
	"fmt"
	"strings"
)

type RunOptions struct {
	Count      bool
	Duplicates bool
	Unique     bool
	SkipFields int
	SkipChars  int
	IgnoreCase bool
}

func getModifiedLine(line string, options RunOptions) string {
	if options.IgnoreCase {
		line = strings.ToUpper(line)
	}

	if options.SkipChars > 0 {
		if len(line) >= options.SkipChars {
			line = line[options.SkipChars:]
		}
	}

	if options.SkipFields > 0 {
		splited := strings.Split(line, " ")
		if len(splited) > 1 {
			line = strings.Join(splited[options.SkipFields:], " ")
		}
	}

	return line
}

func Uniq(lines []string, options RunOptions) []string {
	linesMap := make(map[string][]int)
	linesPositions := make([]int, 0)

	for index, line := range lines {
		line = getModifiedLine(line, options)

		if _, has := linesMap[line]; !has {
			linesMap[line] = make([]int, 0)
			linesPositions = append(linesPositions, index)
		}

		linesMap[line] = append(linesMap[line], index)
	}

	resultLines := make([]string, 0)

	switch {
	case options.Count:
		for _, pos := range linesPositions {
			line := lines[pos]

			if count := len(linesMap[getModifiedLine(line, options)]); count > 0 {
				resultLines = append(resultLines, fmt.Sprint(count)+" "+line)
			}
		}
	case options.Duplicates:
		for _, pos := range linesPositions {
			if line := lines[pos]; len(linesMap[getModifiedLine(line, options)]) > 1 {
				resultLines = append(resultLines, line)
			}
		}
	case options.Unique:
		for _, pos := range linesPositions {
			if line := lines[pos]; len(linesMap[getModifiedLine(line, options)]) == 1 {
				resultLines = append(resultLines, line)
			}
		}
	default:
		for _, pos := range linesPositions {
			resultLines = append(resultLines, lines[pos])
		}
	}

	return resultLines
}
