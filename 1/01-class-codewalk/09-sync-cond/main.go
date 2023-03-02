package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]string)

func main() {
	var wg sync.WaitGroup
	m := sync.Mutex{}
	c := sync.NewCond(&m)

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		c.L.Unlock()
	}()

	// writes changes to sharedRsc
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		sharedRsc["rsc1"] = "foo"
		sharedRsc["rsc2"] = "bar"
		c.Broadcast()
		c.L.Unlock()
	}()

	wg.Wait()
}