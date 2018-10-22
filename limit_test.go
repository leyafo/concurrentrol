package concurrentrol

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestConcurrentCount(t *testing.T) {
	var count int32
	Run(10, 100, func(i int) error {
		atomic.AddInt32(&count, int32(i))
		return nil
	})
	if count != 4950 {
		t.Errorf("concurrent error, expect=4950, actually=%d\n", count)
	}
}

func TestPassParameter(t *testing.T) {
	var strs []string
	for i := 0; i != 10; i++ {
		strs = append(strs, fmt.Sprintf("%d", i))
	}
	err := Run(20, len(strs), func(i int) error {
		if strs[i] != fmt.Sprintf("%d", i) {
			return fmt.Errorf("error occurred")
		}
		return nil
	})

	if err != nil {
		t.Errorf(err.Error())
	}
}

type customError struct {
	taskNumber int
}

func (e *customError) Error() string {
	return fmt.Sprintf("task %d is error!", e.taskNumber)
}

func TestExactTaskError(t *testing.T) {
	var err error
	err = Run(10, 100, func(i int) error {
		if i == 21 {
			return &customError{i}
		}
		return nil
	})
	if err != nil {
		t.Logf(err.Error())
	} else {
		t.Error("Test failed!")
	}
}
