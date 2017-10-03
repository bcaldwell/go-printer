package printer

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTaskSpinner(t *testing.T) {
	s := NewTaskSpinner("hello")
	ticker := time.NewTicker(100 * time.Millisecond)
	s.Ch <- "I work"
	go func() {
		time.Sleep(1 * time.Second)
		s.Ch <- "suppppppp"
		time.Sleep(10 * time.Second)
		close(s.Ch)
	}()

	s.draw()
	for _ = range ticker.C {
		fmt.Print("\x1b[1F")
		s.draw()
	}

}
