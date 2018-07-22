package main

import(
	"net"
	"log"
	//"fmt"
	"io"
	"os"
)

func main(){
	conn,err:=net.Dial("tcp",":14515")
	if err != nil{
		log.Fatal(err)
	}
	//fmt.Fprint(conn,"hello")
	defer conn.Close()
	go mustCopy(os.Stdout,conn)
	mustCopy(conn,os.Stdin)
}

func mustCopy(dst io.Writer,src io.Reader){
	if _,err := io.Copy(dst,src);err!=nil{
		log.Fatal(err)
	}
}