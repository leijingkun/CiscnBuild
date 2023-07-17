package utils

import "CiscnMap/config"

func DeviceDetect(ip string) string {
	return randomFrom(config.Device_list)

}
