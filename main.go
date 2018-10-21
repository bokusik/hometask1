package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func bubble(A []int) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A)-1-i; j++ {
			if A[j] > A[j+1] {
				tmp := A[j]
				A[j] = A[j+1]
				A[j+1] = tmp
			}
		}
	}
}
func main() {
	fmt.Printf("Название файла: ")
	var name_file string
	var name_file2 string
	fmt.Scanln(&name_file)
	file, err := ioutil.ReadFile(name_file)
	check(err)
	s := string(file)
	p := strings.Fields(s)
	ary := make([]int, len(p))
	for i := range ary {
		ary[i], _ = strconv.Atoi(p[i])
	}
	bubble(ary)
	t := make([]string, len(ary))
	for i := 0; i < len(ary); i++ {
		t[i] = strconv.Itoa(ary[i])
	}
	fmt.Printf("Название выходного файла: ")
	fmt.Scanln(&name_file2)
	file1, err1 := os.Create(name_file2)
	check(err1)
	for i := 0; i < len(t); i++ {
		file1.WriteString(t[i])
		file1.WriteString(" ")
	}
	fmt.Printf("Содержимое файла: ")
	f, err := ioutil.ReadFile(name_file2)
	check(err)
	fmt.Printf("%s \n", f)
}
