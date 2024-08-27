//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync/atomic"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  *int64 // in seconds
}

const FreeAvailableTime = 10

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.TimeUsed == nil {
		u.TimeUsed = new(int64)
		*u.TimeUsed = 0
	}

	tickChan := time.Tick(1 * time.Second)

	processFinished := make(chan bool)

	go func() {
		process()
		processFinished <- true
	}()

	for {
		select {
		case <-tickChan:
			newTimeUsed := atomic.AddInt64(u.TimeUsed, 1)
			if newTimeUsed >= FreeAvailableTime {
				return false
			}
		case <-processFinished:
			return true
		}
	}
}

func main() {
	RunMockServer()
}
