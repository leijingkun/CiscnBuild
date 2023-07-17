package utils

import (
	"fmt"
	"net"
	"time"
)

func HoneyPot(ip string, port int) []string {
	var honeypots = []string{"glastopf", "Kippo", "HFish"}

	var results []string

	for _, honeypot := range honeypots {
		address := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.DialTimeout("tcp", address, time.Duration(3)*time.Second)
		if err == nil {
			results = append(results, honeypot)
			conn.Close()
		}
	}

	return results
}
