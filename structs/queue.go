package structs

// Queue simple abstraction of a queue
type Queue struct {

	queue []int
}

// Enqueue an int item to the queue
func (q *Queue) Enqueue(item int){

	q.queue = append(q.queue, item)
}

// Dequeue and return an item from the queue
func (q *Queue) Dequeue() int{
	
	item := q.queue[0]

	q.queue = q.queue[1:]
	return item
}

// Len return the length of the queue
func (q* Queue) Len() int{
	return len(q.queue)
}