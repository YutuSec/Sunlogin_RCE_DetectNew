package DataHandle

import (
	"errors"
	"net"
	"strings"
)

func CheckIP(ips string) (iplist []string, err error) {
	if strings.Contains(ips, "/") {
		ip, ipNet, err := net.ParseCIDR(ips) //返回IP 192.0.2.1 网络192.0.2.0/24
		if err != nil {
			return iplist, errors.New(ips + "格式错误")
		}
		for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) { //获取网段掩码，获取在本网段的IP，复制IP
			iplist = append(iplist, ip.String())
		}
	} else {
		trial := net.ParseIP(ips)
		if trial.To4() == nil {
			return iplist, errors.New(ips + "格式错误")
		}
		iplist = append(iplist, trial.String())
	}
	return iplist, nil
}
func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
