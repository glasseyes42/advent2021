package main

import (
	_ "embed"
	"testing"

	"github.com/glasseyes42/advent2021/internal/six"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestSix(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("process a day", func() {
			result := six.ProcessDay(six.ParseFile(testData))
			So(result, ShouldResemble, []int{1, 1, 2, 1, 0, 0, 0, 0, 0}) // 2,3,2,0,1
		})

		Convey("80 days", func() {
			result := six.ProcessDays(80, six.ParseFile(testData))
			So(result, ShouldEqual, 5934)
		})
	})

	Convey("part 2", t, func() {
		Convey("256 days", func() {
			result := six.ProcessDays(256, six.ParseFile(testData))
			So(result, ShouldEqual, 26984457539)
		})
	})

}
