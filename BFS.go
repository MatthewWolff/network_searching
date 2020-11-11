package main

import "fmt"

// FIFO!
var queue = make([][]*Node, 0) // track the nodes we're gonna visit

func (node *Node) bfs(searchFor int) (path []*Node) {
	enqueue(node)
	for len(queue) != 0 {
		printQueue()
		curr := dequeue()
		fmt.Printf("Current head is %d. ", curr.id)
		if visited[curr] {
			fmt.Printf("Already visited %d!\n", curr.id)
			continue // next iteration!
		}

		visit(curr) // don't come here again
		if curr.id == searchFor {
			fmt.Println("It's a match!")
			return getPathFromQueue()
		} else {
			fmt.Printf("Children to enqueue are %v\n", curr.children)
			for _, child := range curr.children {
				enqueue(child)
			}
		}
	}
	return nil
}

var currPath = make([]*Node, 0)

// enqueue item onto the queue!
func enqueue(node *Node) {
	path := make([]*Node, len(currPath))
	copy(path, currPath)
	path = append(path, node)
	queue = append(queue, path)
}

// dequeue item from the queue!
func dequeue() *Node {
	currPath, queue = queue[0], queue[1:]
	node := currPath[len(currPath)-1]
	return node
}

// recovering a path for a bfs is non-trivial, don't worry about this
func getPathFromQueue() []*Node {
	queue = make([][]*Node, 0)     // reset stack
	visited = make(map[*Node]bool) // reset visited

	path := make([]*Node, len(currPath))
	copy(path, currPath)
	currPath = make([]*Node, 0)
	return path
}

func printQueue() {
	currQueue := make([]*Node, 0)
	for _, path := range queue {
		currQueue = append(currQueue, path[len(path)-1])
	}
	fmt.Printf("Queue: %v\n", currQueue)
}
