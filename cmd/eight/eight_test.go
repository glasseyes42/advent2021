package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/glasseyes42/advent2021/internal/eight"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestEight(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("count all 1,4,7,8", func() {
			result := eight.CountSimple(testData)
			So(result, ShouldEqual, 26)
		})
	})

	Convey("part 2", t, func() {
		Convey("work out patterns", func() {
			result := eight.ComputePatterns(strings.Split("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab", " "))
			So(result, ShouldResemble, map[string]int{
				"abcdefg": 8,
				"bcdef":   5,
				"acdfg":   2,
				"abcdf":   3,
				"abd":     7,
				"abcdef":  9,
				"bcdefg":  6,
				"abef":    4,
				"abcdeg":  0,
				"ab":      1,
			})
		})

		Convey("count all the values having worked out the code", func() {
			result := eight.CountAll(testData)
			So(result, ShouldEqual, 61229)
		})
	})
}
