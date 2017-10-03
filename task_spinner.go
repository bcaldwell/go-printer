package printer

import "fmt"
import "sync"

var spinStates = []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}

const (
	running = iota
	success
	failure
)

func NewTaskSpinner(name string) *TaskSpinner {
	ch := make(chan string)
	t := TaskSpinner{
		Name:  name,
		Ch:    ch,
		tick:  0,
		State: running,
	}
	t.init()
	return &t
}

type TaskSpinner struct {
	State    int
	FinalMSG string
	MSG      string
	Name     string
	Ch       chan string
	tick     int
	m        sync.Mutex
}

func (t *TaskSpinner) init() {
	go func(t *TaskSpinner) {
		for {
			msg, more := <-t.Ch
			t.m.Lock()

			//false when channel is closed
			if !more {
				t.State = success
				t.m.Unlock()
				return
			}
			t.MSG = msg
			t.m.Unlock()
		}
	}(t)
}

func (t *TaskSpinner) draw() {
	t.m.Lock()
	defer t.m.Unlock()

	prefix := ""

	t.tick++
	spin := spinStates[t.tick%len(spinStates)]

	spinColor := 35
	if (t.tick % len(spinStates)) < len(spinStates)/2 {
		spinColor = 36
	}
	switch t.State {
	case running:
		fmt.Printf("\r%s\x1b[%dm%c\x1b[0m %s[%s]\x1b[K\n", prefix, spinColor, spin, t.Name, t.MSG)
	case success:
		fmt.Printf("%s\x1b[K\x1b[32m✓\x1b[0m %s\x1b[1B\r", prefix, t.Name)
	case failure:
		fmt.Printf("%s\x1b[K\x1b[31m✗\x1b[0m %s\x1b[1B\r", prefix, t.Name)
	}
}
