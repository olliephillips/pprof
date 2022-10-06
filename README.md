## Pprof

Simple helper package to set up Pprof debug server. Minimal code footprint on application being inspected.

- Includes `net/http/pprof`
- Creates the http debug server on `http://localhost:9407`
- Can optionally use the `go tool pprof` to take PNG snapshots at desired interval**

** Must be installed, with depenedencies for creating images.
### Usage

1. Include the package.
2. Call `pprof.Start()` at top of `main`
3. Optionally call `defer pprof.Hold()` for programs that would otherwise exit but you wish to inspect

### Snapshots
To capture heap diagrams in `.png` format at a specified interval, pass an optional `snapshotInterval` when starting i.e. 
`pprof.Start(2 * time.Second)`. Files will be written with timestamps to the current directory.

If using`pprof.Hold()`the snapshots will continue at the interval while holding. If not required pass `true` to the 
optional `cancelSnapshot` parameter. i.e. `pprof.Hold(true)`.

### Examples

```golang

package main

import (
	"github.com/olliephillips/pprof"
	"log"
	"time"
)

func main(){
	pprof.Start()
	defer pprof.Hold() // optional, blocks exit

	// some work
	mp := map[int]int{}
	for i:=1; i < 10; i++{
		mp[i] = i
		log.Println("added to map")
		time.Sleep(500 * time.Millisecond)
	}
}
```

With Snapshots at 5 second interval, cancelled once the hold starts

```golang

package main

import (
	"github.com/olliephillips/pprof"
	"log"
	"time"
)

func main(){
	pprof.Start(5 * time.Second) // snapshots at 5 second intervals
	defer pprof.Hold(true) // cancel the snapshots by passing 'true'

	// some work
	mp := map[int]int{}
	for i:=1; i < 10; i++{
		mp[i] = i
		log.Println("added to map")
		time.Sleep(500 * time.Millisecond)
	}
}
```


### License

MIT