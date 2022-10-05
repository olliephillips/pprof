package pprof

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"
)

// Start creates a http server using the default mux on port 9407
// with a link to the Pprof debugging pages
func Start(){
	go func() {
		fmt.Println("Pprof: serving at debug information at `http://localhost:9407/debug/pprof`")
		if  err := http.ListenAndServe("localhost:9407", nil); err != nil {
			log.Fatalf("Pprof server failed to start: %v", err)
		}
	}()
}

// Hold prevents the program exiting so the Pprof data can be inspected
// Not required on programs that provide their own execution blocking like web servers
func Hold(){
	fmt.Println("Pprof: holding...")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("\nPprof: exiting in 3 seconds")
	time.Sleep(3 * time.Second)
}