package models

import (
	"time"

	"github.com/Arthur-Conti/pomodo_cli/styles"
	"github.com/Arthur-Conti/pomodo_cli/ticker"
	"github.com/charmbracelet/lipgloss"
)

const (
	baseEndColor = lipgloss.Color("#04B575")
	baseStartColor = lipgloss.Color("#FF4136")
)

type ModelOpts struct {
	PomodoroDuration time.Duration
	TickInterval time.Duration
	StartColor lipgloss.Color
	EndColor lipgloss.Color
	Ticker ticker.Ticker
	Style styles.Style
}

func BaseModelOpts(tick ticker.Ticker, style styles.Style) ModelOpts {
	return ModelOpts{
		PomodoroDuration: 1 * time.Minute,
		TickInterval: 1 * time.Second,
		StartColor: baseStartColor,
		EndColor: baseEndColor,
		Ticker: tick,
		Style: style,
	}
}