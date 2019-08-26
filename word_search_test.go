package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	matrix := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}
	assert.Equal(t, true, exist(matrix, "ABCCED"))
	assert.Equal(t, true, exist(matrix, "SEE"))
	assert.Equal(t, false, exist(matrix, "ABCB"))
	assert.Equal(t, true, exist([][]byte{{'A'}}, "A"))

}
