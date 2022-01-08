package main

import (
	_ "embed"
	"testing"

	"github.com/glasseyes42/advent2021/internal/eleven"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestTen(t *testing.T) {
	Convey("part 1", t, func() {
		result := eleven.Iterate(eleven.BuildGrid(testData), 10)
		So(result, ShouldEqual, 204)

		result = eleven.Iterate(eleven.BuildGrid(testData), 100)
		So(result, ShouldEqual, 1656)
	})

	Convey("part 2", t, func() {
		result := eleven.IterateTillAllFlash(eleven.BuildGrid(testData))
		So(result, ShouldEqual, 195)
	})
}
