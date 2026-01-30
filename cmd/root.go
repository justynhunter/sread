package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/justynhunter/speedreader/lib"
	"github.com/justynhunter/speedreader/ui"
	"github.com/spf13/cobra"
)

var delayInMs *int

var rootCmd = &cobra.Command{
	Use:   "speedread",
	Short: "speadread is a cli tool for reading a document quickly.",
	Long:  "speadread is a cli tool for reading a document quickly.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	delayInMs = rootCmd.Flags().IntP("delay", "d", 300, "Delay between words in milliseconds")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	wordProcessor, err := lib.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(delayInMs)

	p := tea.NewProgram(ui.UiModel{DelayInMs: *delayInMs, WordProcessor: *wordProcessor}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
