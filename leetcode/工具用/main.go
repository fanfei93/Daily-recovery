package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	//memProfile := flag.String("memProfile", "", "write mem profile to file")

	//flag.Parse()
	//f, err := os.Create(*memProfile)
	//if err != nil {
	//	log.Fatal(err)
	//}

	c := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		idleDuration := 4 * time.Second
		timer := time.NewTimer(idleDuration)
		defer timer.Stop()
		for {
			//timer.Reset(idleDuration)
			select {
			case <-timer.C:
				log.Println("超时")
			case <-c:
				return
			}
		}
	}()
	wg.Wait()



	//pprof.WriteHeapProfile(f)
}