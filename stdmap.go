package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
)

var tableSize uint = 3333331
var table map[string]string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		check(err)
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	table := make(map[string]string, tableSize)

	f, err := os.Open("in")
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	for {
		var op int
		var key string
		var value string
		_, err := fmt.Fscan(f, &op, &key)
		if err != nil {
			if err == io.EOF {
				return
			} else {
				panic(err)
			}
		}

		switch op {
		case 0:
			fmt.Fscan(f, &value)
			table[key] = value
			fmt.Println("inserted", key)
		case 1:
			if table[key] == "" {
				fmt.Printf("cannot delete %s\n", key)
			} else {
				table[key] = ""
				fmt.Println("deleted", key)
			}
		case 2:
			value := table[key]
			if value == "" {
				fmt.Printf("cannot find %s\n", key)
			} else {
				fmt.Println(value)
			}
		}
	}
}
