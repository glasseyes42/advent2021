package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/glasseyes42/advent2021/internal/three"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/3.1-sample.txt
var testData string

func TestThree(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("calculate consumption", func() {
			result := three.CalculateConsumption(strings.Split(testData, "\n"))
			So(result, ShouldEqual, 198)
		})
	})
}
