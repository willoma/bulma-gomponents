package fa

import (
	"strconv"
	"time"

	e "github.com/willoma/gomplements"
)

type Animation interface {
	isAnimation()

	ModifyParent(p e.Element)
	If(bool) Animation
}

type animationBase struct {
	delay     time.Duration
	zeroDelay bool

	direction string

	duration     time.Duration
	zeroDuration bool

	iterationCount     float64
	zeroIterationCount bool

	timing string
}

func (a *animationBase) isAnimation() {}

func (a *animationBase) modifyParent(p e.Element) {
	if a.delay != 0 {
		p.With(e.Style(
			"--fa-animation-delay",
			strconv.FormatFloat(a.delay.Seconds(), 'f', 2, 64)+"s",
		))
	} else if a.zeroDelay {
		p.With(e.Style("--fa-animation-delay", "0s"))
	}

	if a.direction != "" {
		p.With(e.Style("--fa-animation-direction", a.direction))
	}

	if a.duration != 0 {
		p.With(e.Style(
			"--fa-animation-duration",
			strconv.FormatFloat(a.duration.Seconds(), 'f', 2, 64)+"s",
		))
	} else if a.zeroDuration {
		p.With(e.Style("--fa-animation-duration", "0s"))
	}

	if a.iterationCount != 0 {
		p.With(e.Style(
			"--fa-animation-iteration-count",
			strconv.FormatFloat(a.iterationCount, 'f', 2, 64),
		))
	} else if a.zeroIterationCount {
		p.With(e.Style("--fa-animation-iteration-count", "0"))
	}

	if a.timing != "" {
		p.With(e.Style("--fa-animation-timing", a.timing))
	}
}

func Delay(d time.Duration) func(a any) {
	return func(a any) {
		if b, ok := a.(*animationBase); ok {
			b.zeroDelay = d == 0
			b.delay = d
		}
	}
}

func DirectionNormal(a any) {
	if b, ok := a.(*animationBase); ok {
		b.direction = "normal"
	}
}

func DirectionReverse(a any) {
	if b, ok := a.(*animationBase); ok {
		b.direction = "reverse"
	}
}

func DirectionAlternate(a any) {
	if b, ok := a.(*animationBase); ok {
		b.direction = "alternate"
	}
}

func DirectionAlternateReverse(a any) {
	if b, ok := a.(*animationBase); ok {
		b.direction = "alternate-reverse"
	}
}

func Duration(d time.Duration) func(a any) {
	return func(a any) {
		if b, ok := a.(*animationBase); ok {
			b.zeroDuration = d == 0
			b.duration = d
		}
	}
}

func IterationCount(c float64) func(a any) {
	return func(a any) {
		if b, ok := a.(*animationBase); ok {
			b.zeroIterationCount = c == 0
			b.iterationCount = c
		}
	}
}

func (a *animationBase) setTiming(t string) {
	a.timing = t
}

func TimingEase(a any) {
	if b, ok := a.(*animationBase); ok {
		b.timing = "ease"
	}
}

func TimingLinear(a any) {
	if b, ok := a.(*animationBase); ok {
		b.timing = "linear"
	}
}

func TimingEaseIn(a any) {
	if b, ok := a.(*animationBase); ok {
		b.timing = "ease-in"
	}
}

func TimingEaseOut(a any) {
	if b, ok := a.(*animationBase); ok {
		b.timing = "ease-out"
	}
}

func TimingEaseInOut(a any) {
	if b, ok := a.(*animationBase); ok {
		b.timing = "ease-in-out"
	}
}

func TimingStepStart(a any) {
	if b, ok := a.(*animationBase); ok {
		b.timing = "step-start"
	}
}

func TimingStepEnd(a any) {
	if b, ok := a.(*animationBase); ok {
		b.timing = "step-end"
	}
}
