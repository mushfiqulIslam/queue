package queue

import "testing"

// Test Cases
func TestQueue_EnqueueInt(t *testing.T) {
	var q Queue
	q.Enqueue(5)
	q.Enqueue(8)
	q.Enqueue(9)
	if q.Peek() != 9 {
		t.Error("q.Peek() returns Wrong value, Expected 9")
	}
}

func TestQueue_EnqueueString(t *testing.T) {
	var q Queue
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	if q.Peek() != "c" {
		t.Error("q.Peek() returns Wrong value, Expected c")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	var q Queue
	q.Enqueue(5)
	q.Enqueue(8)
	q.Enqueue(9)
	if q.Dequeue() != 9 {
		t.Error("q.Dequeue() returns Wrong value, Expected 9")
	}

	if q.Dequeue() != 8 {
		t.Error("q.Dequeue() returns Wrong value, Expected 8")
	}

	if q.Peek() != 5 {
		t.Error("q.Peek() returns Wrong value, Expected 5")
	}
}

func TestQueue_Peek(t *testing.T) {
	var q Queue
	if q.Peek() != nil {
		t.Error("q.Peek() returns Wrong value, Expected nil")
	}
	q.Enqueue(8)
	if q.Peek() != 8 {
		t.Error("q.Peek() returns Wrong value, Expected 8")
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	var q Queue
	if q.Len() != 0 {
		t.Error("q.Len() returns Wrong value, Expected 0")
	}
	if q.Peek() != nil {
		t.Error("q.Peek() returns Wrong value, Expected nil")
	}
	q.Enqueue(8)
	if q.Len() != 1 {
		t.Error("q.Len() returns Wrong value, Expected 1")
	}
}

// Benchmark Tests
func BenchmarkQueue_Enqueue(b *testing.B) {
	var q Queue
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkQueue_Dequeue(b *testing.B) {
	var q Queue
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = q.Dequeue()
	}
}
