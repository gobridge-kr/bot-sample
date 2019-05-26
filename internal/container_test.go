package bot

import (
	"testing"
)

type helper struct {
	*testing.T
}

func TestContainer(t *testing.T) {
	v := 123
	c := NewContainer(v)
	h := helper{t}
	
	h.assertSame(v, c.Get())
	h.assertSame(false, c.Frozen())

	v = 456
	c.Set(v)
	h.assertSame(v, c.Get())

	c.Freeze()
	h.assertSame(true, c.Frozen())

	v = 789
	h.assertSame(false, c.Set(v))
	h.assertNotSame(v, c.Get())
}

func (h *helper) assertSame(expected interface{}, actual interface{}) {
	if (expected != actual) {
		h.Errorf("Assertion failed\n\tExpected: %s\n\tActual: %s", expected, actual)
	}
}

func (h *helper) assertNotSame(expected interface{}, actual interface{}) {
	if (expected == actual) {
		h.Errorf("Assertion failed\n\tExpected: %s\n\tActual: %s", expected, actual)
	}
}
