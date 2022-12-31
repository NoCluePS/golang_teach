package queue

import "testing"

func TestAddQ(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue elemnt count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
	}
	if q.Append(4) {
		t.Errorf("Should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Should be able to retrieve the three items")
		}

		if item != i {
			t.Errorf("Should be equal to i, each time it comes out: %v, want %v", item, i)
		}
	}
	// Queue is empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("Queue should be empty at this point, got: %v\n", item)
	}
}