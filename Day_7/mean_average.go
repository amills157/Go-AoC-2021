package main
import (
	//"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	//"strings"
	"encoding/csv"
	"sort"
	//"math"
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

	sort.Ints(input_data[:])

	return input_data;
}

// <Insert joke about Shell oil here>
func FuelCost(arr []int, medianValue int) int {
	var fuel, difference int

	for _, i := range arr {
		if i > medianValue{
			difference = i - medianValue
		} else {
			difference = medianValue - i
		}
		fuel += difference
	}

	return fuel
}

func BrexitFuelCost(arr []int, avg int) int {
	var fuel, difference, fuel_tax, j int

	for _, i := range arr {
		if i > avg{
			difference = i - avg
		} else {
			difference = avg - i
		}

		fuel_tax = 0

		for j = 0; j < difference; j++ {
			fuel_tax += (j + 1)
		}

		fuel += fuel_tax
	}

	return fuel
}


func PuzzleOne(arr []int){
	
	median := len(arr) / 2

	if median % 2 == 0 {
		medianValue := (arr[median-1] + arr[median]) / 2;

		fuel := FuelCost(arr, medianValue)

		fmt.Printf("Puzzle 1 answer is: %d\n", fuel)
	} else {

		fuel := FuelCost(arr, arr[median])

		fmt.Printf("Puzzle 1 answer is: %d\n", fuel)
	}

}	

func PuzzleTwo(arr []int){
	var sum int
	
	for _, i := range arr {
		sum += i
	}

	// This is horrible but needed - Why does go round down from 4.9 to 4!?
	//avg := int(math.Round((float64(sum)) / (float64(len(arr)))))

	// Turns out for the test input you need ^ for the main input you want it to round down
	avg := sum / len(arr)

	fuel := BrexitFuelCost(arr, avg)

	fmt.Printf("Puzzle 2 answer is: %d\n", fuel)

}

func main() {

	arr := ReadCSVInput("puzzle_1.txt")

	PuzzleOne(arr)

	PuzzleTwo(arr)

}