package Queue

import (
	"errors"
	"fmt"
	"sort"
)

type customer struct {
	wid      string
	priority int
}

func New(wid string, priority int) customer {
	return customer{
		wid:      wid,
		priority: priority,
	}
}

type ByPriority []customer

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Less(i, j int) bool { return a[i].priority < a[j].priority }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (q *Queue) SortStudents() {
	sort.Sort(ByPriority(q.customers))
}

type Queue2 struct {
	items []string
}

func (q *Queue2) Enqueue2(item string) {
	q.items = append(q.items, item)
}
func (q *Queue2) SearchAndRemove(item string) bool {
	for i, v := range q.items {
		if v == item {
			// Remove the item from the slice
			q.items = append(q.items[:i], q.items[i+1:]...)
			return true
		}
	}
	return false
}

type Queue struct {
	Size      int
	customers []customer
}

func (q *Queue) Enqueue(c customer) (bool, error) {
	if q.Size > 0 && len(q.customers) >= q.Size {
		return true, errors.New("Queue is full") // Queue is full
	}
	q.customers = append(q.customers, c)
	return false, nil
}

func (q *Queue) Dequeue() (string, error) {
	if len(q.customers) == 0 {
		return "No workflow ID ", errors.New("Queue is empty")
	}
	c := q.customers[0]
	q.customers = q.customers[1:]
	return c.wid, nil
}
func (q *Queue) GetLength() int {
	return len(q.customers)
}
func (q *Queue) Display() {
	if len(q.customers) == 0 {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println("Customers in the Queue:")
	for _, c := range q.customers {
		fmt.Printf("WID: %s, Priority: %d\n", c.wid, c.priority)
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.customers) == 0
}

/*type ElementWithTime struct {
	Element int
	index   int
}

type Queue struct {
	customers []ElementWithTime
	Size     int
}

func (q *Queue) Sort() {
	sort.SliceStable(q.customers, func(i, j int) bool {
		return q.customers[i].Element < q.customers[j].Element
	})
}

func (q *Queue) SortQueue() {
	batchSize := 5
	length := q.GetLength()

	// Iterate over the elements in batches of size batchSize
	for i := 0; i < length; i += batchSize {
		end := i + batchSize
		if end > length {
			end = length
		}
		q.sortBatch(i, end)
	}
}

func (q *Queue) sortBatch(start, end int) {
	// Sorts the elements in the queue within the range [start, end)
	for i := start + 1; i < end; i++ {
		key := q.customers[i]
		j := i - 1
		for j >= start && q.customers[j].Element > key.Element {
			q.customers[j+1] = q.customers[j]
			j = j - 1
		}
		q.customers[j+1] = key
	}
}

func (q *Queue) Enqueue(elem int) error {
	if q.GetLength() == q.Size {
		fmt.Println("Overflow")
		return errors.New("queue is full wait ")
	}
	elementWithTime := ElementWithTime{Element: elem}
	q.customers = append(q.customers, elementWithTime)

	return nil
}

func (q *Queue) Dequeue() error {
	if q.IsEmpty() {
		//fmt.Println("UnderFlow")
		return errors.New("empty queue")
	}
	//element := q.customers[0].Element
	//fmt.Println(element)
	if q.GetLength() == 1 {
		q.customers = nil
		return nil
	}
	q.customers = q.customers[1:]
	return nil // Slice off the element once it is dequeued.
}

func (q *Queue) GetLength() int {
	return len(q.customers)
}

func (q *Queue) IsEmpty() bool {
	return len(q.customers) == 0
}

func (q *Queue) Search(elem int) (bool, error) {
	if q.IsEmpty() {
		return false, errors.New("empty queue")
	}
	for i, value := range q.customers {
		if value.Element == elem {
			//fmt.Println("Yes it is present")
			// Dequeue the element if found
			q.customers = append(q.customers[:i], q.customers[i+1:]...)
			return true, nil
		}
		//	fmt.Println(value)
	}
	//fmt.Println("No it is not present")
	return false, nil
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return 0
	}
	return q.customers[0].Element
}

func (q *Queue) Display() {

	if q.Size == 0 {
		fmt.Println("Empty queue")
	}
	for _, elem := range q.customers {
		fmt.Printf("%d\n", elem.Element)
	}
	fmt.Println()
}
*/
