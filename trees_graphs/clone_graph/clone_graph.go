package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	return dfsWithInitialisation(node, map[*Node]*Node{})
}

func dfsWithInitialisation(node *Node, visited map[*Node]*Node) *Node {
	n, ok := visited[node]
	if ok {
		return n
	}

	newNode := &Node{Val: node.Val}
	visited[node] = newNode

	for _, neighbor := range node.Neighbors {
		n, ok := visited[neighbor]
		if !ok {
			n = dfsWithInitialisation(neighbor, visited)
		}

		newNode.Neighbors = append(newNode.Neighbors, n)
	}

	return newNode
}
