package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, error := os.Open("input01.txt")
	check(error)

	var a_list []int
	var b_list []int
	similarity_map := make(map[int]int)
	similarity_score := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := scanner.Text()

		values := strings.Fields(data)
		a_num, a_err := strconv.Atoi(values[0])
		check(a_err)
		b_num, b_err := strconv.Atoi(values[1])
		check(b_err)

		a_list = append(a_list, a_num)
		b_list = append(b_list, b_num)
	}

	file.Close()

	if len(a_list) != len(b_list) {
		panic("Bad Lists")
	}

	for i := 0; i < len(a_list); i++ {
		target_value := a_list[i]
		count := 0

		for _, value := range b_list {
			if value == target_value {
				count++
			}
		}

		_, ok := similarity_map[a_list[i]]

		if ok == false {
			similarity_map[a_list[i]] = count
		}
	}

	fmt.Println(similarity_map)

	for key, value := range similarity_map {
		similarity_score = similarity_score + (key * value)
	}

	fmt.Println(similarity_score)
}
