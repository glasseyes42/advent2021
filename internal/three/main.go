package three

import (
	"fmt"
	"strconv"
	"strings"
)

type BitPosition struct {
	totalCount float64
	ones       float64
}

func (bp *BitPosition) AddVal(val int) {
	bp.totalCount++
	if val > 0 {
		bp.ones++
	}
}

func (bp *BitPosition) PopularValue() int {
	if bp.ones >= (bp.totalCount / 2) {
		return 1
	} else {
		return 0
	}
}

func (bp *BitPosition) UnPopularValue() int {
	if bp.ones >= (bp.totalCount / 2) {
		return 0
	} else {
		return 1
	}
}

func CalculateConsumption(lines []string) int64 {
	commonBits := make([]BitPosition, len(lines[0]))

	for _, line := range lines {
		for idx, bit := range strings.Split(line, "") {
			val, err := strconv.Atoi(bit)
			if err != nil {
				panic(err)
			}
			commonBits[idx].AddVal(val)
		}
	}

	gammaBits := make([]string, len(commonBits))
	epsilonBits := make([]string, len(commonBits))
	for idx, cb := range commonBits {
		gammaBits[idx] = fmt.Sprintf("%d", cb.PopularValue())
		epsilonBits[idx] = fmt.Sprintf("%d", cb.UnPopularValue())
	}

	gamma, err := strconv.ParseInt(strings.Join(gammaBits, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(strings.Join(epsilonBits, ""), 2, 64)
	if err != nil {
		panic(err)
	}

	return gamma * epsilon
}

func getSupportRating(lines []string, idx int, popular bool) string {
	if len(lines) == 1 {
		return lines[0]
	}

	commonBit := BitPosition{}
	for _, line := range lines {
		val, err := strconv.Atoi(string(line[idx]))
		if err != nil {
			panic(err)
		}
		commonBit.AddVal(val)
	}

	n := 0
	for _, line := range lines {
		if popular {
			if string(line[idx]) == fmt.Sprintf("%v", commonBit.PopularValue()) {
				lines[n] = line
				n++
			}
		} else {
			if string(line[idx]) == fmt.Sprintf("%v", commonBit.UnPopularValue()) {
				lines[n] = line
				n++
			}
		}
	}
	lines = lines[:n]

	return getSupportRating(lines, idx+1, popular)
}

func CalculateSupport(lines []string) int64 {
	oxLines := make([]string, len(lines))
	co2Lines := make([]string, len(lines))
	copy(oxLines, lines)
	copy(co2Lines, lines)
	oxygenBits := getSupportRating(oxLines, 0, true)
	co2Bits := getSupportRating(co2Lines, 0, false)

	oxygen, err := strconv.ParseInt(oxygenBits, 2, 64)
	if err != nil {
		panic(err)
	}
	co2, err := strconv.ParseInt(co2Bits, 2, 64)
	if err != nil {
		panic(err)
	}

	return oxygen * co2
}
