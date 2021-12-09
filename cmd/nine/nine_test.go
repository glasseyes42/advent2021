package main

import (
	_ "embed"
	"testing"

	"github.com/glasseyes42/advent2021/internal/nine"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestNine(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("count low levels", func() {
			result := nine.RiskOfLowLevels(testData)
			So(result, ShouldEqual, 15)
		})
	})

	Convey("part 2", t, func() {
		Convey("largest basin size", func() {
			result := nine.LargestBasins(testData)
			So(result, ShouldEqual, 1134)
		})
	})
}
