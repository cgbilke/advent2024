package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	total_diff := 0

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

	sort.Ints(a_list)
	sort.Ints(b_list)

	if len(a_list) != len(b_list) {
		panic("Bad Lists")
	}

	for i := 0; i < len(a_list); i++ {
		a_id := a_list[i]
		b_id := b_list[i]

		diff := b_id - a_id

		if diff < 0 {
			diff *= -1
		}

		total_diff += diff
	}

	fmt.Println(total_diff)
}
