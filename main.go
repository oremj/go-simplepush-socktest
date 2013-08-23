package main

import (
	"log"
	"fmt"
	"runtime"
)


func main() {
	runtime.GOMAXPROCS(8)
	cc := make(chan error, 20000)
	for i := 0; i < 400000; i++ {
		url := fmt.Sprintf("ws://10.249.9.119:%d", 80 + (i % 10))
		go func(url string) {
			_, err := NewClient(url)
			if err != nil {
				cc <- err
				return
			}
			//c.Register("atestchannel")
			cc <- nil
		}(url)
	}

	connected := 0
	for {
		log.Println("Waiting")
		err := <-cc
		if err != nil {
			log.Println(err)
		} else {
			connected++
			log.Printf("connected: %d", connected)
		}
	}
}
