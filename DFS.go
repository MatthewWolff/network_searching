package main

import "fmt"

// LIFO!
var stack = make([]*Node, 0) // track our path

// Question: why can't we just use the stack to keep track of where we've visited?

func (node *Node) dfs(searchFor int) (path []*Node) {
	fmt.Printf("called DFS on %d. ", node.id)
	if visited[node] {
		fmt.Println("already visited!")
		return nil
	}

	push(node)  // add current node to our path
	visit(node) // don't come here again

	fmt.Printf("visiting %d, with kids %v\n", node.id, node.children)

	// check if we have found our node
	if node.id == searchFor {
		return getPathFromStack()
	} else {
		// see if our children have a path
		for _, child := range node.children {
			if p := child.dfs(searchFor); p != nil {
				return p // return the path if so
			}
		}
	}

	pop() // didn't find anything, remove this node from our path
	return nil
}

// push onto the stack!
func push(node *Node) {
	stack = append(stack, node)
}

// pop off the stack!
func pop() {
	stack = stack[:len(stack)-1]
}

// getPath is a helper to clean up our global variables and create our return variable
func getPathFromStack() []*Node {
	var path = make([]*Node, len(stack))
	copy(path, stack)              // store solution
	stack = make([]*Node, 0)       // reset stack
	visited = make(map[*Node]bool) // reset visited
	return path
}
