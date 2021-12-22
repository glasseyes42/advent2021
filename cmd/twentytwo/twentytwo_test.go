package main

import (
	_ "embed"
	"testing"

	"github.com/glasseyes42/advent2021/internal/twentytwo"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestTwentyTwo(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("init", func() {
			result := twentytwo.Init(twentytwo.ParseFile(testData))
			So(result, ShouldEqual, 590784)
		})
	})
}
