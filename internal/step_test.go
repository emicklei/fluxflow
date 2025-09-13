package internal

import (
	"testing"
)

func TestStep(t *testing.T) {
	t.Run("Next", func(t *testing.T) {
		s1 := &step{}
		s2 := &step{}
		s1.Next(s2)
		if s1.next != s2 {
			t.Error("s1.next should be s2")
		}
		if s2.prev != s1 {
			t.Error("s2.prev should be s1")
		}
		// test idempotency
		s1.Next(s2)
		if s1.next != s2 {
			t.Error("s1.next should be s2")
		}
		if s2.prev != s1 {
			t.Error("s2.prev should be s1")
		}
	})

	t.Run("Prev", func(t *testing.T) {
		s1 := &step{}
		s2 := &step{}
		s2.Prev(s1)
		if s2.prev != s1 {
			t.Error("s2.prev should be s1")
		}
		if s1.next != s2 {
			t.Error("s1.next should be s2")
		}
		// test idempotency
		s2.Prev(s1)
		if s2.prev != s1 {
			t.Error("s2.prev should be s1")
		}
		if s1.next != s2 {
			t.Error("s1.next should be s2")
		}
	})
}
