package time

import (
	"testing"
	"time"
)

func TestTrackTime(t *testing.T) {
	defer TrackTime(time.Now()) // get the time easily/fast

	time.Sleep(500 * time.Millisecond)
}
