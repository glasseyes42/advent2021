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

	Convey("completions", t, func() {
		Convey("completes something simple", func() {
			p, err := ten.Parse("[")
			So(err, ShouldBeNil)
			result := p.Complete()
			So(result, ShouldResemble, []string{"]"})
		})

		Convey("completes something nested", func() {
			p, err := ten.Parse("[()<")
			So(err, ShouldBeNil)
			result := p.Complete()
			So(result, ShouldResemble, []string{">", "]"})
		})

		Convey("completes a valid line", func() {
			p, err := ten.Parse("[({(<(())[]>[[{[]{<()<>>")
			So(err, ShouldBeNil)
			result := p.Complete()
			So(result, ShouldResemble, []string{"}", "}", "]", "]", ")", "}", ")", "]"})
		})

		Convey("completes a valid line - 2", func() {
			p, err := ten.Parse("[(()[<>])]({[<{<<[]>>(")
			So(err, ShouldBeNil)
			result := p.Complete()
			So(result, ShouldResemble, []string{")", "}", ">", "]", "}", ")"})
		})
	})

	Convey("part 2", t, func() {
		result := ten.ScoreCompletions(testData)
		So(result, ShouldEqual, 288957)
	})
}
