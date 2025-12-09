package models

import (
	"fmt"
	"time"

	"github.com/Arthur-Conti/pomodo_cli/styles"
	"github.com/Arthur-Conti/pomodo_cli/ticker"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Opts ModelOpts
	timeLeft    time.Duration   // Quanto tempo falta
	totalTime   time.Duration   // O tempo total inicial
	tickInterval time.Duration
	percent     float64         // Porcentagem completada (0.0 a 1.0)
	progress    progress.Model  // O sub-modelo da barra de progresso
	finished    bool            // Se acabou
	tick ticker.Ticker
	style styles.Style
}

func NewModel(opts ModelOpts) *Model {
	prog := progress.New(progress.WithGradient(string(opts.StartColor), string(opts.EndColor)))
	return &Model{
		timeLeft:  opts.PomodoroDuration,
		totalTime: opts.PomodoroDuration,
		tickInterval: opts.TickInterval,
		percent:   1.0, // Começa em 100%
		progress:  prog,
		finished:  false,
		tick: opts.Ticker,
		style: opts.Style,
	}
}

func (m *Model) Init() tea.Cmd {
	return m.tick.Tick()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Se o usuário apertou uma tecla
	case tea.KeyMsg:
		if msg.String() == "q" || msg.Type == tea.KeyCtrlC {
			// Sai do programa
			return m, tea.Quit
		}

	// Se a mensagem for um "tick" do nosso timer
	case ticker.TickMsg:
		if m.timeLeft > 0 {
			m.timeLeft -= m.tickInterval

			// Calcula a nova porcentagem
			m.percent = float64(m.timeLeft) / float64(m.totalTime)

			// Comandos para executar na próxima passada:
			// 1. Agendar o próximo tick
			// 2. Pedir para a barra de progresso animar para a nova porcentagem
			cmdTick := m.tick.Tick()
			cmdProg := m.progress.SetPercent(m.percent)
			return m, tea.Batch(cmdTick, cmdProg)
		}

		// Se o tempo acabou
		m.finished = true
		return m, tea.Quit

	// Mensagem interna da barra de progresso (para animações suaves)
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m *Model) View() string {
	if m.finished {
		return "\n  🍅 POMODORO FINALIZADO! 🍅\n\n"
	}

	// Formata o tempo restante (ex: 24:59)
	minutes := int(m.timeLeft.Minutes())
	seconds := int(m.timeLeft.Seconds()) % 60
	timeString := fmt.Sprintf("%02d:%02d", minutes, seconds)

	// Renderiza a barra de progresso
	barView := m.progress.View()

	// Monta a tela final
	return fmt.Sprintf(
		"\n  Foco total: %s\n\n  %s\n\n  %s\n",
		timeString,    // O texto do tempo
		barView,       // A barra visual
		m.style.Help("(Pressione q para sair)"), // Ajuda
	)
}