package heap

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Value[K constraints.Ordered, V interface{}] struct {
	Priority K
	V        V
}

type Heap[K constraints.Ordered, V interface{}] struct {
	backend []Value[K, V]
	minHeap bool
}

func NewHeap[K constraints.Ordered, V interface{}](minHeap bool) *Heap[K, V] {
	return &Heap[K, V]{
		backend: make([]Value[K, V], 0),
		minHeap: minHeap,
	}
}

func NewValue[K constraints.Ordered, V interface{}](prio K, val V) *Value[K, V] {
	return &Value[K, V]{
		Priority: prio,
		V:        val,
	}
}

func (heap *Heap[K, V]) left(index int) (int, error) {
	lIndex := (index * 2) + 1
	if lIndex < 0 || lIndex > len(heap.backend)-1 {
		return -1, errors.New("index out of bounds")
	}
	return lIndex, nil
}

func (heap *Heap[K, V]) right(index int) (int, error) {
	rIndex := (index * 2) + 2
	if rIndex < 0 || rIndex > len(heap.backend)-1 {
		return -1, errors.New("index out of bounds")
	}
	return rIndex, nil
}

func (heap *Heap[K, V]) parent(index int) int {
	return (index - 1) / 2
}

func (heap *Heap[K, V]) valueAt(index int) (*Value[K, V], error) {
	if index < 0 || index > len(heap.backend)-1 {
		return nil, errors.New("index out of bounds")
	}
	return &heap.backend[index], nil
}

func (heap *Heap[K, V]) Push(v Value[K, V]) {
	heap.backend = append(heap.backend, v)
	index := len(heap.backend) - 1
	bubbleUp := false
	for index > 0 || !bubbleUp {
		parentIndex := heap.parent(index)
		pv := heap.backend[parentIndex]
		if heap.minHeap {
			if v.Priority < pv.Priority {
				bubbleUp = true
			}
		} else {
			if v.Priority > pv.Priority {
				bubbleUp = true
			}
		}
		heap.backend[parentIndex], heap.backend[index] = heap.backend[index], heap.backend[parentIndex]
		index = parentIndex
	}
}

func (heap *Heap[K, V]) lesserChildIndex(index int) (int, error) {
	value := heap.backend[index]
	lIndex, errL := heap.left(index)
	rIndex, errR := heap.right(index)
	if errL == nil && errR == nil {
		minIndex := rIndex
		if heap.backend[lIndex].Priority < heap.backend[rIndex].Priority {
			minIndex = lIndex
		}
		if value.Priority < heap.backend[minIndex].Priority {
			return minIndex, nil
		}
	} else if errL == nil && errR != nil {
		if value.Priority < heap.backend[lIndex].Priority {
			return lIndex, nil
		}
	} else if errL != nil && errR == nil {
		if value.Priority < heap.backend[rIndex].Priority {
			return rIndex, nil
		}
	}
	return -1, errors.New("no lesser child")
}

func (heap *Heap[K, V]) greaterChildIndex(index int) (int, error) {
	value := heap.backend[index]
	lIndex, errL := heap.left(index)
	rIndex, errR := heap.right(index)
	if errL == nil && errR == nil {
		maxIndex := rIndex
		if heap.backend[lIndex].Priority > heap.backend[rIndex].Priority {
			maxIndex = lIndex
		}
		if value.Priority > heap.backend[maxIndex].Priority {
			return maxIndex, nil
		}
	} else if errL == nil && errR != nil {
		if value.Priority > heap.backend[lIndex].Priority {
			return lIndex, nil
		}
	} else if errL != nil && errR == nil {
		if value.Priority > heap.backend[rIndex].Priority {
			return rIndex, nil
		}
	}
	return -1, errors.New("no greater child")
}

func (heap *Heap[K, V]) Pop() (*Value[K, V], error) {
	if len(heap.backend) == 0 {
		return nil, errors.New("underflow")
	}
	deleted := heap.backend[0]
	if len(heap.backend) == 1 {
		heap.backend = make([]Value[K, V], 0)
		return &deleted, nil
	}
	last := len(heap.backend) - 1
	heap.backend[0] = heap.backend[last]
	heap.backend = append([]Value[K, V](nil), heap.backend[:last]...)
	bubbleDown := true

	// Start at the top
	index := 0
	for bubbleDown {
		if heap.minHeap {
			lesserChildIndex, err := heap.lesserChildIndex(index)
			if err != nil {
				bubbleDown = false
			} else {
				heap.backend[lesserChildIndex], heap.backend[index] = heap.backend[index], heap.backend[lesserChildIndex]
				index = lesserChildIndex
			}
		} else {
			greaterChildIndex, err := heap.greaterChildIndex(index)
			if err != nil {
				bubbleDown = false
			} else {
				heap.backend[greaterChildIndex], heap.backend[index] = heap.backend[index], heap.backend[greaterChildIndex]
				index = greaterChildIndex
			}
		}
	}
	return &deleted, nil
}

func (heap *Heap[K, V]) Print() {
	for i, v := range heap.backend {
		fmt.Printf("[%d] = %v\n", i, v)
	}
	fmt.Println()
}
