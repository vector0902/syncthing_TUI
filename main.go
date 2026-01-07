package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	"github.com/pdrolopes/syncthing_TUI/app"
)

func main() {
	homeDir := flag.String("home", "", "Syncthing home directory path (alternative to STHOME env var)")
	flag.Parse()

	zone.NewGlobal()
	p := tea.NewProgram(app.NewModel(*homeDir), tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
