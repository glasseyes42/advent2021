package eight

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func CountSimple(file string) int {
	lines := strings.Split(strings.TrimSpace(file), "\n")

	total := 0
	for _, line := range lines {
		hyphenSplit := strings.Split(strings.TrimSpace(line), "|")
		rhs := hyphenSplit[1]
		segments := strings.Split(strings.TrimSpace(rhs), " ")

		for _, segment := range segments {
			length := len(segment)

			if length == 4 || length == 3 || length == 7 || length == 2 {
				total++
			}
		}
	}

	return total
}

func findStringWithLength(list []string, length int) []string {
	matches := []string{}

	for _, s := range list {
		if len(s) == length {
			matches = append(matches, s)
		}
	}
	return matches
}

func normalize(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)

	return strings.Join(split, "")
}

func combine(list []string) string {
	components := map[string]bool{}

	for _, s := range list {
		for _, c := range strings.Split(s, "") {
			components[c] = true
		}
	}

	toReturn := ""
	for key := range components {
		toReturn = fmt.Sprintf("%s%s", toReturn, key)
	}

	return toReturn
}

func checkIfOneContainsOther(first, second string) bool {
	firstList := strings.Split(first, "")
	secondList := strings.Split(second, "")
	sort.Strings(firstList)
	sort.Strings(secondList)

	contains := true
	for _, component := range secondList {
		contains = contains && strings.Contains(strings.Join(firstList, ""), component)
	}
	return contains
}

func ComputePatterns(line []string) map[string]int {
	patterns := map[string]int{}
	one := normalize(findStringWithLength(line, 2)[0])
	patterns[one] = 1
	four := normalize(findStringWithLength(line, 4)[0])
	patterns[four] = 4
	seven := normalize(findStringWithLength(line, 3)[0])
	patterns[seven] = 7
	patterns[normalize(findStringWithLength(line, 7)[0])] = 8

	nine := ""
	zero := ""
	zeroNineAndSix := findStringWithLength(line, 6)
	for len(zeroNineAndSix) > 0 {
		for idx, option := range zeroNineAndSix {
			if checkIfOneContainsOther(option, combine([]string{four, seven})) {
				nine = normalize(option)
				patterns[nine] = 9
				zeroNineAndSix = append(zeroNineAndSix[:idx], zeroNineAndSix[idx+1:]...)
				break
			}
			if nine != "" && checkIfOneContainsOther(option, seven) {
				zero = normalize(option)
				patterns[zero] = 0
				zeroNineAndSix = append(zeroNineAndSix[:idx], zeroNineAndSix[idx+1:]...)
				break
			}
			if nine != "" && zero != "" {
				patterns[normalize(zeroNineAndSix[0])] = 6
				zeroNineAndSix = append(zeroNineAndSix[:idx], zeroNineAndSix[idx+1:]...)
				break
			}
		}
	}

	twoThreeFive := findStringWithLength(line, 5)
	three := ""
	five := ""
	for len(twoThreeFive) > 0 {
		for idx, option := range twoThreeFive {
			if checkIfOneContainsOther(option, one) {
				three = normalize(option)
				patterns[three] = 3
				twoThreeFive = append(twoThreeFive[:idx], twoThreeFive[idx+1:]...)
				break
			}
			if three != "" && checkIfOneContainsOther(nine, option) {
				five = normalize(option)
				patterns[five] = 5
				twoThreeFive = append(twoThreeFive[:idx], twoThreeFive[idx+1:]...)
				break
			}
			if three != "" && five != "" {
				patterns[normalize(twoThreeFive[0])] = 2
				twoThreeFive = append(twoThreeFive[:idx], twoThreeFive[idx+1:]...)
				break
			}
		}
	}

	return patterns
}

func CountAll(file string) int {
	lines := strings.Split(strings.TrimSpace(file), "\n")

	total := 0
	for _, line := range lines {
		hyphenSplit := strings.Split(strings.TrimSpace(line), "|")
		lhs := hyphenSplit[0]
		rhs := hyphenSplit[1]
		segments := strings.Split(strings.TrimSpace(rhs), " ")
		patterns := strings.Split(strings.TrimSpace(lhs), " ")
		conversionMap := ComputePatterns(patterns)

		strRep := ""
		for _, segment := range segments {
			strRep = fmt.Sprintf("%s%d", strRep, conversionMap[normalize(segment)])
		}

		val, err := strconv.Atoi(strRep)
		if err != nil {
			panic(err)
		}
		total += val
	}

	return total
}
