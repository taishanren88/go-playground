package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	seats := []int{1, 0, 0, 0, 1, 0, 1}
	assert.Equal(t, 2, maxDistToClosest(seats))

	seats = []int{1, 0, 0, 0}
	assert.Equal(t, 3, maxDistToClosest(seats))
}
