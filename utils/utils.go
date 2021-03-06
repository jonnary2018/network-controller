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

// GenerateVethName : Use pod-uuid and container-vethName to generate vtxXXXXXXXXX for bridge's interface name
func GenerateVethName(podUUID, containerVethName string) string {
	str := sha256String(podUUID + containerVethName)
	return fmt.Sprintf("vtx%s", str[0:9])
}

// IsValidIP : Check IP address is valided
func IsValidIP(str string) bool {
	return net.ParseIP(str) != nil
}

// IsValidCIDR : Check CIDR is valided
func IsValidCIDR(str string) bool {
	_, _, err := net.ParseCIDR(str)
	return err == nil
}

// IsValidVLANTag : Check VLAN tag is valided
func IsValidVLANTag(tag int32) bool {
	if tag < 0 || tag > 4095 {
		return false
	}
	return true
}
