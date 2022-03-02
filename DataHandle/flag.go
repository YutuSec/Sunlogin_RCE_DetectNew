package DataHandle

import (
	"flag"
	"fmt"
	"os"
)

var IP string
var Cmd string
var Thread int

func init() {

	flag.StringVar(&IP, "i", "", "探测的IP")
	flag.StringVar(&Cmd, "c", "", "执行命令")
	flag.IntVar(&Thread, "t", 100, "并发数量")

	// 修改提示信息
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: main.exe -i 192.168.0.1 -t 500 -c whoami\n不加-c参数仅探测漏洞")
		flag.PrintDefaults()
	}

	flag.Parse()

}
