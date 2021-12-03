package main
import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"
)

func ReadInput(r string) (map[int]string) {
	var result = make(map[int]string)
	var count int

	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		x := scanner.Text()

		result[count] = x

		count += 1 
	}
	return result
}


func puzzleOne(items map[int]string) {
	var i int
	var gamma []string
	var epsilon []string

	for i = 0; i < len(items[0]); i++ {
		var count_0, count_1 int
		for _, value := range items {
			first_bit, err := strconv.Atoi(string(value[i]))
			if err != nil {
				return;
			}
			if first_bit >= 1{
				count_1 += 1
			} else {
				count_0 += 1
			}
		}

		if count_0 > count_1{
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		} else {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		}
		count_0 = 0
		count_1 = 0
	}

	gamma_dec, err := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)  
	if err != nil {   
		return  
	}

	epsilon_dec, err := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)  
	if err != nil {   
		return  
	}

	fmt.Printf( "Puzzle 1 answer is: %d\n", gamma_dec * epsilon_dec );

	return 
}

func puzzleTwo_OxMap(items map[int]string) {
	var i, j int
	var o2_map = make(map[int]string)

	for k,v := range items {
		o2_map[k] = v
	}
	
	out:
	for i = 0; i < len(items[0]); i++ {
		var count_0, count_1 []int
		for key, value := range o2_map {
			first_bit, err := strconv.Atoi(string(value[i]))
			if err != nil {
				return ;
			}
			if first_bit >= 1{
				count_1 = append(count_1, key)
			} else {
				count_0 = append(count_0, key)
			}
		}

		if len(count_0) > 0 && len(count_1) > 0{
			if len(count_0) > len(count_1){
				for j = 0; j < len(count_1); j++ { 
					delete(o2_map, count_1[j]);
				}
			} else {
				for j = 0; j < len(count_0); j++ { 
					delete(o2_map, count_0[j]);
				}
			}
		} else {
			break out
		}

	}

	for _, value := range o2_map {
		o2_value, err := strconv.ParseInt(value, 2, 64)  
		if err != nil {   
			return ;
		}
		fmt.Printf( "O2 Value is: %d\n", o2_value );
	}

	return ;
}

func puzzleTwo_Co2Map(items map[int]string) {
	var i, j int
	var co2_map = make(map[int]string)

	for k,v := range items {
		co2_map[k] = v
	}
	
	out:
	for i = 0; i < len(items[0]); i++ {
		var count_0, count_1 []int
		for key, value := range co2_map {
			first_bit, err := strconv.Atoi(string(value[i]))
			if err != nil {
				return ;
			}
			if first_bit >= 1{
				count_1 = append(count_1, key)
			} else {
				count_0 = append(count_0, key)
			}
		}

		if len(count_0) > 0 && len(count_1) > 0{
			if len(count_0) > len(count_1){
				for j = 0; j < len(count_0); j++ { 
					delete(co2_map, count_0[j]);
				}
			} else {
				for j = 0; j < len(count_1); j++ { 
					delete(co2_map, count_1[j]);
				}
			}
		} else {
			break out
		}

	}

	for _, value := range co2_map {
		co2_value, err := strconv.ParseInt(value, 2, 64) 
		if err != nil {   
			return ;
		}
		fmt.Printf( "CO2 Value is: %d\n", co2_value );
	}

	return ;
}

// Need to sort the return typ confusion / casting for these as right now I'm doing the math
// manually - Rush job!
func main() {

	puzzle_one_dict := ReadInput("puzzle_1.txt")

	puzzleOne(puzzle_one_dict)	

	puzzleTwo_OxMap(puzzle_one_dict)

	puzzleTwo_Co2Map(puzzle_one_dict)
	

}