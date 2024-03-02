package fa

import (
	"fmt"

	b "github.com/willoma/bulma-gomponents"
)

type Animation interface {
	attrs() (Class, b.Styles)
}

type AnimationDirection string

const (
	AnimationDirectionNormal           AnimationDirection = "normal"
	AnimationDirectionReverse          AnimationDirection = "reverse"
	AnimationDirectionAlternate        AnimationDirection = "alternate"
	AnimationDirectionAlternateReverse AnimationDirection = "alternate-reverse"
)

type AnimationTiming string

const (
	AnimationTimingEase      AnimationTiming = "ease"
	AnimationTimingLinear    AnimationTiming = "linear"
	AnimationTimingEaseIn    AnimationTiming = "ease-in"
	AnimationTimingEaseOut   AnimationTiming = "ease-out"
	AnimationTimingEaseInOut AnimationTiming = "ease-in-out"
	AnimationTimingStepStart AnimationTiming = "step-start"
	AnimationTimingStepEnd   AnimationTiming = "step-end"
)

type AnimationBase struct {
	Delay          float64 // seconds
	Direction      AnimationDirection
	Duration       float64 // seconds
	IterationCount float64
	Timing         AnimationTiming
}

func (a AnimationBase) baseStyles() b.Styles {
	styles := b.Styles{}
	if a.Delay != 0 {
		styles["--fa-animation-delay"] = fmt.Sprintf("%vs", a.Delay)
	}
	if a.Direction != "" {
		styles["--fa-animation-direction"] = string(a.Direction)
	}
	if a.Duration != 0 {
		styles["--fa-animation-duration"] = fmt.Sprintf("%vs", a.Duration)
	}
	if a.IterationCount != 0 {
		styles["--fa-animation-iteration-count"] = fmt.Sprintf("%v", a.IterationCount)
	}
	if a.Timing != "" {
		styles["--fa-animation-timing"] = string(a.Timing)
	}
	return styles
}

// Beat adds a beat animation to a FA icon.
type Beat struct {
	AnimationBase
	MaxScale float64
}

func (a Beat) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	if a.MaxScale != 0 {
		styles["--fa-beat-scale"] = fmt.Sprintf("%v", a.MaxScale)
	}
	return "fa-beat", styles
}

// Fade adds a fade animation to a FA icon.
type Fade struct {
	AnimationBase
	MinOpacity float64
}

func (a Fade) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	if a.Duration != 0 {
		styles["--fa-animation-duration"] = fmt.Sprintf("%vs", a.Duration)
	}
	if a.MinOpacity != 0 {
		styles["--fa-fade-opacity"] = fmt.Sprintf("%v", a.MinOpacity)
	}
	return "fa-fade", styles
}

// BeatFade adds a beat and a fade animations to a FA icon.
type BeatFade struct {
	AnimationBase
	MaxScale   float64
	MinOpacity float64
}

func (a BeatFade) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	if a.Duration != 0 {
		styles["--fa-animation-duration"] = fmt.Sprintf("%vs", a.Duration)
	}
	if a.MaxScale != 0 {
		styles["--fa-beat-fade-scale"] = fmt.Sprintf("%v", a.MaxScale)
	}
	if a.MinOpacity != 0 {
		styles["--fa-beat-fade-opacity"] = fmt.Sprintf("%v", a.MinOpacity)
	}
	return "fa-beat-fade", styles
}

// Bounce adds a bounce animation to a FA icon.
type Bounce struct {
	AnimationBase
	Rebound     float64
	Height      float64
	StartScaleX float64
	StartScaleY float64
	JumpScaleX  float64
	JumpScaleY  float64
	LandScaleX  float64
	LandScaleY  float64
}

func (a Bounce) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	if a.Duration != 0 {
		styles["--fa-animation-duration"] = fmt.Sprintf("%vs", a.Duration)
	}
	if a.Rebound != 0 {
		styles["--fa-bounce-rebound"] = fmt.Sprintf("%v", a.Rebound)
	}
	if a.Height != 0 {
		styles["--fa-bounce-height"] = fmt.Sprintf("%v", a.Height)
	}
	if a.StartScaleX != 0 {
		styles["--fa-bounce-start-scale-x"] = fmt.Sprintf("%v", a.StartScaleX)
	}
	if a.StartScaleY != 0 {
		styles["--fa-bounce-start-scale-y"] = fmt.Sprintf("%v", a.StartScaleY)
	}
	if a.JumpScaleX != 0 {
		styles["--fa-bounce-jump-scale-x"] = fmt.Sprintf("%v", a.JumpScaleX)
	}
	if a.JumpScaleY != 0 {
		styles["--fa-bounce-jump-scale-y"] = fmt.Sprintf("%v", a.JumpScaleY)
	}
	if a.LandScaleX != 0 {
		styles["--fa-bounce-land-scale-x"] = fmt.Sprintf("%v", a.LandScaleX)
	}
	if a.LandScaleY != 0 {
		styles["--fa-bounce-land-scale-y"] = fmt.Sprintf("%v", a.LandScaleY)
	}

	return "fa-bounce", styles
}

// Flip adds a flip animation to a FA icon.
type Flip struct {
	AnimationBase
	X     float64
	Y     float64
	Z     float64
	Angle float64
}

func (a Flip) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	if a.Duration != 0 {
		styles["--fa-animation-duration"] = fmt.Sprintf("%vs", a.Duration)
	}
	if a.X != 0 {
		styles["--fa-flip-x"] = fmt.Sprintf("%v", a.X)
	}
	if a.Y != 0 {
		styles["--fa-flip-y"] = fmt.Sprintf("%v", a.Y)
	}
	if a.Z != 0 {
		styles["--fa-flip-z"] = fmt.Sprintf("%v", a.Z)
	}
	if a.Angle != 0 {
		styles["--fa-flip-angle"] = fmt.Sprintf("%v", a.Angle)
	}
	return "fa-flip", styles
}

// Shake adds a shake animation to a FA icon.
type Shake struct {
	AnimationBase
}

func (a Shake) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	return "fa-shake", styles
}

// Spin adds a spin animation to a FA icon.
type Spin struct {
	AnimationBase
}

func (a Spin) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	return "fa-spin", styles
}

// SpinReverse adds a reverse spin animation to a FA icon.
type SpinReverse struct {
	AnimationBase
}

func (a SpinReverse) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	return "fa-spin fa-spin-reverse", styles
}

// SpinPulse adds a 8-steps spin animation to a FA icon.
type SpinPulse struct {
	AnimationBase
}

func (a SpinPulse) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	return "fa-spin-pulse", styles
}

// SpinPulseReverse adds a 8-steps reverse spin animation to a FA icon.
type SpinPulseReverse struct {
	AnimationBase
}

func (a SpinPulseReverse) attrs() (Class, b.Styles) {
	styles := a.baseStyles()
	return "fa-spin-pulse fa-spin-reverse", styles
}
