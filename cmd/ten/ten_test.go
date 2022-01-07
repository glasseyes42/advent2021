package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/glasseyes42/advent2021/internal/ten"
	. "github.com/smartystreets/goconvey/convey"
)

//go:embed data/sample.txt
var testData string

func TestTen(t *testing.T) {
	Convey("parse", t, func() {
		Convey("parses a line", func() {
			_, err := ten.Parse("[<>({}){}[([])<>]]")
			So(err, ShouldBeNil)

			_, err = ten.Parse("[({(<(())[]>[[{[]{<()<>>")
			So(err, ShouldBeNil)
		})

		Convey("parses an unfinished line", func() {
			_, err := ten.Parse("[(()[<>])]({[<{<<[]>>(")
			So(err, ShouldBeNil)
		})

		Convey("finds corruption", func() {
			_, err := ten.Parse("{([(<{}[<>[]}>{[]{[(<()>")
			So(err, ShouldResemble, fmt.Errorf("expected ], but found }"))
		})
	})

	Convey("part 1", t, func() {
		result := ten.ScoreIllegalLines(testData)
		So(result, ShouldEqual, 26397)
	})

}
