package cmd

import (
	"fmt"
	"os"
	"qparser/parser"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qparser [file]",
	Short: "qparser is a CLI tool for parsing log files from Quake 3 Arena server",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		gameID, _ := cmd.Flags().GetInt("game")
		parser.ParserQuakeGameFile(file, gameID)
	},
}

func init() {
	rootCmd.PersistentFlags().Int("game", -1, "Specify the game ID to summarize")
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
