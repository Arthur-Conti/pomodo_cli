package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Arthur-Conti/pomodo_cli/models"
	"github.com/Arthur-Conti/pomodo_cli/styles"
	"github.com/Arthur-Conti/pomodo_cli/ticker"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Cria e roda o programa Bubbletea
	p := tea.NewProgram(setup())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro ao rodar o programa: %v", err)
		os.Exit(1)
	}
}

func setup() *models.Model {
	style := styles.NewStyle()
	tickInterval := 1 * time.Second
	tick := ticker.NewTicker(tickInterval)
	modelOpts := models.BaseModelOpts(*tick, *style)
	return models.NewModel(modelOpts)
}
