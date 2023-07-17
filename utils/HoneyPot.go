package utils

import "CiscnMap/config"

func HoneyPot(ip string) string {
	return randomFrom(config.Honeypot_list)
}
