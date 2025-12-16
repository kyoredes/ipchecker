package checker

import (
	"net"
)

func validateIp(ip string) bool {
	res := net.ParseIP(ip)
	return res != nil
}
