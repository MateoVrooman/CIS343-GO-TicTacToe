/*
CIS 343
Prof. Woodring
Code written by: Mateo Vrooman
12/10/2024
Final Project
This program is a tic tac toe game in the console
Run with: go run app/main.go
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Main()
// Main function contains the primary logic for the function of the game. First, the board is initialized
// then the main game loop occurs. The board is printed, and the game checks for a win condition.
// If a player has won, the loop ends, otherwise the player can make a move and update the board.
func main(){

	var board [3][3]string
	var currentPlayer string

	initializeBoard(&board) // Pass the address of board as a parameter, since initializeBoard needs to modify the variable board
	currentPlayer = "X" // String currentPlayer stores which character is going to be playing next, "X" or "O"
	// For loop without condition creates infinite loop. Game loop ends when a player wins, or the board is full
	for {
		printBoard(board)
		if currentPlayer == "X" { // Player moves
			playerMove(&board, currentPlayer) // There is no end-game condition, so player moves
		} else {
			computerMove(&board, currentPlayer) // Computer randomly selects a move
		}
		
		if winner := checkWinner(board, currentPlayer); winner != "" { // := operator is used to declare and initialize variable
			printBoard(board)
			fmt.Printf("Player %s wins!\n", winner) // Print the winner
			break // Exit the game loop
		}
		if isBoardFull(board) {
			fmt.Println("It's a draw!") // No winner, board is full
			break // Exit the game loop
		}
		switchPlayer(&currentPlayer) // Pass the address of currentPlayer as a parameter
	}
}

// initializeBoard()
// Initializes the board as a 3x3 2 dimensional array.
// Parameters: 
// board: A pointer to a 3x3 2 dimensional array of strings
func initializeBoard(board *[3][3]string) {
	// Loop through both dimensions of board
	for i := range *board {
		for j := range board[i] {
			// Initialize each space as a string with a space
			board[i][j] = " "
		}
	}
}

// printBoard()
// Prints the board state while formatting it to be readable
// Parameters:
// board: A 3x3 2 dimensional array of strings
func printBoard(board [3][3]string) {
	fmt.Println("Current Board:")
	//Loop only through the rows of board
	for i := range board {
		// Print one row at a time
		fmt.Printf(" %s | %s | %s \n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			// Print a line in between rows to be more readable
			fmt.Println("---+---+---")
		}
	}
}

// playerMove()
// Prompts the player for a move, and checks for errors. If the move is valid it updates the board
// Parameters:
// board: A pointer to a 3x3 2 dimensional array of strings
// currentPlayer: A string containing the current player
func playerMove(board *[3][3]string, currentPlayer string) {
	var row, col int
	for {
		fmt.Printf("Player %s, enter your move (row and column): ", currentPlayer)
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err == nil && row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == " " {
			board[row][col] = currentPlayer
			break
		}
		fmt.Println("Invalid move. Try again.")
	}
}

// computerMove()
// Selects a random move and applies it to the board if that cell is empty
// Parameters:
// board: A pointer to a 3x3 2 dimensional array of strings
// currentPlayer: A string containing the current player
func computerMove(board *[3][3]string, currentPlayer string){
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		row := rng.Intn(3) // Random row between 0-2
		col := rng.Intn(3) // Random column between 0-2
		if (*board)[row][col] == " " {
			fmt.Printf("Computer chooses move: %d %d\n", row, col)
			(*board)[row][col] = currentPlayer // Place the O in the cell
			break // Exit the loop
		}
	}
}

// switchPlayer()
// Switches the currentPlayer string to the other player
// Parameters:
// currentPlayer: A pointer to a string containing the current player
func switchPlayer(currentPlayer *string) {
	if *currentPlayer == "X" {
		*currentPlayer = "O"
	} else {
		*currentPlayer = "X"
	}
}

// checkWinner()
// Checks the board for a winning condition, and the player that has won
// Parameters:
// board: A 3x3 2 dimensional array
// currentPlayer: A string containing the current player
// Returns the string representation of the winning player
// If there is no win condition returns an empty string
func checkWinner(board [3][3]string, currentPlayer string) string {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return currentPlayer
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return currentPlayer
		}
	}
	// Check diagonals
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return currentPlayer
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return currentPlayer
	}
	// No winner yet, return an empty string
	return ""
}

// isBoardFull()
// Checks to see if the board is completely full, indicating a draw
// Parameters:
// board: A 3x3 2 dimensional array of strings
// Returns a boolean that indicates if the board is full or not
func isBoardFull(board [3][3]string) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}