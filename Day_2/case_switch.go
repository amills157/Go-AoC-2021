package main
import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"
)

// This one was simply enough to do live 
// Storing values to re-iterate would have been work for no real gain here
func Both_Puzzles(r string) {
	var split_str []string
	var horizontal_pos_1, horizontal_pos_2, forward_pos int

	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		x := scanner.Text()
		split_str = strings.Split(x, " ")
		value, err := strconv.Atoi(split_str[1])
		if err != nil {
			return
		}
		
		switch split_str[0] {
		case "forward":
			forward_pos += value
			horizontal_pos_2 += horizontal_pos_1*value
		case "up":
			horizontal_pos_1 -= value
		case "down":
			horizontal_pos_1 += value
		default:
			fmt.Println("Here be dragons")
		}
	}

	fmt.Printf( "Puzzle 1 answer is: %d\n", forward_pos * horizontal_pos_1 );

	fmt.Printf( "Puzzle 2 answer is: %d\n", forward_pos * horizontal_pos_2 );
	
	return
}

/*
func Puzzle_2(r string) {
	var split_str []string
	var horizontal_pos, forward_pos, aim int

	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		x := scanner.Text()
		split_str = strings.Split(x, " ")
		value, err := strconv.Atoi(split_str[1])
		if err != nil {
			return
		}
		
		switch split_str[0] {
		case "forward":
			forward_pos += value
			horizontal_pos += aim*value
		case "up":
			aim -= value
		case "down":
			aim += value
		default:
			fmt.Println("Here be dragons")
		}
	}

	fmt.Printf( "Puzzle 2 answer is: %d\n", forward_pos * horizontal_pos );
	
	return
}
*/

func main() {

	Both_Puzzles("puzzle_1.txt")

}