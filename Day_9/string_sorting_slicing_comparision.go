package main
import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"
	//"encoding/csv"
	"sort"
	//"math"
)

type low_points struct{
	value int
	key int
	idx int
}

func ReadInput(r string) map[int][]int{
	var result = make(map[int][]int)
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
		str_arr := strings.Split(x, "")
		for _, i := range str_arr {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			result[count] = append(result[count], j)
		}
		count +=1
	}

	return result

}


func CheckAllFour(i int, idx int, v[]int, upper_arr[]int,lower_arr[]int) bool{
	switch {
	case i > v[idx-1]:
		break
	case i > v[idx+1]:
		break
	case i > upper_arr[idx]:
		break
	case i > lower_arr[idx]:
		break
	default:
		return true;
	}

	return false;
}

func RowBasedCheck(i int, idx int, v[]int, arr[]int) bool{
	switch {
	case i > v[idx-1]:
		break
	case i > v[idx+1]:
		break
	case i > arr[idx]:
		break
	default:
		return true;
	}

	return false;
}

func CheckEdge(i int, idx int, v[]int, upper_arr[]int,lower_arr[]int) bool{
	if idx == 0{
		switch {
		case i > v[idx+1]:
			break
		case i > upper_arr[idx]:
			break
		case i > lower_arr[idx]:
			break
		default:
			return true;
		}
	} else {
		switch {
		case i > v[idx-1]:
			break
		case i > upper_arr[idx]:
			break
		case i > lower_arr[idx]:
			break
		default:
			return true;
		}
	}

	return false;
	
}


func RowBasedCheckEdge(i int, idx int, v[]int, arr[]int) bool{
	if idx == 0{
		switch {
		case i > v[idx+1]:
			break
		case i > arr[idx]:
			break
		default:
			return true;
		}
	} else {
		switch {
		case i > v[idx-1]:
			break
		case i > arr[idx]:
			break
		default:
			return true;
		}
	}
	return false;
}

func PuzzleOne(dict map[int][]int) map[int]low_points{
	var risk, count int
	lp_dict := map[int]low_points{}
	
	for k, v := range dict {
		if k == 0 || k == (len(dict) -1){
			var arr []int
			if k == 0{
				arr = dict[k+1]
			} else {
				arr = dict[k-1]
			}
			for idx, i := range v {
				if idx == 0 || idx == (len(v) - 1){
					if RowBasedCheckEdge(i,idx,v,arr){
						if i != 9 {
							lp_dict[count] = low_points{value: i,key: k, idx: idx}
							risk += (i + 1)
							count += 1
						}
					}
				} else {
					if RowBasedCheck(i,idx,v,arr){
						if i != 9 {
							lp_dict[count] = low_points{value: i,key: k, idx: idx}
							risk += (i + 1)
							count += 1
						}
					}
				}
			}
		} else {
			lower_arr := dict[k+1] 
			upper_arr := dict[k-1]
			for idx, i := range v {
				if idx == 0 || idx == (len(v) - 1){
					if CheckEdge(i,idx,v,upper_arr,lower_arr){
						if i != 9 {
							lp_dict[count] = low_points{value: i,key: k, idx: idx}
							risk += (i + 1)
							count += 1
						}
					}
				} else {
					if CheckAllFour(i,idx,v,upper_arr,lower_arr){
						if i != 9 {
							lp_dict[count] = low_points{value: i,key: k, idx: idx}
							risk += (i + 1)
							count += 1
						}
					}
				}
			}
		}
				 
	}
	fmt.Printf("Puzzle 1 answer is: %d\n", risk)

	return lp_dict;
}

// Changed tact and gone for basin mapping - I'm probably missing a trick here
// I'm missing an edge cases here as I don't get the full 14
func PuzzleTwo(lp_dict map[int]low_points, dict map[int][]int){
	var i, j, k int
	var basin_sizes []int
	
	for i = 0; i < len(lp_dict); i++{
		//fmt.Printf("Itteration %d\n",i)
		var basin_size []int
		value := lp_dict[i].value
		key := lp_dict[i].key
		idx := lp_dict[i].idx
		starting_row := dict[key]
		basin_size = append(basin_size, value)

		for k=(key+1); k < len(dict); k++{
			next_row := dict[k]
			if next_row[idx] != 9{
				basin_size = append(basin_size, next_row[idx])
				//fmt.Printf("Row %d, column %d, value %d\n",k, idx, next_row[idx])
			} else{
				break
			}
		}

		for k=(key-1); k >= 0; k--{
			previous_row := dict[k]
			if previous_row[idx] != 9{
				basin_size = append(basin_size, previous_row[idx])
				//fmt.Printf("Row %d, column %d, value %d\n",k, idx, previous_row[idx])
			} else{
				break
			}
		}

		for j=(idx+1); j < len(starting_row); j++{
			if starting_row[j] != 9{
				basin_size = append(basin_size, starting_row[j])
				//fmt.Printf("Row %d, column %d, value %d\n",key, j, starting_row[j])
				for k=(key+1); k < len(dict); k++{
					next_row := dict[k]
					if next_row[j] != 9{
						basin_size = append(basin_size, next_row[j])
						//fmt.Printf("Row %d, column %d, value %d\n",k, j, next_row[j])
					} else{
						break
					}
				}

				for k=(key-1); k >= 0; k--{
					previous_row := dict[k]
					if previous_row[j] != 9{
						basin_size = append(basin_size, previous_row[j])
						//fmt.Printf("Row %d, column %d, value %d\n",k, j, previous_row[j])
					} else{
						break
					}
				}

			} else {
				break
			}

		}

		for j=(idx-1); j >= 0; j-- {
			if starting_row[j] != 9{
				basin_size = append(basin_size, starting_row[j])
				//fmt.Printf("Row %d, column %d, value %d\n",k, j, starting_row[j])
				for k=(key+1); k < len(dict); k++{
					next_row := dict[k]
					if next_row[j] != 9{
						basin_size = append(basin_size, next_row[j])
						//fmt.Printf("Row %d, column %d, value %d\n",k, j, next_row[j])
					} else{
						break
					}
				}

				for k=(key-1); k >= 0; k--{
					previous_row := dict[k]
					if previous_row[j] != 9{
						basin_size = append(basin_size, previous_row[j])
						//fmt.Printf("Row %d, column %d, value %d\n",k, j, previous_row[j])
					} else{
						break
					}
				}
			} else {
				break
			}

		}

		//fmt.Printf("Basin size:")
		//fmt.Println(basin_size)
		//fmt.Printf("Basin size is %d:\n", len(basin_size))
		basin_sizes = append(basin_sizes, len(basin_size))
	}

	sort.Ints(basin_sizes[:])
	for i=1; i < 4; i++{
		fmt.Println(basin_sizes[(len(basin_sizes) -i)])
	}

}

func main() {

	dict := ReadInput("test_input.txt")

	lp_dict := PuzzleOne(dict)

	//fmt.Println(lp_dict)

	PuzzleTwo(lp_dict, dict)

}