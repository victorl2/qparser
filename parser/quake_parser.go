package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParserQuakeGameFile(logFilePath string, gameID int) *QuakeGames {
	fmt.Println(logFilePath, gameID)

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
	fmt.Println("Creating games from log file")
	groupedQuakeGames := NewGroupQuakeGames()
	var currentGameId = 0
	var currentGame *Game = nil

	for scanner.Scan() {
		logLine := scanner.Text()
		event := logLine[7:12]
		fmt.Println(event)

		switch event {
		case "Kill:":
			if targetGame == currentGameId || targetGame == -1 {
				fmt.Println("Adding kill")
				killing := parseKillLine(logLine)
				currentGame.AddKill(killing)
			}
		case "InitG":
			currentGameId++
			if targetGame == currentGameId || targetGame == -1 {
				fmt.Println("Creating new game obj")
				currentGame = NewGame()
			}
		case "Shutd":
			if targetGame == currentGameId || targetGame == -1 {
				fmt.Println("Flushing game")
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

	fmt.Printf("Killer: %s, Killed: %s, Weapon: %s\n", killerName, killedName, weaponUsed)
	return &Killing{killerName, killedName, weaponUsed}
}
