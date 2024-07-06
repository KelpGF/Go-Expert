package examples

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

var counter int64 = 0

// ab -n 1000 -c 100 http://localhost:3000/
// identify race condition: go run -race main.go
func RaceCondition() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter++ // problem: any goroutine can access and modify the counter at the same time

		w.Write([]byte(fmt.Sprintf("Counter: %d", counter)))
	})

	http.ListenAndServe(":3000", nil)
}

func RaceConditionMutexSolution() {
	mutex := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		counter++
		mutex.Unlock()

		w.Write([]byte(fmt.Sprintf("Counter: %d", counter)))
	})

	http.ListenAndServe(":3000", nil)
}

func RaceConditionAtomicSum() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&counter, 1)

		w.Write([]byte(fmt.Sprintf("Counter: %d", counter)))
	})

	http.ListenAndServe(":3000", nil)
}
