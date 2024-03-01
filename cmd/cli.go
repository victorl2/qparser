package cmd

import (
	"encoding/json"
	"fmt"
	"log"
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
		groupedGames := parser.ParserQuakeGameFile(file, gameID)

		jsonData, err := json.MarshalIndent(groupedGames, "", "  ")
		if err != nil {
			log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
		}
		fmt.Println(string(jsonData))
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
