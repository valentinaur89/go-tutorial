package main

import (
	"fmt"
	"os"
)

const (
	arrsize = 7
)

//HashTable hashtable
type HashTable struct {
	arr [arrsize]*bucket
}
type bucket struct {
	head *bucketNode
}
type bucketNode struct {
	key  string
	next *bucketNode
}

func (ht *HashTable) put(k string) {
	index := hash(k)
	ht.arr[index].insert(k)

}
func (ht *HashTable) delete(k string) {
	index := hash(k)
	ht.arr[index].remove(k)
}

func (b *bucket) remove(k string) {
	if b.search(k) {
		if b.head.key == k {
			b.head = b.head.next
			return
		}
		for b.head.next != nil {
			if b.head.next.key == k {
				b.head = b.head.next.next
				return
			}
			b.head = b.head.next
		}

	} else {
		fmt.Println("element doesnt exists")
	}
}
func (b *bucket) insert(k string) {
	node := &bucketNode{key: k}
	if !b.search(k) {
		node.next = b.head
		b.head = node
	} else {
		fmt.Println("already exist")
	}
}
func (b *bucket) search(k string) bool {
	if b.head != nil {
		if b.head.key == k {
			return true
		}
		node := b.head.next
		for node != nil {
			if node.key == k {
				return true
			}
			node = node.next
		}
	}
	return false
}
func (ht *HashTable) find(k string) bool {
	index := hash(k)
	return ht.arr[index].search(k)
}

func hash(k string) int {
	if k != "" {
		sum := 0
		for s := range k {
			sum += int(s)
		}
		return sum % arrsize
	}
	return -1
}

func initiate() *HashTable {
	ht := &HashTable{}
	for i := range ht.arr {
		ht.arr[i] = &bucket{}
	}
	return ht
}

func main() {
	ht := initiate()
	ht.put("first")
	ht.put("second")
	ht.put("third")
	ht.put("fourth")
	// go run hashtable.go second second
	fmt.Println(ht.find(os.Args[1]))
	ht.delete(os.Args[2])
	fmt.Println(ht.find(os.Args[1]))

}
