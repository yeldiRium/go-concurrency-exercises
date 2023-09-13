//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import (
	"strconv"
	"sync"

	"github.com/stretchr/testify/assert"
)

const (
	cycles        = 15
	callsPerCycle = 100
)

// RunMockServer simulates a running server, which accesses the
// key-value database through our cache
func RunMockServer(cache *KeyStoreCache, as *assert.Assertions) {
	var wg sync.WaitGroup

	for c := 0; c < cycles; c++ {
		wg.Add(1)
		go func() {
			for i := 0; i < callsPerCycle; i++ {

				wg.Add(1)
				go func(i int) {
					value := cache.Get("Test" + strconv.Itoa(i))
					if as != nil {
						as.Equal("Test" + strconv.Itoa(i), value)
					}
					wg.Done()
				}(i)

			}
			wg.Done()
		}()
	}

	wg.Wait()
}
