package main

import (
		"fmt"
		"../config"
		"time"
		"net"
		"os"
		)
		
const DEFAULT_PORT  string ="1888"	

var serverIP,serverPort,localPort string
var sock_server net.Conn
var isConnectedServer bool=false

func main(){
	config.Init("../jws_config.json")
	serverIP=config.GetServerIPstring()
	serverPort=config.GetServerPortString()
	localPort=config.GetLocalPortString()
	if len(serverIP)==0 || len(serverPort)==0 || len(localPort)==0 {
		fmt.Println("[main] config.Init fail...")
		time.Sleep(time.Second*2)
		os.Exit(1)
	}
	go toServer_loop()
	for {
	}
}

func toServer_loop() {
	for ;!isConnectedServer; {
		fmt.Println("[main] Start connecting to server:"+serverIP+"["+serverPort+"]......")
		tsock,err:=net.Dial("tcp",serverIP+":"+serverPort)
		if (err!=nil){
			isConnectedServer=false
			tsock=nil
			fmt.Println("[main] Connect to server("+serverIP+":"+serverPort+") error:"+err.Error())
			time.Sleep(time.Second*10)
		}
		sock_server=tsock
		isConnectedServer=true
	}
}

