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
	//iplist, _ := DataHandle.CheckIP(DataHandle.IP)
	//fmt.Println(iplist)
	//var ips = make(chan string, len(iplist)*25535)
	//DataHandle.RunPing(iplist, ips)
	//close(ips)
	//
	//for val1 := range ips {
	//	for i := 40000; i < 65535; i++ {
	//		fmt.Println(val1, i)
	//	}
	//}

	AliveIP := DataHandle.AliveCheck()
	ch2 := DataHandle.Port(AliveIP)
	for _, url := range ch2 {
		wg.Add(1)
		go DataHandle.HTTPProtScan(url, &wg)
	}
	wg.Wait()
	fmt.Printf("总耗时%v", time2.Since(time))

}
