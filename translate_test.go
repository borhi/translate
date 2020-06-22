package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, "h2ll4", Encode("hello"))
}

func TestDecode(t *testing.T) {
	assert.Equal(t, "hi there", Decode("h3 th2r2"))
}