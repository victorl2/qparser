package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"qparser/parser"
	"sort"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qparser [file]",
	Short: "qparser is a CLI tool for parsing log files from Quake 3 Arena server",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		gameID, _ := cmd.Flags().GetInt("game")
		quakeGames := parser.ParserQuakeGameFile(file, gameID)
		printJsonQuakeGames(quakeGames)
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Int("game", -1, "Specify the game ID to summarize")
}

func printJsonQuakeGames(quakeGames *parser.QuakeGames) {
	keys := make([]int, 0)
	for k := range quakeGames.GameDetails {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, gameID := range keys {
		game := quakeGames.GameDetails[gameID]
		gameMap := map[string]*parser.Game{fmt.Sprintf("game_%d", gameID): game}
		gameJSON, err := json.MarshalIndent(gameMap, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling game %d: %v", gameID, err)
		}
		fmt.Println(string(gameJSON))
		fmt.Println()
	}
}
