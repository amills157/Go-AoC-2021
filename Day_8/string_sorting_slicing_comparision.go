package main
import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"
	"sort"
	//"math"
)

type clock_data struct{
	input []string
	output []string
}

// cf
var clock_1 = 2

// acf
var clock_7 = 3

// bcdf
var clock_4 = 4

// acdeg only digit w/o f
var clock_2 = 5
// acdfg = 7+4 (+g -b) & only 5 with 1 & 4 (+ag -b)
var clock_3 = 5
//abdfg = 7+4 (+g -c) & 4 (+ag -c)
var clock_5 = 5

// abcefg = 7+4 (+eg) & 4 (+aeg -d)
var clock_0 = 6
// abdefg = 7+4 (+eg -c) & only 6 without 1 & 4 (+ aeg -c)
var clock_6 = 6
// abcdfg = 7+4 (+g) & 4 (+ag)
var clock_9 = 6

// abcdefg
var clock_8 = 7


func StringSort(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))

	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func StringDiff(str_1 string, str_2 string) string{
	var i int
	var lead_str, short_str, diff_str string

	if len(str_1) > len(str_2){
		lead_str = str_1
		short_str = str_2
	} else {
		lead_str = str_2
		short_str = str_1
	}

	for i=0; i < len(lead_str); i++{
		if !(strings.Contains(short_str, string(lead_str[i]))){
			diff_str += string(lead_str[i])
		}
	}

	return diff_str
}

// Avoids duplicate chars being added
func StringMerge(str_1 string, str_2 string) string{
	var i int
	var lead_str, return_str string

	if len(str_1) > len(str_2){
		lead_str = str_1
		return_str = str_2
	} else {
		lead_str = str_2
		return_str = str_1
	}

	for i=0; i < len(lead_str); i++{
		if !(strings.Contains(return_str, string(lead_str[i]))){
			return_str += string(lead_str[i])
		}
	}

	return_str = StringSort(return_str)

	return return_str
}

func RemoveIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}


func ReadInput(r string) map[int]clock_data{
	var result = make(map[int]clock_data)
	var count int

	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		x := scanner.Text()
		split_str := strings.Split(x, " | ")
		input_arr := strings.Split(split_str[0], " ")
		for i := range input_arr {
			input_arr[i] = StringSort(input_arr[i])
		}
		output_arr := strings.Split(split_str[1], " ")
		for i := range output_arr {
			output_arr[i] = StringSort(output_arr[i])
		}

		temp := clock_data{input: input_arr, output: output_arr}

		result[count] = temp

		count +=1
	}

	return result

}


func PuzzleOne(dict map[int]clock_data){
	var i, answer int
	for _, v := range dict {
        for i=0; i < len(v.output); i++{
			if len(v.output[i]) != 5 && len(v.output[i]) != 6{
				answer += 1
			}
			
		}
    }
	fmt.Printf("Puzzle 1 answer is: %d\n", answer)
}	

func PuzzleTwo(dict map[int]clock_data){
	var i, answer int
	var str_zero, str_one, str_two, str_three, str_four, str_five, str_six, str_seven, str_eight, str_nine, str_sev_for string
	var result_arr []int
	
	for _, v := range dict {
		var reading_arry []string

		// Sort the input by length so we can easily work out the rules
		sort.Slice(v.input, func(x, y int) bool {
			l1, l2 := len(v.input[x]), len(v.input[y])
			if l1 != l2 {
				return l1 < l2
			}
			return v.input[x] < v.input[y]
		})

		// We work out the 5 length cases last - Each case_5 ends up at idx 3 as we shift them to the end
		for j:=0; j <3; j++{
			case_5 := v.input[3]
			v.input = RemoveIndex(v.input, 3)
			v.input = append(v.input, case_5)
		}
		
        for i=0; i < len(v.input); i++{

			switch len(v.input[i]) {
			case clock_1:
				str_one = v.input[i]
			case clock_4:
				str_four = v.input[i]
				// Use this for diff checking later on
				str_sev_for = StringMerge(str_four, str_seven)
			case clock_7:
				str_seven = v.input[i]
			case clock_8:
				str_eight = v.input[i]
			case 5:
				if strings.Contains(v.input[i], string(str_one[0])) && strings.Contains(v.input[i], string(str_one[1])) {
					str_three = v.input[i]
				} else {
					diff_str := StringDiff(v.input[i], str_six)
					if len(diff_str) == 1{
						str_five = v.input[i]
					} else {
						str_two = v.input[i]
					}
				}
			case 6:
				if !(strings.Contains(v.input[i], string(str_one[0]))) || !(strings.Contains(v.input[i], string(str_one[1]))) {
					str_six = v.input[i]
				} else {
					diff_str := StringDiff(v.input[i], str_sev_for)
					if len(diff_str) == 1{
						str_nine = v.input[i]
					} else {
						str_zero = v.input[i]
					}
				}
			default:
				fmt.Println("Here be dragons")
			}

		}

		// Not a huge fan of this, but it will do
		for i=0; i < len(v.output); i++{
			switch v.output[i] {
			case str_zero:
				reading_arry = append(reading_arry, "0")
			case str_one:
				reading_arry = append(reading_arry, "1")
			case str_two:
				reading_arry = append(reading_arry, "2")
			case str_three:
				reading_arry = append(reading_arry, "3")
			case str_four:
				reading_arry = append(reading_arry, "4")
			case str_five:
				reading_arry = append(reading_arry, "5")
			case str_six:
				reading_arry = append(reading_arry, "6")
			case str_seven:
				reading_arry = append(reading_arry, "7")
			case str_eight:
				reading_arry = append(reading_arry, "8")
			case str_nine:
				reading_arry = append(reading_arry, "9")
			default:
				fmt.Println("Here be dragons")
			}
		}

		reading := strings.Join(reading_arry, "")

		c, err := strconv.Atoi(reading)
		if err != nil {
			panic(err)
		}

		result_arr = append(result_arr, c)

    }

	for j := range result_arr{
		answer += result_arr[j]
	}

	fmt.Printf("Puzzle 2 answer is: %d\n", answer)

}

func main() {

	dict := ReadInput("puzzle_1.txt")

	PuzzleOne(dict)

	PuzzleTwo(dict)

}