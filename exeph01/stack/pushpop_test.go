package stack

import (
	"testing"
)

func TestPushPop(t *testing.T) {
	c := new(Stack)
	c.Push(5)
	if c.Pop() != 5 {
		t.Log("pop err ")
		t.Fail()
	}
}
