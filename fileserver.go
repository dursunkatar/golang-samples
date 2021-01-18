package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("")
	fmt.Println("F İ L E   S E R V E R | dursunkatar.com")
	fmt.Println("")
	getLocalIP()
	fmt.Println("Port:  8081 dinleniyor...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getLocalIP() {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}
}
