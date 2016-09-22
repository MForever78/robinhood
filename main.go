package main

import "fmt"

var tableSize = 333331

type hashNode struct {
	key, value string
	hashValue  uint
}

type hashTable struct {
	nodes []hashNode
}

func (t *hashTable) create(s uint) {
	t.nodes = make([]hashNode, s)
}

func (t *hashTable) insert(key, value string) {
	hashValue = hash(key)
	hashValue = hashValue%tableSize + 1

	initPos := hashValue
	if t.nodes[initPos].hashValue == 0 {
		node := t.nodes[initPos]
		node = hashNode{key, value, hashValue}
	} else {

	}
}

func (t *hashTable) remove(key string) (err error) {
}

func (t *hashTable) query(key string) (err error, value string) {
}

func hash(key string) (hashValue uint) {
	hashValue = 5381

	for _, c := range key {
		hashValue = ((hashValue << 5) + hashValue) + c
	}
}

func main() {
	fmt.Println("vim-go")
}
