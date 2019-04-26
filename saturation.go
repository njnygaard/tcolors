package main

import (
	"fmt"
	"math"

	"github.com/gdamore/tcell"
	"github.com/teacat/noire"
)

func roundFloat(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

const (
	satMin   = 0.0
	satMax   = 100.0
	satIncr  = 0.5
	satCount = int(satMax/satIncr) + 1
)

type SaturationBar struct {
	*NavBar
	scale [satCount]float64
}

func NewSaturationBar(s *State) *SaturationBar {
	bar := &SaturationBar{NavBar: NewNavBar(s, satCount)}

	i := satMin
	for n, _ := range bar.scale {
		bar.scale[n] = i
		i += 0.5
	}
	bar.scale[0] = 0.01

	return bar
}

// Draw redraws bar at given coordinates and screen, returning the number
// of rows occupied
func (bar *SaturationBar) Draw(x, y int, s tcell.Screen) int {
	h := bar.NavBar.Draw(x, y, s)
	return h + 1
}

// State change handler
func (bar *SaturationBar) Handle(change StateChange) {
	var nc *noire.Color

	if change.Includes(HueChanged, ValueChanged) {
		nc = bar.state.BaseColor()

		for n, val := range bar.scale {
			nc = noire.NewHSV(bar.state.Hue(), val, bar.state.Value())
			bar.items[n] = toTColor(nc)
		}
	}

	if change.Includes(SaturationChanged) {
		bar.SetPos(roundFloat(bar.state.Saturation() / satIncr))
		bar.SetLabel(fmt.Sprintf("%5.1f ", bar.scale[bar.pos]))
	}
}

func (bar *SaturationBar) Up(step int) {
	bar.up(step)
	bar.setState()
}

func (bar *SaturationBar) Down(step int) {
	bar.down(step)
	bar.setState()
}

func (bar *SaturationBar) setState() {
	bar.state.SetSaturation(bar.scale[bar.pos])
}
