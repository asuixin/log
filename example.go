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

func InitLog(path string) error {

	logger, err := log.LoggerFromConfigAsFile(path)
	if err != nil {
		return err
	}
	err = log.ReplaceLogger(logger)
	if err != nil {
		return err
	}
	gpath = path
	return nil

}
func Reload() error {

	logger, err := log.LoggerFromConfigAsFile(gpath)
	if err != nil {
		return err
	}

	err = log.ReplaceLogger(logger)
	if err != nil {
		return err
	}
	defer log.Flush()
	return nil
}

func main() {

	var wg sync.WaitGroup

	err := InitLog("./config.xml")
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
			if err := Reload(); err != nil {
				fmt.Println("reload failed:", err)
				return
			}
			time.Sleep(3 * time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
}
