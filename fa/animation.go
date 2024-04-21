package fa

import (
	"strconv"
	"time"

	b "github.com/willoma/bulma-gomponents"
)

type AnimationDirection string

const (
	Normal           AnimationDirection = "normal"
	Reverse          AnimationDirection = "reverse"
	Alternate        AnimationDirection = "alternate"
	AlternateReverse AnimationDirection = "alternate-reverse"
)

type AnimationTiming string

const (
	Ease      AnimationTiming = "ease"
	Linear    AnimationTiming = "linear"
	EaseIn    AnimationTiming = "ease-in"
	EaseOut   AnimationTiming = "ease-out"
	EaseInOut AnimationTiming = "ease-in-out"
	StepStart AnimationTiming = "step-start"
	StepEnd   AnimationTiming = "step-end"
)

type AnimationType int

const (
	Beat AnimationType = iota
	Fade
	BeatFade
	Bounce
	Flip
	Shake
	Spin
)

type Animation struct {
	Type AnimationType

	// Common parameters
	Delay          time.Duration
	Direction      AnimationDirection
	Duration       time.Duration
	IterationCount float64
	Timing         AnimationTiming

	// Beat and BeatFade parameters
	MaxScale float64

	// Fade and BeatFade parameters
	MinOpacity float64

	// Bounce parameters
	Rebound     float64
	Height      float64
	StartScaleX float64
	StartScaleY float64
	JumpScaleX  float64
	JumpScaleY  float64
	LandScaleX  float64
	LandScaleY  float64

	// Flip parameters
	X     float64
	Y     float64
	Z     float64
	Angle float64

	// Spin parameters
	Pulse   bool
	Reverse bool
}

func (a Animation) ClassesAndStyles() ([]b.Class, b.Styles) {
	var (
		classes    []b.Class
		stylesArgs []string
	)
	switch a.Type {
	case Beat:
		classes = append(classes, "fa-beat")
		if a.MaxScale != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-beat-scale", strconv.FormatFloat(a.MaxScale, 'f', 2, 64),
			)
		}

	case Fade:
		classes = append(classes, "fa-fade")
		if a.MinOpacity != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-fade-opacity", strconv.FormatFloat(a.MinOpacity, 'f', 2, 64),
			)
		}

	case BeatFade:
		classes = append(classes, "fa-beat-fade")
		if a.MaxScale != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-beat-fade-scale", strconv.FormatFloat(a.MaxScale, 'f', 2, 64),
			)
		}
		if a.MinOpacity != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-beat-fade-opacity", strconv.FormatFloat(a.MinOpacity, 'f', 2, 64),
			)
		}

	case Bounce:
		classes = append(classes, "fa-bounce")
		if a.Rebound != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-bounce-rebound", strconv.FormatFloat(a.MaxScale, 'f', 2, 64),
			)
		}
		if a.Height != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-bounce-height", strconv.FormatFloat(a.Height, 'f', 2, 64),
			)
		}
		if a.StartScaleX != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-bounce-start-scale-x", strconv.FormatFloat(a.StartScaleX, 'f', 2, 64),
			)
		}
		if a.StartScaleY != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-bounce-start-scale-y", strconv.FormatFloat(a.StartScaleY, 'f', 2, 64),
			)
		}
		if a.JumpScaleX != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-bounce-jump-scale-x", strconv.FormatFloat(a.JumpScaleX, 'f', 2, 64),
			)
		}
		if a.JumpScaleY != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-bounce-jump-scale-y", strconv.FormatFloat(a.JumpScaleY, 'f', 2, 64),
			)
		}
		if a.LandScaleX != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-bounce-land-scale-x", strconv.FormatFloat(a.LandScaleX, 'f', 2, 64),
			)
		}
		if a.LandScaleY != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-bounce-land-scale-y", strconv.FormatFloat(a.LandScaleY, 'f', 2, 64),
			)
		}

	case Flip:
		classes = append(classes, "fa-flip")
		if a.X != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-flip-x", strconv.FormatFloat(a.X, 'f', 2, 64),
			)
		}
		if a.Y != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-flip-y", strconv.FormatFloat(a.Y, 'f', 2, 64),
			)
		}
		if a.Z != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-flip-z", strconv.FormatFloat(a.Z, 'f', 2, 64),
			)
		}
		if a.Angle != 0 {
			stylesArgs = append(
				stylesArgs,
				"--fa-flip-angle", strconv.FormatFloat(a.Angle, 'f', 2, 64),
			)
		}

	case Shake:
		classes = append(classes, "fa-shake")

	case Spin:
		if a.Pulse {
			classes = append(classes, "fa-spin-pulse")
		} else {
			classes = append(classes, "fa-spin")
		}
		if a.Reverse {
			classes = append(classes, "fa-spin-reverse")
		}

	default:
		return nil, nil
	}

	if a.Delay != 0 {
		stylesArgs = append(
			stylesArgs,
			"--fa-animation-delay", strconv.FormatFloat(a.Delay.Seconds(), 'f', 2, 64)+"s",
		)
	}

	if a.Direction != "" {
		stylesArgs = append(
			stylesArgs,
			"--fa-animation-direction", string(a.Direction),
		)
	}

	if a.Duration != 0 {
		stylesArgs = append(
			stylesArgs,
			"--fa-animation-duration", strconv.FormatFloat(a.Duration.Seconds(), 'f', 2, 64)+"s",
		)
	}

	if a.IterationCount != 0 {
		stylesArgs = append(
			stylesArgs,
			"--fa-animation-iteration-count", strconv.FormatFloat(a.IterationCount, 'f', 2, 64),
		)
	}

	if a.Timing != "" {
		stylesArgs = append(
			stylesArgs,
			"--fa-animation-timing", string(a.Timing),
		)
	}

	return classes, b.Style(stylesArgs...)
}
