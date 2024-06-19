package Queue

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

type customer struct {
	wid       string
	rid       string
	ctx       context.Context
	priority  int32
	timestamp time.Time
	leno      int
	flag      int
}

func New(wid string, rid string, ctx context.Context, priority int32, timestamp time.Time, leno int, flag int) customer {
	return customer{
		wid:       wid,
		rid:       rid,
		ctx:       ctx,
		priority:  priority,
		timestamp: timestamp,
		leno:      leno,
		flag:      flag,
	}
}

type ByPriority struct {
	customers []customer
	offset    int
}

func (a ByPriority) Len() int { return len(a.customers) - a.offset }
func (a ByPriority) Less(i, j int) bool {
	return a.customers[a.offset+i].priority < a.customers[a.offset+j].priority
}
func (a ByPriority) Swap(i, j int) {
	a.customers[a.offset+i], a.customers[a.offset+j] = a.customers[a.offset+j], a.customers[a.offset+i]
}

type Queue struct {
	Size          int
	customers     []customer
	overdueOffset int
	mutex         sync.Mutex
}

func (q *Queue) Enqueue(c customer) (bool, error) {
	q.mutex.Lock() // Lock for concurrency safety

	defer q.mutex.Unlock()
	if q.Size > 0 && len(q.customers) >= q.Size {
		q.mutex.Unlock()
		return true, errors.New("Queue is full") // Queue is full
	}

	// Check special cases for priority and flag
	if c.priority == 4 || c.priority == 5 || c.priority == 6 {
		// Insert at the end
		q.customers = append(q.customers, c)
	} else if c.priority == 1 || c.priority == 2 || c.priority == 3 {
		// Insert based on flag condition, searching from behind
		insertIndex := len(q.customers)
		for insertIndex > 0 && q.customers[insertIndex-1].flag != 1 {
			insertIndex--
		}
		q.customers = append(q.customers[:insertIndex], append([]customer{c}, q.customers[insertIndex:]...)...)
	}

	fmt.Printf("wid:%s pid: %d pos:%d time:%s\n", c.wid, c.priority, c.leno, c.timestamp)

	for i := 0; i < len(q.customers); i++ {
		fmt.Print(q.customers[i].priority, " ")

	}
	fmt.Println()
	return false, nil
}

func (q *Queue) Dequeue() (string, string, context.Context, int32, time.Time, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	// Lock for concurrency safety

	if len(q.customers) == 0 {
		return " ", " ", nil, 0, time.Time{}, errors.New("Queue is empty")
	}
	c := q.customers[0]
	q.customers = q.customers[1:]
	if q.overdueOffset > 0 && (c.priority == 4 || c.priority == 5 || c.priority == 6) {
		q.overdueOffset--
	}

	return c.wid, c.rid, c.ctx, c.priority, c.timestamp, nil
}

func (q *Queue) GetLength() int {
	q.mutex.Lock() // Lock for concurrency safety
	defer q.mutex.Unlock()

	return len(q.customers)
}

func (q *Queue) Display() {
	q.mutex.Lock() // Lock for concurrency safety
	defer q.mutex.Unlock()
	if len(q.customers) == 0 {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println("Customers in the Queue:")
	for _, c := range q.customers {
		fmt.Printf("WID: %s, Priority: %d, Timestamp: %v\n", c.wid, c.priority, c.timestamp)
	}
}

func (q *Queue) IsEmpty() bool {
	q.mutex.Lock() // Lock for concurrency safety
	defer q.mutex.Unlock()
	return len(q.customers) == 0
}

func (q *Queue) SortCustomers() {
	q.mutex.Lock() // Lock for concurrency safety
	defer q.mutex.Unlock()

	sort.Sort(ByPriority{q.customers, q.overdueOffset})
}

func (q *Queue) MoveToFrontIfOverdue(id int32) {
	for {
		time.Sleep(time.Millisecond * 1)

		q.mutex.Lock()

		for i := 0; i < len(q.customers); i++ {
			if q.customers[i].flag == 0 && time.Since(q.customers[i].timestamp) > 19*time.Second {
				// Extract the overdue customer
				overdueCustomer := q.customers[i]
				overdueCustomer.flag = 1

				// Find the insertion position
				insertIndex := -1
				for j := 0; j < i; j++ {
					if q.customers[j].priority == 1 && q.customers[j].timestamp.After(overdueCustomer.timestamp) {
						insertIndex = j
						break
					}
				}

				// Only move the customer if a valid insertion index was found
				if insertIndex != -1 {
					// Remove the overdue customer from the current position
					q.customers = append(q.customers[:i], q.customers[i+1:]...)

					// Insert the overdue customer at the found position
					q.customers = append(q.customers[:insertIndex], append([]customer{overdueCustomer}, q.customers[insertIndex:]...)...)
				}

			}
		}

		q.mutex.Unlock()
	}
}
