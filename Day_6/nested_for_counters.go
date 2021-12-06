package main
import (
	//"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	//"strings"
	"encoding/csv"
)

func ReadCSVInput(r string) []int{
	//var result = make(map[int]string)
	var input_data []int

	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	scanner := csv.NewReader(file)
    records, _ := scanner.ReadAll()

	// It returns it nested- No idea why
	for _, i := range records[0] {
		j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        input_data = append(input_data, j)
    }

	return input_data;
}

func PuzzleOne(arr []int){
	var i, j int

	for i = 0; i < 80; i++ {
		for j = 0; j < len(arr); j++ {
			if arr[j] == 0{
				arr[j] = 6
				arr = append(arr, 9)
			} else {
				arr[j] = arr[j] - 1
			}
		}
	}

	fmt.Printf("Puzzle 1 answer is: %d\n", len(arr))

}	

// Tried to doing this with a map - Got lost in the woods amongst the trees
func BothPuzzles(arr []int){
	var i, total int
	
	counter := []int{0,0,0,0,0,0,0,0,0}

	for _, i := range arr {
		counter[i] +=1
    }

	for i = 0; i < 256; i++{
		new_fish := counter[0]
		// Shift the count down one - Day 3 value becomes Day 2 etc
		counter = counter[1:]
		counter[6] += new_fish 
		counter = append(counter, new_fish)

		if i == 79{
			for _, i := range counter{
				total += i
			}
			fmt.Printf("Puzzle 1 answer is: %d\n", total)
		}
	}  

	total = 0

	for _, i := range counter{
		total += i
	}

	fmt.Printf("Puzzle 2 answer is: %d\n", total)
}	

// Keeping PuzzleOne just because
func main() {

	arr := ReadCSVInput("puzzle_1.txt")

	//PuzzleOne(arr)

	BothPuzzles(arr)

}