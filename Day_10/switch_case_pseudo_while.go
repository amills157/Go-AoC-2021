package main
import (
	"bufio"
	"fmt"
	"os"
	"log"
	//"strconv"
	"strings"
	//"sort"
	//"math"
)

type chunk_location struct{
	open int
	close int
}


func ReadInput(r string) map[int][]string{
	var result = make(map[int][]string)
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
		split_str := strings.Split(x, "")

		result[count] = split_str

		count +=1
	}

	return result

}

func RemoveIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}

func PuzzleOne(dict map[int][]string){
	var answer int
	var points = make(map[int]int)
	
	for _,v := range dict{
		arr := v
		var next_closed string
		var open_idx, closed_idx int
		
		for {

			//fmt.Println(arr)

			check_str := strings.Join(arr,"")

			if !strings.ContainsAny(check_str, ")]}>"){
				goto NEXT
			}

			for idx,bracket := range arr{
				
				switch bracket{
				case "(":
					next_closed = ")"
					open_idx = idx
				case "[":
					next_closed = "]"
					open_idx = idx
				case "{":
					next_closed = "}"
					open_idx = idx
				case "<":
					next_closed = ">"
					open_idx = idx
				default:
					if bracket != next_closed{
						//fmt.Printf("ERROR! Expected %s but got %s\n", next_closed, bracket)
						switch bracket{
							case ")":
								points[3] += 3
							case "]":
								points[57] += 57
							case "}":
								points[1197] += 1197
							case ">":
								points[25137] += 25137
							default:
						}
						goto NEXT
					} else {
						closed_idx = idx
						goto RERUN
					}
				}
			
			}

			RERUN:
			arr = RemoveIndex(arr, open_idx)
			arr = RemoveIndex(arr, (closed_idx-1))
	
		}


		NEXT:
	        
    }
	//fmt.Println(points)

	for _,v := range points{
		answer += v
	}

	fmt.Printf("Puzzle 1 answer is: %d\n", answer)

}	


func main() {

	dict := ReadInput("puzzle_1.txt")

	PuzzleOne(dict)

	//PuzzleTwo(dict)

}