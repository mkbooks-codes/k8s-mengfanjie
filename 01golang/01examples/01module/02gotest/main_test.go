package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func add(a int, b int) int {
	return a + b
}
func TestIncrease(t *testing.T) {
	t.Log("Start testing")
	assert.Equal(t, add(1, 2), 3, "Add Error!")
}
