package queue

const minCapacity = 16

// Queue Initializes Empty Queue of integer
type Queue struct {
	queue       []interface{}
	length      int
	minCapacity int
}

// Len returns length of Queue
func (q *Queue) Len() int {
	if q == nil {
		return 0
	}
	return q.length
}

// Enqueue adds an item to the Queue
func (q *Queue) Enqueue(value interface{}) {
	q.queue = append(q.queue, value)
	q.length++
	if q.minCapacity == 0 {
		q.minCapacity = minCapacity
	}
}

// Dequeue removes the first element from the Queue and returns it
func (q *Queue) Dequeue() interface{} {
	if q.length <= 0 {
		return nil
	}
	val := q.queue[q.length-1]
	q.queue = q.queue[:q.length-1]
	q.length--
	return val
}

// Peek returns first value from the Queue
func (q *Queue) Peek() interface{} {
	if q.length <= 0 {
		return nil
	}
	return q.queue[q.length-1]
}

// IsEmpty returns  true if Queue is empty, false otherwise
func (q *Queue) IsEmpty() bool {
	if q.length <= 0 {
		return true
	}
	return false
}

/*
	IsFull checks length of the Queue is greater or equal to the minCapacity. if
	equal or greater return true, false otherwise.
*/
func (q *Queue) IsFull() bool {
	if q.length >= q.minCapacity {
		return true
	}
	return false
}
