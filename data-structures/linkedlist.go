package main

import "fmt"

//Node struct
type Node struct {
	key  string
	next *Node
}

//LinkedList struct
type LinkedList struct {
	head   *Node
	length int
}

//Add element
func (ll *LinkedList) insert(k string) {
	node := &Node{key: k}
	node.next, ll.head = ll.head, node
	ll.length++
}

//Remove element
func (ll *LinkedList) remove(k string) {
	if ll.length == 0 {
		return
	}
	if ll.find(k) {
		if ll.head.key == k {
			ll.head = ll.head.next
			ll.length--
			return
		}

		temp := ll.head
		for temp.next.key != k {
			if temp.next.next == nil {
				return
			}
			temp = temp.next
		}
		temp.next = temp.next.next
		ll.length--
	} else {
		fmt.Println("element not found")
	}
}

//Find element
func (ll *LinkedList) find(k string) bool {

	if ll.head.key == k {
		return true
	}
	previous := ll.head.next
	for previous != nil {
		if previous.key == k {
			return true
		}
		previous = previous.next
	}
	return false

}

func main() {
	n := Node{key: "first"}
	ll := &LinkedList{head: &n, length: 1}
	ll.insert("second")
	ll.insert("third")
	ll.insert("fourth")
	ll.insert("fifth")
	fmt.Println("length :: ", ll.length)
	fmt.Println(ll.find("fourth"))
	ll.remove("fourth")
	fmt.Println(ll.find("fourth"))
	ll.remove("fourth")
	fmt.Println("length :: ", ll.length)
}
