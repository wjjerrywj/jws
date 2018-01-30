
package channel

import (
	"net"
	"fmt"
	"../config"
)

type SCH struct{
    Pre *SCH
    Next *SCH
    state uint32
    SockIn net.Conn
    SockOut net.Conn
}

var serverIP string
var serverPort string

func Init(){
	var tstr *string
	if (config.GetServerIP(&tstr)>0){
		serverIP=*tstr
	}else{
		serverIP=""
	}
	if (config.GetServerPort(&tstr)>0){
		serverPort=*tstr
	}else{
		serverPort=""
	}
	fmt.Println("[Channel] Server=",serverIP,":",serverPort)
}


func InitSCH(sch *SCH){
	sch.state=0;
	sch.Pre=nil
	sch.Next=nil
	sch.SockIn=nil
	sch.SockOut=nil
}


func NewChannel(sch *SCH){
	go channelMainLoop(sch)    
}

func channelMainLoop(sch *SCH){

}