package main

import "CiscnMap/utils"

type Service struct {
	port        int
	Protocol    string
	service_app string
}

type Host struct {
	services   []Service
	deviceinfo string
	honeypot   string
	timestamp  string
}

func main() {
	utils.PortScan()
}
