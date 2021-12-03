package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/glasseyes42/advent2021/internal/one"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/1.1-sample.txt
var testData string

func TestOne(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("counts increases", func() {
			result := one.CountIncreases(strings.Split(testData, "\n"))
			So(result, ShouldEqual, 7)
		})
	})

	Convey("part 2", t, func() {
		Convey("counts increases with a sliding window", func() {
			result := one.CountIncreasesWindowed(strings.Split(testData, "\n"))
			So(result, ShouldEqual, 5)
		})
	})
}
