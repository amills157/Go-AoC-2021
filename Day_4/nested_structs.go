package main
import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"
)

type bingoCard struct {
	cells [5][5]cell
	won bool
}

type cell struct {
	value  int
	marked bool
}


func makebingoCard(cardRows [][]int) bingoCard {
	var newCard bingoCard

	for i := range cardRows {
		for j := range cardRows[i]{
			newCard.cells[i][j] = cell{value: cardRows[i][j]}
		}
		
	}

	return newCard
}

func winningBoard(card bingoCard) bool {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !card.cells[i][j].marked {
				break
			// Tried to use goto END here - But it seemed to break it. No idea why
			} 
			if j == 4 {
				return true
			}
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !card.cells[j][i].marked {
				break
			}
			if j == 4 {
				return true
			}
		}
	}

	return false
}

func unmarkedSum(card bingoCard) int  {
	var sum int

	for i := range card.cells {
		for j := range card.cells[i] {
			if !card.cells[i][j].marked {
				sum += card.cells[i][j].value
			}
		}
	}
	return sum
}


func ReadInput(r string) ([]int, []bingoCard) {
	var bingoNumbers []int
	var bingCardRows [][]int
	var bingoCards []bingoCard
	
	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	idx := 0

	for scanner.Scan() {
		x := scanner.Text()
		//Bingo numbers are on the first line 
		if idx == 0 {
			split_str := strings.Split(x, ",")

			for _, i := range split_str {
				j, err := strconv.Atoi(i)
				if err != nil {
					panic(err)
				}
				bingoNumbers = append(bingoNumbers, j)
			}
			idx++

		} else {
			if len(x) != 0 { // add card row to list
				var temp [] int
				str_line := strings.Fields(x)

				for _, i := range str_line {
					j, err := strconv.Atoi(i)
					if err != nil {
						panic(err)
					}
					temp = append(temp, j)
				}

				bingCardRows = append(bingCardRows, temp)
				
				// initially did this using the else, but always missed last card if no trailing line
				if len(bingCardRows) == 5{ 
					newCard := makebingoCard(bingCardRows)
					bingoCards = append(bingoCards, newCard)
					bingCardRows = nil
				}
			}
		}
	}

	return bingoNumbers, bingoCards
}


func bothPuzzles(numbers []int, cards []bingoCard) {
	var answer, lastNum int
	var card bingoCard

	partOneAnswered := -1

	for _, num := range numbers {
		for i := range cards {
			for j := range cards[i].cells {
				for k := range cards[i].cells[j] {
					if cards[i].cells[j][k].value == num {
						cards[i].cells[j][k].marked = true
					}
				}
			}
		}

		for i := range cards {
			if winningBoard(cards[i]) {
				if !cards[i].won{

					card = cards[i]
					lastNum = num
					// Breaks it if we try card.won = true
					cards[i].won = true

					if partOneAnswered == -1{
						answer = unmarkedSum(card)

						fmt.Printf("Puzzle 1 answer is: %d\n", answer*num)

						partOneAnswered = 0
					}
				} 
			}
		}

	}

	answer = unmarkedSum(card)

	fmt.Printf("Puzzle 2 answer is: %d\n", answer*lastNum)

}


func main() {

	numbers, cards := ReadInput("puzzle_1.txt")

	bothPuzzles(numbers, cards)	
	
}