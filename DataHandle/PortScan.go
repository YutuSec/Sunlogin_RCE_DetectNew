package DataHandle

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func PortScan(addr string, ResultChan chan<- string, wg *sync.WaitGroup) {
	conn, err := net.DialTimeout("tcp", addr, time.Duration(3)*time.Second)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err == nil {
		fmt.Printf("监测到端口开放：%v\n", addr)
		wg.Add(1)
		ResultChan <- fmt.Sprintf("http://%v", addr)
	}

}

var IPAlive chan string

func AliveCheck() chan string { //多线程用chan接受返回值

	iplist, _ := CheckIP(IP)
	IPAlive := make(chan string, len(iplist))
	RunPing(iplist, IPAlive)
	return IPAlive
}
func Port(alive chan string) (Ch2 []string) {
	var wg sync.WaitGroup
	close(alive)
	var ScanChan = make(chan string, 25535*len(alive))
	var ResultChan = make(chan string, 25535*len(alive))

	go func() { /*输出结果的*/
		for found := range ResultChan {
			Ch2 = append(Ch2, found)
			wg.Done()
		}
	}()
	for i := 0; i < Thread; i++ {
		go func() {
			for addr := range ScanChan {
				PortScan(addr, ResultChan, &wg)
				wg.Done()
			}
		}()
	}
	for val1 := range alive {
		for i := 40000; i < 65535; i++ {
			wg.Add(1)
			ScanChan <- fmt.Sprintf("%v:%v", val1, i) //扫描目标

		}
	}
	//fmt.Println(ScanChan) //输出存货IP端口扫描队列
	wg.Wait()

	close(ScanChan)
	close(ResultChan)

	return Ch2
}
