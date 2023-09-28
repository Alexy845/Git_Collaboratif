package main

import (
	"fmt"
	"strings"
)

var (
	board    [3][3]string
	player   string
	gameOver = false
)

func main() {
	initializeBoard()
	player = "X"

	for !gameOver {
		printBoard()
		move := getInput()
		if IsValidMove(move) {
			board[move[0]][move[1]] = player
			if hasWon(player) {
				printBoard()
				fmt.Printf("Le joueur %s a gagné !\n", player)
				gameOver = true
			} else if isBoardFull() {
				printBoard()
				fmt.Println("Match nul !")
				gameOver = true
			} else {
				switchPlayer()
			}
		} else {
			fmt.Println("Mouvement invalide. Réessayez.")
		}
	}
}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "-"
		}
	}
}

func printBoard() {
	fmt.Println("Tableau du Morpion :")
	for i := 0; i < 3; i++ {
		fmt.Printf("%s | %s | %s\n", strings.Join(board[i][:], " | "))
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
	fmt.Println()
}

func getInput() [2]int {
	var row, col int
	fmt.Print("Entrez la ligne (0, 1, ou 2) : ")
	fmt.Scan(&row)
	fmt.Print("Entrez la colonne (0, 1, ou 2) : ")
	fmt.Scan(&col)
	return [2]int{row, col}
}

func IsValidMove(move [2]int) bool {
	row, col := move[0], move[1]
	if row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == "-" {
		return true
	}
	return false
}

func hasWon(player string) bool {
	// Vérification des lignes
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}

	// Vérification des colonnes
	for j := 0; j < 3; j++ {
		if board[0][j] == player && board[1][j] == player && board[2][j] == player {
			return true
		}
	}

	// Vérification des diagonales
	if (board[0][0] == player && board[1][1] == player && board[2][2] == player) ||
		(board[0][2] == player && board[1][1] == player && board[2][0] == player) {
		return true
	}

	return false
}

func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	return true
}

func switchPlayer() {
	if player == "X" {
		player = "O"
	} else {
		player = "X"
	}
}
