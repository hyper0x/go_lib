package pool

import (
	"runtime/debug"
	"testing"
	"time"
)

var count = 0

func TestRedisCacheProviderSync(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error (Test, sync): %s\n", err)
		}
	}()
	sign := make(chan bool, 2)
	pool := &Pool{Id: "Test", Size: 100}
	err := pool.Init(initFunc)
	if err != nil {
		debug.PrintStack()
		t.Errorf("Init Error (Test, sync): %s\n", err)
		t.FailNow()
		return
	}
	go getFromLoop(pool, sign, 0)
	go putToLoop(pool, sign, 0)
	<-sign
	<-sign
}

func TestRedisCacheProviderAsync(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error (Test, async): %s\n", err)
		}
	}()
	sign := make(chan bool, 2)
	pool := &Pool{Id: "Test", Size: 100}
	err := pool.Init(initFunc)
	if err != nil {
		debug.PrintStack()
		t.Errorf("Init Error (Test, async): %s\n", err)
		t.FailNow()
		return
	}
	go getFromLoop(pool, sign, 100)
	go putToLoop(pool, sign, 100)
	<-sign
	<-sign
}

func BenchmarkRedisCacheProviderSync(b *testing.B) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			b.Errorf("Fatal Error (Benchmark, sync): %s\n", err)
		}
	}()
	sign := make(chan bool, 2)
	pool := &Pool{Id: "Test", Size: 100}
	err := pool.Init(initFunc)
	if err != nil {
		debug.PrintStack()
		b.Errorf("Init Error (Benchmark, sync): %s\n", err)
		b.FailNow()
		return
	}
	go getFromLoop(pool, sign, 0)
	go putToLoop(pool, sign, 0)
	<-sign
	<-sign
}

func BenchmarkRedisCacheProviderAsync(b *testing.B) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			b.Errorf("Fatal Error (Benchmark, async): %s\n", err)
		}
	}()
	sign := make(chan bool, 2)
	pool := &Pool{Id: "Test", Size: 100}
	err := pool.Init(initFunc)
	if err != nil {
		debug.PrintStack()
		b.Errorf("Init Error (Benchmark, async): %s\n", err)
		b.FailNow()
		return
	}
	go getFromLoop(pool, sign, 100)
	go putToLoop(pool, sign, 100)
	<-sign
	<-sign
}

func getFromLoop(pool *Pool, sign chan bool, timeoutMs int) {
	result := false
	for {
		element, ok := pool.Get(time.Duration(timeoutMs))
		//		infoMsg := fmt.Sprintf("Get element '%v'. (%v)\n", element, ok)
		//		LogInfo(infoMsg)
		if element == nil {
			result = false
		} else {
			if !ok {
				result = false
			} else {
				result = true
			}
		}
		if !ok {
			break
		}
	}
	logger.Infoln("Getting finish.")
	sign <- result
}

func putToLoop(pool *Pool, sign chan bool, timeoutMs int) {
	result := false
	for i := 0; i < 50; i++ {
		element, _ := initFunc()
		result = pool.Put(element, time.Duration(timeoutMs))
		//		infoMsg := fmt.Sprintf("Put element '%v': %v", element, result)
		//		LogInfo(infoMsg)
	}
	pool.Close()
	logger.Infoln("Putting finish.")
	sign <- result
}

func initFunc() (interface{}, error) {
	count++
	return count, nil
}
