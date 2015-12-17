package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	var a string = "Foo"
	var b string = foo()

	assert.Equal(t, a, b, "The two words should be the same.")
}