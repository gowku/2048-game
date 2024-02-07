package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/eiannone/keyboard"
)
const boardSize = 4
var errEndGame = errors.New("GameOverError")
var isGameOver = false
const ProbSeuil2 = 70
var emptycellsarr = [][2]int{}


type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
	NO_DIR
)

// GetCharKeystroke returns the key pressed by the user
func GetCharKeystroke() (Dir, error) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	char, key, err := keyboard.GetKey()
	ans := int(char)
	if ans == 0 {
		ans = int(key)
	}

	// fmt.Printf("the key is: %v \n", ans)

	if err != nil {
		return NO_DIR, err
	}
	switch ans {
	case 119, 65517, 107:
		return UP, nil
	case 97, 65515, 104:
		return LEFT, nil
	case 115, 65516, 106:
		return DOWN, nil
	case 100, 65514, 108:
		return RIGHT, nil
	case 3:
		// return NO_DIR, errEndGame
		return NO_DIR, err
	}
	return NO_DIR, nil
}

// PrintBoard prints the 2048 game board
func PrintBoard(board [boardSize][boardSize]int) {
	fmt.Println("----------------------------------------")
	for i := 0; i < boardSize; i++ {
		fmt.Print("|")
		for j := 0; j < boardSize; j++ {
			fmt.Printf("%4s", "")
			fmt.Printf("%-4d |", board[i][j])
		}
		fmt.Println("\n----------------------------------------")
	}
}

func KeystrokeReact (arr [4]*int) {
	for i := 0; i < 3; i++ {
		switch {
		case *arr[i + 1] == 0:
			*arr[i + 1] = *arr[i]
			*arr[i] = 0
		case *arr[i] == *arr[i + 1]:
			*arr[i + 1] = *arr[i] + *arr[i + 1]
			*arr[i] = 0			
		}
	}
	for j := 0; j < 3; j++ {
		if *arr[j + 1] == 0 {
			*arr[j + 1] = *arr[j]
			*arr[j] = 0
		}
	}
}

func GetEmptyCell(board [boardSize][boardSize]int) [][2]int {
	// emptycellsarr := [][2]int{}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == 0 {
				emptycellsarr = append(emptycellsarr, [2]int{i, j})
				
			}
		}
	}
	return emptycellsarr
}

func AddRandom(board *[boardSize][boardSize]int, emptycellsarr [][2]int){
	if len(emptycellsarr) > 0 {
		randomIndex := rand.Intn(len(emptycellsarr))
		row, col := emptycellsarr[randomIndex][0], emptycellsarr[randomIndex][1]
		
		// Randomly choose between 2 and 4
		randomValue := 2
		if rand.Intn(2) == 1 {
			randomValue = 4
		}
		(*board)[row][col] = randomValue
	}
	}

func main() {
	fmt.Printf("Utiliser {w, a, s, d} pour jouer\n")
	fmt.Printf("Appuyer sur n´importe quelle touche pour commencer\n")
	_,err := GetCharKeystroke()
	if err != nil {
			fmt.Println("Error: ", err)
		}
	
	_clearScreenSequence := "\033[H\033[2J"
	
	board := [boardSize][boardSize]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	
	L10,L11,L12,L13 := &board[0][0],&board[0][1],&board[0][2],&board[0][3]
	L20,L21,L22,L23 := &board[1][0],&board[1][1],&board[1][2],&board[1][3]
	L30,L31,L32,L33 := &board[2][0],&board[2][1],&board[2][2],&board[2][3]
	L40,L41,L42,L43 := &board[3][0],&board[3][1],&board[3][2],&board[3][3]
	
	ligne1, rev_ligne1 := [4]*int{L10, L11, L12, L13}, [4]*int{L13, L12, L11, L10}
	ligne2, rev_ligne2 := [4]*int{L20, L21, L22, L23}, [4]*int{L23, L22, L21, L20}
	ligne3, rev_ligne3 := [4]*int{L30, L31, L32, L33}, [4]*int{L33, L32, L31, L30}
	ligne4, rev_ligne4 := [4]*int{L40, L41, L42, L43}, [4]*int{L43, L42, L41, L40}
	
	col1, rev_col1 := [4]*int{L10, L20, L30, L40}, [4]*int{L40, L30, L20, L10}
	col2, rev_col2 := [4]*int{L11, L21, L31, L41}, [4]*int{L41, L31, L21, L11}
	col3, rev_col3 := [4]*int{L12, L22, L32, L42}, [4]*int{L42, L32, L22, L12}
	col4, rev_col4 := [4]*int{L13, L23, L33, L43}, [4]*int{L43, L33, L23, L13}
	
	fmt.Println(_clearScreenSequence)
	GetEmptyCell(board)
	AddRandom(&board, emptycellsarr)
	PrintBoard(board)
	
	for true {
		if isGameOver {
			break
		}
		dir, err := GetCharKeystroke()
		if err != nil {
			fmt.Println("Error: ", err)
		}
		switch dir {
		case UP:
			KeystrokeReact(rev_col1)
			KeystrokeReact(rev_col2)
			KeystrokeReact(rev_col3)
			KeystrokeReact(rev_col4)
		case DOWN:
			KeystrokeReact(col1)
			KeystrokeReact(col2)
			KeystrokeReact(col3)
			KeystrokeReact(col4)
		case LEFT:
			KeystrokeReact(rev_ligne1)
			KeystrokeReact(rev_ligne2)
			KeystrokeReact(rev_ligne3)
			KeystrokeReact(rev_ligne4)	
		case RIGHT:
			KeystrokeReact(ligne1)
			KeystrokeReact(ligne2)
			KeystrokeReact(ligne3)
			KeystrokeReact(ligne4)
		}
		fmt.Println(_clearScreenSequence)
		AddRandom(&board, emptycellsarr)
		PrintBoard(board)
	}
		
		// fmt.Println("ligne1[0]: %v,ligne1[1]: %v,ligne1[2]: %v,ligne1[3]: %v", ligne1[0],ligne1[1],ligne1[2],ligne1[3])
		// fmt.Println("ligne2[0]: %v,ligne2[1]: %v,ligne2[2]: %v,ligne2[3]: %v", ligne2[0],ligne2[1],ligne2[2],ligne2[3])
		// fmt.Println("ligne3[0]: %v,ligne3[1]: %v,ligne3[2]: %v,ligne3[3]: %v", ligne3[0],ligne3[1],ligne3[2],ligne3[3])
		// fmt.Println("ligne4[0]: %v,ligne4[1]: %v,ligne4[2]: %v,ligne4[3]: %v", ligne4[0],ligne4[1],ligne4[2],ligne4[3])
		
		// fmt.Println("col1[0]: %v,col1[1]: %v,col1[2]: %v,col1[3]: %v", col1[0],col1[1],col1[2],col1[3])
		// fmt.Println("col2[0]: %v,col2[1]: %v,col2[2]: %v,col2[3]: %v", col2[0],col2[1],col2[2],col2[3])
		// fmt.Println("col3[0]: %v,col3[1]: %v,col3[2]: %v,col3[3]: %v", col3[0],col3[1],col3[2],col3[3])
		// fmt.Println("col4[0]: %v,col4[1]: %v,col4[2]: %v,col4[3]: %v", col4[0],col4[1],col4[2],col4[3])
		
		// fmt.Println("rev_ligne1[0]: %v,rev_ligne1[1]: %v,rev_ligne1[2]: %v,rev_ligne1[3]: %v", rev_ligne1[0],rev_ligne1[1],rev_ligne1[2],rev_ligne1[3])
		// fmt.Println("rev_ligne2[0]: %v,rev_ligne2[1]: %v,rev_ligne2[2]: %v,rev_ligne2[3]: %v", rev_ligne2[0],rev_ligne2[1],rev_ligne2[2],rev_ligne2[3])
		// fmt.Println("rev_ligne3[0]: %v,rev_ligne3[1]: %v,rev_ligne3[2]: %v,rev_ligne3[3]: %v", rev_ligne3[0],rev_ligne3[1],rev_ligne3[2],rev_ligne3[3])
		// fmt.Println("rev_ligne4[0]: %v,rev_ligne4[1]: %v,rev_ligne4[2]: %v,rev_ligne4[3]: %v", rev_ligne4[0],rev_ligne4[1],rev_ligne4[2],rev_ligne4[3])
		
		// fmt.Println("rev_col1[0]: %v,rev_col1[1]: %v,rev_col1[2]: %v,rev_col1[3]: %v", rev_col1[0],rev_col1[1],rev_col1[2],rev_col1[3])
		// fmt.Println("rev_col2[0]: %v,rev_col2[1]: %v,rev_col2[2]: %v,rev_col2[3]: %v", rev_col2[0],rev_col2[1],rev_col2[2],rev_col2[3])
		// fmt.Println("rev_col3[0]: %v,rev_col3[1]: %v,rev_col3[2]: %v,rev_col3[3]: %v", rev_col3[0],rev_col3[1],rev_col3[2],rev_col3[3])
		// fmt.Println("rev_col4[0]: %v,rev_col4[1]: %v,rev_col4[2]: %v,rev_col4[3]: %v", rev_col4[0],rev_col4[1],rev_col4[2],rev_col4[3])
}