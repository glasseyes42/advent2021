package main

import (
	_ "embed"
	"testing"

	"github.com/glasseyes42/advent2021/internal/eight"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestSix(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("count all 1,4,7,8", func() {
			result := eight.CountSimple(testData)
			So(result, ShouldEqual, 26)
		})
	})

}
