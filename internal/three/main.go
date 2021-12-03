package three

import (
	"fmt"
	"strconv"
	"strings"
)

type bitPosition struct {
	totalCount int
	ones       int
}

func (bp *bitPosition) AddVal(val int) {
	bp.totalCount++
	if val > 0 {
		bp.ones++
	}
}

func (bp *bitPosition) PopularValue() int {
	if bp.ones > (bp.totalCount / 2) {
		return 1
	} else {
		return 0
	}
}

func (bp *bitPosition) UnPopularValue() int {
	if bp.ones > (bp.totalCount / 2) {
		return 0
	} else {
		return 1
	}
}

func CalculateConsumption(lines []string) int64 {
	commonBits := make([]bitPosition, len(lines[0]))

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
