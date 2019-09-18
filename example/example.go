package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	mylog "github.com/asuixin/log"
)

var log *mylog.MLog

func main() {

	var wg sync.WaitGroup

	log, err := mylog.NewMLog("./config.xml")
	if err != nil {
		fmt.Println("create MLog falied:", err)
		return
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {

			for {
				log.Logger().Infof("Pid=%d app.Version=%s", os.Getpid(), "hello")
				time.Sleep(10 * time.Nanosecond)
			}
			wg.Done()
		}()
	}

	wg.Add(1)
	go func() {

		for {
			if err := log.Reload(); err != nil {
				fmt.Println("reload failed:", err)
				return
			}
			time.Sleep(100 * time.Nanosecond)
		}
		wg.Done()
	}()
	wg.Wait()
}
