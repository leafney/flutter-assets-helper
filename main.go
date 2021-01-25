package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"

	"sync"
	"time"
)

func main() {

	run()

	//allFiles := loadFiles()
	//copyIosFiles(allFiles)

	//test()
}

func test2() (interface{}, error) {
	fmt.Println("start test")
	time.Sleep(3 * time.Second)
	return 100, nil
}

func test() {
	g := singleflight.Group{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			res, err, shared := g.Do("test", test2)
			fmt.Println(res, err, shared)

		}()
	}
	wg.Wait()
}
