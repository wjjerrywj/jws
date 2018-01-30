package config

import (
	"os"
	"encoding/json"
	"fmt"
)

const KEY_PORT 		string = "LocalPort"
const KEY_SERV_IP	string = "ServerIP"
const KEY_SERV_PORT	string = "ServerPort"
const KEY_SERV_NAME	string = "ServerName"


var state uint32
var myMap map[string]string

func Init(fn string) {
	state=0
	finfo,err:=os.Stat(fn)
	if (err!=nil){
		fmt.Println("[Config] File("+fn+") not found!")
		return
	}
	filelen:=finfo.Size()
	if (filelen<8){
		fmt.Println("[Config] File("+fn+") have no content!")
		return
	}
	file,err:=os.Open(fn)
	if (err!=nil){
		fmt.Println("[Config] Open file \""+fn +"\"fail...")
		return
	}
	state=1
	//var buffer []uint8
	buffer:=make([]byte,filelen)
	n,err:=file.Read(buffer)
	if (err!=nil){
		fmt.Println("[Config] Read file("+fn+") fail...")
		file.Close()
		state=0
		return
	}
	file.Close()
//	fmt.Print("[Config] Read ",n," bytes")
//	fmt.Println(buffer)
	if n>8 {
		//keyval:=make([]key_value,256)
		//var kvs keyvals
		myMap=make(map[string]string,256)
		err:=json.Unmarshal(buffer,&myMap)
		if (err!=nil){
			fmt.Println("[Config] Json Unmarshal err:",err)
		}
		state=3
		fmt.Printf("[Config] Read config :%v",myMap)
		fmt.Println()
/*		for key,value:=range myMap {
			fmt.Println("[Config] ",key,"=",value)
		}
		fmt.Println(myMap[KEY_PORT])*/
	}
}


func GetLocalPort(out **string) int{
	if (state<3){
		return 0
	}
	tstr:=myMap[KEY_PORT]
	if (len(tstr)>0){
		*out=&tstr
		return 1
	}
	return 0
}

func GetLocalPortString() string{
	if (state>=3){
		return myMap[KEY_PORT]
//		if (len(tstr)>0){
//			return *tstr
//		}
	}
	return ""
}

func GetServerIP(out **string) int{
	if (state<3){
		return 0
	}
	tstr:=myMap[KEY_SERV_IP]
	if (len(tstr)>0){
		*out=&tstr
		return 1
	}
	return 0
}

func GetServerIPstring() string{
	if (state>=3){
		return myMap[KEY_SERV_IP]
	}
	return ""
}


func GetServerPort(out **string) int{
	if (state<3){
		return 0
	}
	tstr:=myMap[KEY_SERV_PORT]
	if (len(tstr)>0){
		*out=&tstr
		return 1
	}
	return 0
}

func GetServerPortString() string {
	if (state>=3){
		return myMap[KEY_SERV_PORT]
	}
	return ""
}


