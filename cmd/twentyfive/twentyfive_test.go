package main

import (
	_ "embed"
	"testing"

	"github.com/glasseyes42/advent2021/internal/twentyfive"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestTwentyFive(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("map is linked correctly", func() {
			seed := twentyfive.ParseFile(testData)
			So(seed[0][len(seed[0])-1].Right, ShouldEqual, seed[0][0])
			So(seed[0][0].Below, ShouldEqual, seed[1][0])
			So(seed[len(seed)-1][0].Below, ShouldEqual, seed[0][0])
		})

		Convey("process till no movement", func() {
			result := twentyfive.IterateTillStopped(twentyfive.ParseFile(testData))
			So(result, ShouldEqual, 58)
		})
	})
}
