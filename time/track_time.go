package time

import (
	"fmt"
	"time"
)

func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	fmt.Println("elapsed:", elapsed)

	return elapsed
}
