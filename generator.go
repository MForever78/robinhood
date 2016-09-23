package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var dictionary []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildDictionary() {
	f, e := os.Open("dictionary.txt")
	check(e)
	defer func() {
		e := f.Close()
		check(e)
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dictionary = append(dictionary, scanner.Text())
	}

}

func main() {
	buildDictionary()

	var n int
	fmt.Scanf("%d", &n)

	r := rand.New(rand.NewSource(233))
	for i := 0; i < n; i++ {
		op := r.Intn(3)
		index := r.Intn(len(dictionary))
		// 0: insert
		// 1: remove
		// 2: query
		if op == 0 {
			valueIndex := r.Intn(len(dictionary))
			fmt.Printf("%d %s %s\n", op, dictionary[index], dictionary[valueIndex])
		} else {
			fmt.Printf("%d %s\n", op, dictionary[index])
		}
	}
}
