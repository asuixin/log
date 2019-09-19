package main

import (
	"fmt"
	"sync"
	"time"

	log "github.com/cihub/seelog"
)

var x int = 0
var m sync.Mutex
var gpath string

func LogInit() error {

	logger, err := log.LoggerFromConfigAsFile(gpath)
	if err != nil {
		return err
	}
	err = log.ReplaceLogger(logger)
	if err != nil {
		return err
	}
	return nil

}
func LogReload() error {
	return LogInit()
}

func main() {

	var wg sync.WaitGroup
	gpath = "./config.xml"
	err := LogInit()
	if err != nil {
		fmt.Println("InitLog err:", err)
		return
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				m.Lock()
				x++
				x1 := x
				m.Unlock()
				log.Infof("%s x=%dv", "hello", x1)
				time.Sleep(10 * time.Nanosecond)
			}
			wg.Done()
		}()
	}

	wg.Add(1)
	go func() {
		for {
			if err := LogReload(); err != nil {
				fmt.Println("reload failed:", err)
				return
			}
			time.Sleep(3 * time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
}
