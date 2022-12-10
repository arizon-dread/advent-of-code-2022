package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Open file
	dat, err := os.Open("input.txt")
	var games []game
	if err != nil {
		fmt.Printf("Could not read game input file, %v\n", err)
	}
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split((bufio.ScanLines))
	fmt.Printf("Reading lines in file...\n")
	for fileScanner.Scan() {
		t := fileScanner.Text()
		strArr := strings.Split(t, " ")
		var game game
		for i, item := range strArr {
			var err error
			fmt.Printf("index: %d, item: %v\n", i, item)
			if i == 0 {
				game.opponent, err = getRPS(item)
			} else {
				game.player, err = getRPS(item)
				fmt.Printf("After aligning values: opponent: %v, player: %v\n", game.opponent, game.player)
			}
			if err != nil {
				fmt.Printf("got unexpected input, here: %v or here: %v\n", game.opponent, game.player)
			}

		}
		calculateScoreForGame(&game)
		fmt.Printf("gamescore: %d\n", game.score)
		games = append(games, game)
	}
	fmt.Printf("number of entries in games: %d\n", len(games))
	var totalScore int
	for _, g := range games {
		fmt.Printf("totalscore: %d, gamescore to add: %d. ", totalScore, g.score)
		totalScore += g.score
		fmt.Printf("after adding: %d\n", totalScore)
	}
	fmt.Printf("Total score for player: %d\n", totalScore)

}

func getRPS(r string) (rune, error) {
	if r == "A" || r == "X" {
		return 'A', nil
	}
	if r == "B" || r == "Y" {
		return 'B', nil
	}
	if r == "C" || r == "Z" {
		return 'C', nil
	}
	return 'D', errors.New("unexpected input")
}

func calculateScoreForGame(game *game) {
	//find out who won
	if game.opponent == game.player {
		fmt.Printf("DRAW, opponent == player\n")
		game.score = 3
	} else if game.player == game.opponent+1 || game.player+2 == game.opponent {
		fmt.Printf("WIN! %d < %d!\n", game.opponent, game.player)
		game.score = 6
	} else {
		fmt.Printf("LOSE!\n")
	}
	//add score for rock/paper/scissors
	if game.player == 'A' {
		game.score += 1
	} else if game.player == 'B' {
		game.score += 2
	} else if game.player == 'C' {
		game.score += 3
	}
}
