package helpers

import (
	"fmt"
	"time"
)

func GetCurrentTime() string {
	currentTime := time.Now()
	stime := fmt.Sprintf(currentTime.Format("01-02-2006"))
	return stime
}
