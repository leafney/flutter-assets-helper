/**
 * @Author:      leafney
 * @Date:        2021-01-25 09:59
 * @Project:     flutter-assets-helper
 * @Description:
 */

package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"golang.org/x/sync/singleflight"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func waitUntilFind(filename string) error {
	for {
		time.Sleep(1 * time.Second)
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return err
			}
		}
		break
	}
	return nil
}

func run() {
	//if err := cmd.Execute(); err != nil {
	//	os.Exit(1)
	//}

	filename := "./tmp"
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	err = watcher.Add(filename)
	if err != nil {
		log.Fatalln(err)
	}

	addNewCh := make(chan bool, 1)

	go func() {

		for {
			select {
			case event := <-watcher.Events:
				switch {
				//case event.Op&fsnotify.Write == fsnotify.Write:
				//	log.Printf("Write:  %s: %s", event.Op, event.Name)
				case event.Op&fsnotify.Create == fsnotify.Create:
					log.Printf("Create: %s: %s", event.Op, event.Name)

					addNewCh <- true
					//case event.Op&fsnotify.Remove == fsnotify.Remove:
					//	log.Printf("Remove: %s: %s", event.Op, event.Name)
					//	removeCh <- true
					//case event.Op&fsnotify.Rename == fsnotify.Rename:
					//	log.Printf("Rename: %s: %s", event.Op, event.Name)
					//	renameCh <- true
					//case event.Op&fsnotify.Chmod == fsnotify.Chmod:
					//	log.Printf("Chmod:  %s: %s", event.Op, event.Name)
				}
			case err := <-watcher.Errors:
				log.Print(err)

			}
		}
	}()

	go func() {
		g := singleflight.Group{}
		for {
			select {
			case <-addNewCh:
				fmt.Println("开始处理了")

				go func() {
					res, err, shared := g.Do("ttt2", func() (interface{}, error) {
						// 设置等待一定时间后开始处理
						time.Sleep(5 * time.Second)
						fmt.Println("start test")

						allFiles := loadFiles()
						copyIosFiles(allFiles)

						return 1, nil
					})
					fmt.Println(res, err, shared)
				}()

			}
		}
	}()

	fmt.Println("Press Ctrl+C to stop")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
