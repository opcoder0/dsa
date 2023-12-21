package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushPop(t *testing.T) {

	type nothing interface{}

	h := NewHeap[int, nothing](false) // a max heap
	v100 := NewValue[int, nothing](100, nil)
	h.Push(*v100)
	v30 := NewValue[int, nothing](30, nil)
	h.Push(*v30)
	v50 := NewValue[int, nothing](50, nil)
	h.Push(*v50)

	v, err := h.Pop()
	assert.Nil(t, err)
	assert.Equal(t, v.Priority, 100)

	v, err = h.Pop()
	assert.Nil(t, err)
	assert.Equal(t, v.Priority, 50)
}
