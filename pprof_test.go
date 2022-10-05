package pprof_test

import (
	"github.com/olliephillips/pprof"
	"log"
	"testing"
	"time"
)

// TestStart isn't a true test, it performs some work which should cause allocations
// and is used to check Pprof is serving on Start and that Hold prevents termination
func TestStart(t *testing.T) {
	pprof.Start()
	defer pprof.Hold() // optional

	// some work
	mp := map[int]int{}
	for i:=1; i < 10; i++{
		mp[i] = i
		log.Println("added to map")
		time.Sleep(500 * time.Millisecond)
	}
}


