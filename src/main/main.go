
package main

import (
    "fmt"
    "os"
    "net"
    "../channel"
    "../config"
)
/*
type SCH struct{
    pre *SCH
    next *SCH
    sock net.Conn
}*/

const DEFAULT_PORT  string ="1888"

//var myConfig *(map[string]string)
var sch *(channel.SCH)

func checkError(err error){
    if err!=nil {
        fmt.Println("Error: %s",err.Error())
        os.Exit(1)
    }
}

func main(){
	var port string=DEFAULT_PORT
	//myConfig=config.ConfigInit("jws_config.json")
//	myConfig=make(map[string]string)
	config.Init("../jws_config.json")
	channel.Init()
	var tstr *string
	if config.GetLocalPort(&tstr)>0 {
		port=*tstr
	}
    sch=new (channel.SCH)
    channel.InitSCH(sch)
    fmt.Println("[Main] Start listenning Port:"+port+" ...\n")
    listen_sock, err :=net.Listen("tcp",":"+port)
    checkError(err)
    defer listen_sock.Close()
    for {
        new_sock, err :=listen_sock.Accept()
        if err!=nil {
            fmt.Println("Accept error: %s",err.Error())
            continue
        }
        newClient(new_sock)
    }
}

func newClient(newSock net.Conn){
	fmt.Println("[Main]New client in... ")
    channel.NewChannel(newSCH(newSock))
}

func newSCH(newSock net.Conn) *(channel.SCH) {
    tsch:=sch
    for ;tsch!=nil; {
        if tsch!=nil {
            if (tsch.SockIn==newSock) {
                return tsch
            }
            if tsch.Next==nil {
                break;
            }
        }
        tsch=tsch.Next
    }
    ttsch:= new(channel.SCH)
    channel.InitSCH(ttsch)
    ttsch.SockIn=newSock
    ttsch.Pre=tsch
    tsch.Next=ttsch
    return ttsch
}