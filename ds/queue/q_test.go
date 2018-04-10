package queue_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/rakin-ishmam/sw-median/ds/queue"
)

func TestQ(t *testing.T) {
	q := queue.New()

	tt := []struct {
		name       string
		op         func()
		exSize     int
		exFront    interface{}
		exFrontErr error
	}{
		{
			"push 1",
			func() {
				q.Push(1)
			},
			1,
			1,
			nil,
		},
		{
			"push 2",
			func() {
				q.Push(2)
			},
			2,
			1,
			nil,
		},
		{
			"push 3",
			func() {
				q.Push(3)
			},
			3,
			1,
			nil,
		},
		{
			"pop 1",
			func() {
				q.Pop()
			},
			2,
			2,
			nil,
		},
		{
			"pop 2",
			func() {
				q.Pop()
			},
			1,
			3,
			nil,
		},
		{
			"pop empty",
			func() {
				q.Pop()
			},
			0,
			nil,
			errors.New("empty"),
		},
	}

	for _, st := range tt {
		t.Run(st.name, func(t *testing.T) {
			st.op()

			if st.exSize != q.Size() {
				t.Errorf("exptected queue size is %v but got %v", st.exSize, q.Size())
			}

			fv, err := q.Front()
			if fmt.Sprintf("%v", err) != fmt.Sprintf("%v", st.exFrontErr) {
				t.Errorf("exptected queue front error is %v but got %v", st.exFrontErr, err)
			}

			if fv != st.exFront {
				t.Errorf("exptected queue front value is %v but got %v", st.exFront, fv)
			}
		})
	}
}
