package main

import (
    "fmt"
    "net"
    "os"
)

func main(){

    listener,err :=net.Listen("tcp","192.168.124.12:25")
    ServerHandleError(err,"net.listen")



    for {
        conn,e := listener.Accept()
        ServerHandleError(e,"listener.accept")
        go ChatWith(conn)
    }


}

func ServerHandleError(err error,when string) {
    if err != nil {
        fmt.Println(err, when)
        os.Exit(1)
    }
}

func ChatWith(conn net.Conn){

    buffer := make([]byte, 1024)
    for {

        n,err := conn.Read(buffer)
        ServerHandleError(err,"conn.read buffer")

        clientMsg := string(buffer[0:n])
        fmt.Printf("get message",conn.RemoteAddr(),clientMsg)

        if clientMsg != "im off" {
            conn.Write([]byte("read:"+clientMsg))
        } else {
            conn.Write([]byte("bye"))
            break
        }
    }
    conn.Close()
    fmt.Printf("Client disconnect",conn.RemoteAddr())

}