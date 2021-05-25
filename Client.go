package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func ClientHandleError(err error, when string){
    if err != nil {
        fmt.Println(err, when)
        os.Exit(1)
    }

}


func main(){

    conn, err := net.Dial("tcp","192.168.124.12:25")
    ClientHandleError(err, "client conn error")

    buffer := make([]byte,1024)


    reader := bufio.NewReader(os.Stdin)

    for {
        lineBytes,_,_ := reader.ReadLine()
        conn.Write(lineBytes)
        n,err := conn.Read(buffer)
        ClientHandleError(err, "client read error")

        serverMsg := string(buffer[0:n])
        fmt.Printf("Server msg",serverMsg)
        if serverMsg == "bye" {
            break
        }

    }

}