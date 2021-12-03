package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/glasseyes42/advent2021/internal/two"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/2.1-sample.txt
var testData string

func TestTwo(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("counts increases", func() {
			result := two.ProcessInstructions(strings.Split(testData, "\n"))
			So(result.Depth*result.Horizonal, ShouldEqual, 150)
		})
	})
}
