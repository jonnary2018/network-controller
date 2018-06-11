package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
)

func sha256String(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	md := hash.Sum(nil)
	return fmt.Sprintf("%s", hex.EncodeToString(md))
}

// GenerateVethName : Use pod-uuid and container-vethName to generate vethXXXXXXXX for bridge's interface name
func GenerateVethName(podUUID, containerVethName string) string {
	str := sha256String(podUUID + containerVethName)
	return fmt.Sprintf("veth%s", str[0:8])
}

func VerifyIP(str string) bool {
	return net.ParseIP(str) != nil
}

func VerifyCIDR(str string) bool {
	_, _, err := net.ParseCIDR(str)
	return err == nil
}
