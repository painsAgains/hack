package main

// 需要将Sprintf中的地址替换为目标地址
import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start scan port")
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
