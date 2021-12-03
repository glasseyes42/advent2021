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

//go:embed data/3.1.txt
var realData string

func TestThree(t *testing.T) {
	Convey("part 1", t, func() {
		Convey("calculate consumption", func() {
			result := three.CalculateConsumption(strings.Split(testData, "\n"))
			So(result, ShouldEqual, 198)
		})
	})

	Convey("part 2", t, func() {
		Convey("calculate support", func() {
			result := three.CalculateSupport(strings.Split(testData, "\n"))
			So(result, ShouldEqual, 230)
		})

		Convey("wrong answer", func() {
			result := three.CalculateSupport(strings.Split(realData, "\n"))
			So(result, ShouldNotEqual, 5128275)
		})
	})

	Convey("BitPosition popularity", t, func() {
		Convey("Determines 1 as popular", func() {
			bp := three.BitPosition{}
			bp.AddVal(1)
			bp.AddVal(1)
			bp.AddVal(0)

			So(bp.PopularValue(), ShouldEqual, 1)
		})

		Convey("Determines 1 as popular when equal", func() {
			bp := three.BitPosition{}
			bp.AddVal(1)
			bp.AddVal(1)
			bp.AddVal(0)
			bp.AddVal(0)

			So(bp.PopularValue(), ShouldEqual, 1)
		})

		Convey("Determines 0 as popular", func() {
			bp := three.BitPosition{}
			bp.AddVal(1)
			bp.AddVal(0)
			bp.AddVal(0)

			So(bp.PopularValue(), ShouldEqual, 0)
		})

		Convey("Determines 1 as unpopular", func() {
			bp := three.BitPosition{}
			bp.AddVal(1)
			bp.AddVal(0)
			bp.AddVal(0)

			So(bp.UnPopularValue(), ShouldEqual, 1)
		})

		Convey("Determines 0 as unpopular", func() {
			bp := three.BitPosition{}
			bp.AddVal(1)
			bp.AddVal(1)
			bp.AddVal(0)

			So(bp.UnPopularValue(), ShouldEqual, 0)
		})

		Convey("Determines 0 as unpopular when equal", func() {
			bp := three.BitPosition{}
			bp.AddVal(1)
			bp.AddVal(1)
			bp.AddVal(0)
			bp.AddVal(0)

			So(bp.UnPopularValue(), ShouldEqual, 0)
		})
	})
}
