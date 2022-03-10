package main

import (
	"awesomeProject3/SunClientRCE/DataHandle"
	"fmt"
	"sync"
	time2 "time"
)

var wg sync.WaitGroup

func main() {
	time := time2.Now()
	AliveIP := DataHandle.AliveCheck()
	ch2 := DataHandle.Port(AliveIP)
	for _, url := range ch2 {
		wg.Add(1)
		go DataHandle.HTTPProtScan(url, &wg)
	}
	wg.Wait()
	fmt.Printf("总耗时%v", time2.Since(time))

}
