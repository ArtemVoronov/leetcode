package dfs

func dfs(node *Node, target int) bool {
	if node == nil {
		return false
	}

	if node.v == target {
		return true
	}

	if node.l != nil && dfs(node.l, target) {
		return true
	}

	if node.r != nil && dfs(node.r, target) {
		return true
	}

	return false
}

type Node struct {
	v int
	l *Node
	r *Node
}
