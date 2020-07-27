package src

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var n, d, b = 5, 1, 1
var array = []int{1, 0, 0, 0, 4}
var n2, d2, b2 = 6, 1, 2
var array2 = []int{3, 8, 0, 1, 0, 0}

func Test1(t *testing.T) {
	excpect := 1
	assert.Equal(t, Solution(n, d, b, array), excpect)

}

func Test2(t *testing.T) {
	excpect := 2
	assert.Equal(t, Solution(n2, d2, b2, array2), excpect)
}
