package pprof

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

// we use this to cancel the snapshots
var done chan struct{}

// Start creates a http server using the default mux on port 9407
// with a link to the Pprof debugging pages
func Start(snapshotInterval ...time.Duration){
	go func() {
		fmt.Println("Pprof: serving debug information at `http://localhost:9407/debug/pprof`")
		if  err := http.ListenAndServe("localhost:9407", nil); err != nil {
			log.Fatalf("Pprof server failed to start: %v", err)
		}
	}()

	// have we set a snapshot time?
	if len(snapshotInterval) == 1 {
		done = make(chan struct{})
		go func(){
			for {
				select {
				case t:= <-time.After(snapshotInterval[0]):
					go func() {
						if err := snapshot(t); err != nil {
							fmt.Println("Pprof: snapshot error...", err)
						}
					}()
				case <-done:
					return
				}
			}
		}()
	}
}

// Hold prevents the program exiting so the Pprof data can be inspected
// Not required on programs that provide their own execution blocking like web servers
func Hold(cancelSnapshot ...bool){
	fmt.Println("Pprof: holding...")
	if len(cancelSnapshot) == 1 && cancelSnapshot[0] == true {
		// avoid send on nil channel if, if we are passing true, but snapshots
		// have not been requested.
		if done != nil {
			fmt.Println("Pprof: cancelling snapshot...")
			done <- struct{}{}
		}

	} else {
		fmt.Println("Pprof: continuing to snapshot...")
	}
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("\nPprof: exiting in 3 seconds...")
	time.Sleep(3 * time.Second)
}

// utility to create a snapshot in PNG format
func snapshot(t time.Time) error {
	file := fmt.Sprintf("./snapshot-heap-%v.png", t)
	cmd := exec.Command("go", "tool", "pprof", "-png", "http://localhost:9407/debug/pprof/heap")
	outfile, err := os.Create(file)
	if err != nil {
		return err
	}
	cmd.Stdout = outfile
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}