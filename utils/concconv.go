package utils

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
	"net"
)

func IsIPv6(ipdr net.IP) bool {
	return ipdr != nil && ipdr.To4() == nil
}

func IP2Intstr(ipaddr string) (string, error) {
	// this two functions is used to parse X-Real-IP for recording
	// the X-Real-IP can only be done on LB side
	var ipdr net.IP = net.ParseIP(ipaddr)
	if ipdr == nil {
		return "0", errors.New("Invalid IP Address")
	}
	ipint := big.NewInt(0)
	// detect if ipv6
	if IsIPv6(ipdr) {
		ipint.SetBytes(ipdr.To16())
	} else {
		ipint.SetBytes(ipdr.To4())
	}
	ipstr := ipint.String()
	return ipstr, nil
}

func Pack2BinData(data []byte) primitive.Binary {
	var bd = primitive.Binary{
		Subtype: 0x00,
		Data:    data,
	}
	return bd
}
