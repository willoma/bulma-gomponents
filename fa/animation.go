package fa

import (
	"strconv"
	"time"

	b "github.com/willoma/bulma-gomponents"
)

type Animation interface {
	b.ParentModifier

	setDelay(time.Duration)
	setDirection(string)
	setDuration(time.Duration)
	setIterationCount(float64)
	setTiming(string)
}

type animationBase struct {
	delay          time.Duration
	direction      string
	duration       time.Duration
	iterationCount float64
	timing         string
}

func (a *animationBase) ModifyParent(e b.Element) {
	if a.delay != 0 {
		e.With(b.Style(
			"--fa-animation-delay",
			strconv.FormatFloat(a.delay.Seconds(), 'f', 2, 64)+"s",
		))
	}

	if a.direction != "" {
		e.With(b.Style("--fa-animation-direction", a.direction))
	}

	if a.duration != 0 {
		e.With(b.Style(
			"--fa-animation-duration",
			strconv.FormatFloat(a.duration.Seconds(), 'f', 2, 64)+"s",
		))
	}

	if a.iterationCount != 0 {
		e.With(b.Style(
			"--fa-animation-iteration-count",
			strconv.FormatFloat(a.iterationCount, 'f', 2, 64),
		))
	}

	if a.timing != "" {
		e.With(b.Style("--fa-animation-timing", a.timing))
	}
}

func (a *animationBase) setDelay(d time.Duration) {
	a.delay = d
}

func Delay(d time.Duration) func(a Animation) {
	return func(a Animation) {
		a.setDelay(d)
	}
}

func (a *animationBase) setDirection(d string) {
	a.direction = d
}

func DirectionNormal(a Animation) {
	a.setDirection("normal")
}

func DirectionReverse(a Animation) {
	a.setDirection("reverse")
}

func DirectionAlternate(a Animation) {
	a.setDirection("alternate")
}

func DirectionAlternateReverse(a Animation) {
	a.setDirection("alternate-reverse")
}

func (a *animationBase) setDuration(d time.Duration) {
	a.duration = d
}

func Duration(d time.Duration) func(a Animation) {
	return func(a Animation) {
		a.setDuration(d)
	}
}

func (a *animationBase) setIterationCount(c float64) {
	a.iterationCount = c
}

func IterationCount(c float64) func(a Animation) {
	return func(a Animation) {
		a.setIterationCount(c)
	}
}

func (a *animationBase) setTiming(t string) {
	a.timing = t
}

func TimingEase(a Animation) {
	a.setTiming("ease")
}

func TimingLinear(a Animation) {
	a.setTiming("linear")
}

func TimingEaseIn(a Animation) {
	a.setTiming("ease-in")
}

func TimingEaseOut(a Animation) {
	a.setTiming("ease-out")
}

func TimingEaseInOut(a Animation) {
	a.setTiming("ease-in-out")
}

func TimingStepStart(a Animation) {
	a.setTiming("step-start")
}

func TimingStepEnd(a Animation) {
	a.setTiming("step-end")
}
