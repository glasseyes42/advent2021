package main

import (
	_ "embed"
	"testing"

	"github.com/glasseyes42/advent2021/internal/seven"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestSeven(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("calc fuel", func() {
			result := seven.CalcFuel(seven.ParseFile(testData))
			So(result, ShouldEqual, 37)
		})
	})
}
