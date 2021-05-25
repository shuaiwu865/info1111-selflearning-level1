package main


import (
	"fmt"
	"net"
)

func main() {
	for i:=20; i<120; i++ {
		address := fmt.Sprintf("192.168.124.12:%d", i)
		conn, err :=net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s closed\n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s open！！！\n",address)
	}
}