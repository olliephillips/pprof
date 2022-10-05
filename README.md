## Pprof

Simple helper package to set up Pprof debug server. Minimal code footprint on application being inspected.

- Includes `net/http/pprof`
- Creates the http debug server on `http://localhost:9407`

### Usage

1. Include the package.
2. Call `pprof.Start()` at top of `main`
3. Optionally call `defer pprof.Hold()` for programs that will otherwise exit


### Example

```golang
package main

import (
	"github.com/olliephillips/pprof"
	"log"
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
