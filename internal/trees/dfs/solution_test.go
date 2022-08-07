package dfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFound(t *testing.T) {
	root := &Node{
		v: 1,
		l: &Node{
			v: 2,
			l: &Node{
				v: 4,
			},
			r: &Node{
				v: 5,
			},
		},
		r: &Node{
			v: 3,
			l: &Node{
				v: 6,
			},
			r: &Node{
				v: 7,
			},
		},
	}
	assert.True(t, dfs(root, 6))
}

func TestNotFound(t *testing.T) {
	root := &Node{
		v: 1,
		l: &Node{
			v: 2,
			l: &Node{
				v: 4,
			},
			r: &Node{
				v: 5,
			},
		},
		r: &Node{
			v: 3,
			l: &Node{
				v: 6,
			},
			r: &Node{
				v: 7,
			},
		},
	}

	assert.False(t, dfs(root, 100))
}

func TestEmpty(t *testing.T) {
	root := &Node{}

	assert.False(t, dfs(root, 100))
}

func TestFoundOnlyRight(t *testing.T) {
	root := &Node{
		v: 1,
		l: &Node{
			v: 2,
			l: &Node{
				v: 4,
			},
			r: &Node{
				v: 5,
			},
		},
		r: &Node{
			v: 3,
			l: &Node{
				v: 6,
			},
			r: &Node{
				v: 7,
			},
		},
	}

	assert.True(t, dfs(root, 7))
}

func TestNotFoundOnlyRight(t *testing.T) {
	root := &Node{
		v: 1,
		l: &Node{
			v: 2,
			l: &Node{
				v: 4,
			},
			r: &Node{
				v: 5,
			},
		},
		r: &Node{
			v: 3,
			l: &Node{
				v: 6,
			},
			r: &Node{
				v: 7,
			},
		},
	}

	assert.False(t, dfs(root, 8))
}
