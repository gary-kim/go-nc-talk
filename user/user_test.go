package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContains(t *testing.T) {
	testSlice1 := []string{"index0", "index1", "index2", "index3", "index4"}

	assert.True(t, sliceContains(testSlice1, "index0"))
	assert.True(t, sliceContains(testSlice1, "index4"))
	assert.True(t, sliceContains(testSlice1, "index3"))
	assert.True(t, sliceContains(testSlice1, "index2"))
	assert.True(t, sliceContains(testSlice1, "index1"))

	assert.False(t, sliceContains(testSlice1, ""))
	assert.False(t, sliceContains(testSlice1, "index-1"))
	assert.False(t, sliceContains(testSlice1, "index234"))
	assert.False(t, sliceContains(testSlice1, "dkvwdfs"))
}
