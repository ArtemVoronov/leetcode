package issame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	assert.True(t, isSameTree(p, q))
}

func TestCase2(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	assert.False(t, isSameTree(p, q))
}
func TestCase3(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}

	assert.False(t, isSameTree(p, q))
}
