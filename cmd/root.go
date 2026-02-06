package cmd

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/justynhunter/speedreader/lib"
	"github.com/justynhunter/speedreader/ui"
	"github.com/spf13/cobra"
)

var highlightColor *string
var noHighlight *bool
var fileName *string
var wordsPerMinute *int

var rootCmd = &cobra.Command{
	Use:   "speedread [FILE]",
	Short: "speadread is a cli tool for reading a document quickly.",
	Long:  "speadread is a cli tool for reading a document quickly.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fileName = &args[0]
		}
	},
}

func init() {
	highlightColor = rootCmd.Flags().StringP("highlight-color", "c", "#98FF98", "Color of the highlighted character in hex (e.g. #98FF98")
	noHighlight = rootCmd.Flags().BoolP("no-highlight", "n", false, "Don't highlight the 'center' character in the word")
	wordsPerMinute = rootCmd.Flags().IntP("words-per-minute", "w", 300, "words per minute")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	var wordProcessor *lib.WordProcessor
	var err error

	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	isTerminal := (fileInfo.Mode() & os.ModeCharDevice) != 0

	if isTerminal {
		if fileName == nil {
			if err = rootCmd.Help(); err != nil {
				log.Fatal(err)
			}
			return
		}

		wordProcessor, err = lib.ReadFile(*fileName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		wordProcessor, err = lib.ReadInput()
		if err != nil {
			log.Fatal(err)
		}
	}

	p := tea.NewProgram(ui.UiModel{WordsPerMinute: *wordsPerMinute, HighlightColor: *highlightColor, NoHighlight: *noHighlight, WordProcessor: *wordProcessor}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
