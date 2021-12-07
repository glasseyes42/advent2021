package main

import (
	_ "embed"
	"testing"

	"github.com/glasseyes42/advent2021/internal/five"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestSix(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("process straight lines", func() {
			result := five.CountPointsWithOverlap(2, five.FilterStraight(five.ParseFile(testData)))
			So(result, ShouldEqual, 5)
		})
	})

	Convey("part 2", t, func() {
		Convey("process all lines", func() {
			result := five.CountPointsWithOverlap(2, five.ParseFile(testData))
			So(result, ShouldEqual, 12)
		})
	})
}
