package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParserQuakeGameFile(logFilePath string, gameID int) *QuakeGames {
	logFile, err := os.Open(logFilePath)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return nil
	}
	defer logFile.Close()
	scanner := bufio.NewScanner(logFile)

	return createGamesFromLogFile(scanner, gameID)

}

func createGamesFromLogFile(scanner *bufio.Scanner, targetGame int) *QuakeGames {
	groupedQuakeGames := NewGroupQuakeGames()
	var currentGameId = 0
	var currentGame *Game = nil

	for scanner.Scan() {
		logLine := scanner.Text()
		event := logLine[7:12]

		switch event {
		case "Kill:":
			if targetGame == currentGameId || targetGame == -1 {
				killing := parseKillLine(logLine)
				currentGame.AddKill(killing)
			}
		case "InitG":
			currentGameId++
			if targetGame == currentGameId || targetGame == -1 {
				currentGame = NewGame()
			}
		case "Shutd":
			if targetGame == currentGameId || targetGame == -1 {
				groupedQuakeGames.AddGame(fmt.Sprintf("game_%d", currentGameId), currentGame)
			}
		}
	}

	return groupedQuakeGames
}

func parseKillLine(logLine string) *Killing {
	var infoStart = 12 // Skip the time prefix
	info := logLine[infoStart:]

	infoParts := strings.SplitN(info, ": ", 2)

	detailParts := strings.SplitN(infoParts[1], " by ", 2)
	names := strings.SplitN(detailParts[0], " killed ", 2)

	killerName := names[0]
	killedName := names[1]
	weaponUsed := detailParts[1]
	return &Killing{killerName, killedName, weaponUsed}
}
