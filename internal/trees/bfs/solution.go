package bfs

func bfs(root *Node, target int) bool {
	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		if node.v == target {
			return true
		}
		queue = queue[1:]
		if node.l != nil {
			queue = append(queue, node.l)
		}
		if node.r != nil {
			queue = append(queue, node.r)
		}
	}

	return false
}

type Node struct {
	v int
	l *Node
	r *Node
}
