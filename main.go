package main

import (
	"fmt"
	"math/rand"
)

type Key struct {
	key int
	// value   int //add later
	nextNode *Node
}

func NewKey(key int) *Key {
	return &Key{
		key: key,
	}
}

func NewEmptyKey() *Key {
	return &Key{}
}

func (k *Key) UpdateKey(key int) {
	k.key = key
}

func (k *Key) UpdateNextNode(node *Node) {
	k.nextNode = node
}

type Node struct {
	pointer  int
	capacity int
	parent   *Node
	linkNode *Node
	key      []*Key
	isLeaf   bool //false
	// counter  int  // remove all allocation
}

// bug is here
func (n *Node) Link(node *Node) {
	if n.linkNode != nil {
		node.linkNode = n.linkNode
	}

	n.linkNode = node
}

func NewNode(degree int) *Node {
	return &Node{
		key: make([]*Key, degree+1),
	}
}

func (n *Node) SetParent(parent *Node) {
	n.parent = parent
}

// change node to be leaf
func (n *Node) SetLeaf() {
	n.isLeaf = true
}

type BPTree struct {
	degree   int // max value in node
	capacity int //for better performance, capacity allow us preallocation of node
	root     *Node
	memory   []*Node
	position int
}

// Constructor for new B+ Tree
func NewBPTree(capacity, degree int) *BPTree {
	return &BPTree{
		capacity: capacity,                   //capacity for memory allocation
		degree:   degree,                     //max number of key in node
		memory:   make([]*Node, 0, capacity), //allocate memory for new nodes
	}
}

// check if is root
func (t BPTree) IsRoot() bool {
	return t.root == nil
}

// create root
func (t *BPTree) CreateRoot(key int) {
	t.root = t.CreateNode() //need to create something to add in memory
	t.root.InsertKey(key)   // add key
	t.root.SetLeaf()        //Set to bee leaf
}

// Create node
func (t *BPTree) CreateNode() *Node {
	t.memory = append(t.memory, NewNode(t.degree)) //Create new node

	node := t.memory[len(t.memory)-1]
	node.key[0] = NewEmptyKey()

	return node //return to us new created node
}

// add just one key in the node
func (n *Node) AppendKey(position, key int) {
	n.key[position] = NewKey(key)
	n.pointer++
}

// append more keys to node
func (n *Node) AppendKeys(position int, key []*Key) {
	copy(n.key[position:], key)
	n.pointer += len(key)
}

// value will be added later
func (t *BPTree) Insert(key int) {
	if t.IsRoot() { //check if exist root
		t.CreateRoot(key) //create new node who is root
		return
	}

	leaf := t.Search(key) //find right node

	current := leaf.AppendToLeaf(key, t) //append to leaf

	for current != nil && current.pointer == t.degree {

		parent := current.parent //get parent

		middleKey := (t.degree / 2) //which key we will use in the parent
		//update parent
		// fmt.Println(n.key[4])

		// fmt.Println("yes", update.linkNode.key[0])

		newNode := t.CreateNode()                        //create new node
		newNode.AppendKeys(0, current.key[middleKey+1:]) //move to next node half // we will remove free pointer (0)
		newNode.pointer--

		// update := current.key[middleKey+1].nextNode

		// //update part of leaf
		// for i := 0; i < newNode.pointer+1; i++ {
		// 	update.SetParent(newNode)
		// 	update = update.linkNode //go to next leaf node
		// }

		// update := current.key[middleKey+1].nextNode

		//update part of leaf
		for i := 0; i < newNode.pointer+1; i++ {
			current.key[middleKey+1+i].nextNode.SetParent(newNode)
			// update = update.linkNode //go to next leaf node
		}

		if parent == nil { // if not exist then create them
			parent = t.CreateNode() //create parent
			current.parent = parent //kako je moguce da ja ovdje ovo moram updatovati ako je ovo node
			t.root = parent         // set new root
		}

		newNode.SetParent(parent) //set the parent of this node

		i := parent.InsertKey(current.key[middleKey].key)
		parent.key[i].UpdateNextNode(current)

		parent.key[i+1].UpdateNextNode(newNode)
		current.pointer -= (len(current.key[:middleKey]) + 1)

		current = current.parent
	}
}

// add key to leaf
func (n *Node) AppendToLeaf(key int, t *BPTree) *Node {
	n.InsertKey(key) //insert to leaf

	if n.pointer == t.degree { //Check if is full
		parent := n.parent //get parent

		if parent == nil { // if not exist then create them
			parent = t.CreateNode() //create parent
			n.parent = parent
			t.root = parent // set new root
		}

		middleKey := (t.degree / 2) //which key we will use in the parent

		newNode := t.CreateNode()                        //create new node
		newNode.AppendKeys(0, n.key[middleKey:t.degree]) //move to next node half // we will remove free pointer (0)

		// newNode.pointer--
		newNode.SetLeaf()         //set that node to be leaf
		newNode.SetParent(parent) //set the parent of this node

		i := parent.InsertKey(n.key[middleKey].key)

		parent.key[i].UpdateNextNode(n)

		parent.key[i+1].UpdateNextNode(newNode)
		n.pointer -= (len(n.key[:middleKey]) + 1)

		n.Link(newNode) // link current node with newNode
	}

	return n.parent
}

// need to inset key on right position in node
func (n *Node) InsertKey(key int) int {
	i := Find(n.key, key, n.pointer) //find position
	copy(n.key[i+1:], n.key[i:])     //make space for this position
	n.AppendKey(i, key)              //append key

	return i
}

func (t *BPTree) Search(key int) *Node {
	current := t.root //start searching from root, from start

	for current.isLeaf != true { //go inside the tree unit come to leaf

		current = current.key[Find(current.key, key, current.pointer)].nextNode //next deeper node
	}

	return current //return leaf
}

// find right position in node, for inserting place or for next deeper node
func Find(list []*Key, key int, pointer int) int {
	for i := 0; i < pointer; i++ {
		if list[i].key > key {
			return i
		}
	}

	return pointer
}

func (t *BPTree) All() {
	current := t.root

	//go to left first key
	for current.isLeaf != true {
		current = current.key[0].nextNode
	}

	var counter int

	var less int = -1

	for current != nil {
		for i := 0; i < current.pointer; i++ {
			counter++

			if less <= current.key[i].key {
				less = current.key[i].key
			} else {
				break
			}

			fmt.Println(current.key[i])
		}

		fmt.Println()

		current = current.linkNode
	}

	fmt.Println(counter)
}

func main() {
	tree := NewBPTree(5000, 5)

	var counter int = 150

	// test := make([]int, counter)

	for i := 0; i < counter; i++ {
		rand := rand.Intn(5000)

		tree.Insert(rand)
		// test[i] = rand
	}

	tree.All()
}
