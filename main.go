package main

import (
	"errors"
	"fmt"
)

var tableSize uint = 333331

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

func (t *hashTable) decpos(index uint) (pos uint) {
	pos = index - 1
	if pos == 0 {
		pos = uint(len(t.nodes) - 1)
	}
	return pos
}

//TODO: what if table is full
func (t *hashTable) insert(key, value string) {
	hashValue := hash(key)
	hashValue = hashValue%tableSize + 1

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

		for pos := initPos; pos != lastPos-1; pos = t.incPos(pos) {
			t.nodes[pos] = t.nodes[pos+1]
		}
	} else {
		err = ok
	}

	return err
}

func (t *hashTable) queryIndex(key string) (err error, index uint) {
	hashValue := hash(key)
	hashValue = hashValue%tableSize + 1

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

func main() {
	table := new(hashTable)
	table.create(tableSize + 1)

	table.insert("foo", "bar")
	table.insert("bar", "foo")
	table.insert("foobar", "foo")
	fmt.Println(table.query("foo"))
	table.remove("foo")
	fmt.Println(table.query("foo"))
}
