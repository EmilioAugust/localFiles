package utils

import (
	"log"
	"main/app/internal/repository"
	"net"
	"strings"
	"math/rand/v2"
)

func CheckIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Found an error", err)
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Println("Found an error", err)
		}
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.IsPrivate() {
				if strings.HasPrefix(ipNet.IP.String(), "192") {
					return ipNet.IP.String()
				}
			}
		}
	}

	return "localhost"
}

func CheckForDevice(ua string) repository.FindDevice {
	var device repository.FindDevice
	if strings.Contains(ua, "Windows") {
		device.Name = "Windows"
		device.Type = "desktop"
	} else if strings.Contains(ua, "Mac") {
		device.Name = "Mac"
		device.Type = "desktop"
	} else if strings.Contains(ua, "Android") {
		device.Name = "Android"
		device.Type = "mobile"
	} else if strings.Contains(ua, "iPhone") {
		device.Name = "iOS"
		device.Type = "mobile"
	} else if strings.Contains(ua, "iPad") {
		device.Name = "iPad"
		device.Type = "mobile"
	} else {
		device.Name = "Windows"
		device.Type = "desktop"
	}
	return device
}

func GetRandomNumber(min, max int) int {
	return rand.IntN(max - min) + max
}