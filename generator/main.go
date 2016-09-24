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

	var inserts, deletes, queries int
	fmt.Scanf("%d %d %d", &inserts, &deletes, &queries)

	r := rand.New(rand.NewSource(233))
	for i := 0; i < inserts; i++ {
		op := 0
		index := i
		valueIndex := r.Intn(len(dictionary))
		fmt.Printf("%d %s %s\n", op, dictionary[index], dictionary[valueIndex])
	}

	for i := 0; i < deletes; i++ {
		op := 1
		index := r.Intn(len(dictionary))
		fmt.Printf("%d %s\n", op, dictionary[index])
	}

	for i := 0; i < queries; i++ {
		op := 2
		index := r.Intn(len(dictionary))
		fmt.Printf("%d %s\n", op, dictionary[index])
	}
}
