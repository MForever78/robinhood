package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
)

var tableSize uint = 3333331

type hashNode struct {
	key, value string
	hashValue  uint
}

type hashTable struct {
	nodes    []hashNode
	capacity uint
}

func (t *hashTable) String() string {
	return fmt.Sprint(t.nodes)
}

func (t *hashTable) create(s uint) {
	// hash value zero means empty bucket
	t.nodes = make([]hashNode, s+1)
	t.capacity = s
}

func (t *hashTable) dib(index uint) (distance uint) {
	distance = index - t.nodes[index].hashValue
	if distance < 0 {
		distance += uint(len(t.nodes))
	}
	return distance
}

func (t *hashTable) incPos(index uint) (pos uint) {
	pos = index + 1
	if pos >= uint(len(t.nodes)) {
		pos = 1
	}
	return pos
}

func (t *hashTable) insert(key, value string) {
	hashValue := hash(key)
	hashValue = hashValue%t.capacity + 1

	initPos := hashValue
	newNode := hashNode{key, value, hashValue}

	pos := initPos
	for t.nodes[pos].hashValue != 0 {
		if t.dib(pos) < pos-initPos {
			tempNode := newNode
			newNode = t.nodes[pos]
			t.nodes[pos] = tempNode
		}
		pos = t.incPos(pos)
	}

	t.nodes[pos] = newNode
}

func (t *hashTable) remove(key string) (err error) {
	ok, initPos := t.queryIndex(key)

	if ok == nil {
		err = nil
		lastPos := initPos + 1
		for t.nodes[lastPos].hashValue != 0 && t.dib(lastPos) != 0 {
			lastPos = t.incPos(lastPos)
		}

		for pos := initPos; pos != lastPos; pos = t.incPos(pos) {
			t.nodes[pos] = t.nodes[pos+1]
		}
	} else {
		err = ok
	}

	return err
}

func (t *hashTable) queryIndex(key string) (err error, index uint) {
	hashValue := hash(key)
	hashValue = hashValue%t.capacity + 1

	initPos := hashValue
	pos := initPos
	for t.nodes[pos].hashValue != 0 && t.dib(pos) > pos-initPos {
		pos = t.incPos(pos)
	}

	if t.nodes[pos].key == key {
		index = pos
		err = nil
	} else {
		err = errors.New("cannot find key")
	}
	return err, index
}

func (t *hashTable) query(key string) (err error, value string) {
	ok, pos := t.queryIndex(key)

	if ok == nil {
		value = t.nodes[pos].value
		err = nil
	} else {
		err = errors.New("cannot find key")
	}

	return err, value
}

func hash(key string) (hashValue uint) {
	hashValue = 5381

	for _, c := range key {
		hashValue = ((hashValue << 5) + hashValue) + uint(c)
	}
	return hashValue
}

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
	table := new(hashTable)
	table.create(tableSize)

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
			table.insert(key, value)
		case 1:
			err := table.remove(key)
			if err != nil {
				fmt.Printf("cannot delete %s\n", key)
			}
		case 2:
			err, value := table.query(key)
			if err != nil {
				fmt.Printf("cannot find %s\n", key)
			} else {
				fmt.Println(value)
			}
		}
	}
}
