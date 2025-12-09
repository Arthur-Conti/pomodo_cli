package main

import (
	"fmt"
	"time"
)

func main() {
	totalTime := 5 * time.Second

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for range totalTime {
		select {
		case <- ticker.C: 
			totalTime -= time.Second * 1
			if totalTime <= 0 {
				fmt.Print("\033[H\033[2J") // Limpa a tela
				fmt.Println("POMODORO FINALIZADO! 🍅")
				return // Sai do programa
			}
			
			minutes := int(totalTime.Minutes())
			seconds := int(totalTime.Seconds()) % 60
			
			fmt.Print("\033[H\033[2J") // Limpa a tela
			fmt.Printf("Tempo restante: %02d:%02d\n", minutes, seconds)
		}
	}
}
