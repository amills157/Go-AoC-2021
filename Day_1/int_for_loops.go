package main
import (
    "bufio"
    "fmt"
	"os"
	"log"
    "strconv"
)

func ReadInts(r string) ([]int, error) {
	file, err := os.Open(r)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}

func puzzleOne(arr []int) int {
	var i, count int

	count = 0

	for i = 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			count += 1
		}
	}
	return count;
}

func puzzleTwo(arr []int) []int {
	var i, sum int

	array_sums:= make([]int, 0)

	for i = 0; i < len(arr); i++ {
		if i+2 < len(arr){
			sum = arr[i] + arr[i+1] + arr[i+2]
			array_sums = append(array_sums, sum)
		}
	}

	return 	array_sums;
}

func main() {
	var increase_count, increase_count_2 int

    puzzle_one_array, err := ReadInts("puzzle_1.txt")
	if err != nil {
        log.Fatal(err)
    }
	
    increase_count = puzzleOne(puzzle_one_array)

	puzzle_two_array := puzzleTwo(puzzle_one_array)

	/* output the returned value */
	fmt.Printf( "Puzzle 1 answer is: %d\n", increase_count );

	increase_count_2 = puzzleOne(puzzle_two_array)

	fmt.Printf( "Puzzle 2 answer is: %d\n", increase_count_2 );


}