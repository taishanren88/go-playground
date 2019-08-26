package main

import "fmt"

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	fmt.Println("HERE")
	 assert.Equal(t, uniquePaths(3, 2), int(3));
     assert.Equal(t, uniquePaths(7, 3), int(28));

}