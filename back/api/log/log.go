package log

import (
	"fmt"
	"time"
)

// Logging -
type Logger struct {
}

// PrintLog -
func (l *Logger) Log(level, content string) {
	t := time.Now()
	fmt.Println(fmt.Sprintf("%v : %s : %v", t.Format(time.RFC3339), level, content))
}
