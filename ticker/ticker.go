package ticker

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

type Ticker struct {
	TickInterval time.Duration
}

func NewTicker(interval time.Duration) *Ticker {
	return &Ticker{
		TickInterval: interval,
	}
}

func (t *Ticker) Tick() tea.Cmd {
	return tea.Tick(t.TickInterval, func(timer time.Time) tea.Msg {
		return TickMsg(timer)
	})
}
