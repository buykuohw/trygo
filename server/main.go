package main

import(
	"net"
	"log"
	"fmt"
	"io"
	//"os"
	"bufio"
)

func main(){
	listener,err:=net.Listen("tcp",":14515")
	if err!=nil{
		log.Fatal(err)
	}
	for{
		conn,err:=listener.Accept()
		if err!= nil{
			log.Print(err)
			continue
		}
		go handleConn(conn)
		//mustCopy(os.Stdout,conn)
	}
}
func handleConn(c net.Conn){
	input := bufio.NewScanner(c)
	for input.Scan(){
		resp := input.Text()
		fmt.Println(resp)
		fmt.Fprintln(c,"Server :" + resp)
	}
	c.Close()
}
func mustCopy(dst io.Writer,src io.Reader){
	if _,err := io.Copy(dst,src);err!=nil{
		log.Fatal(err)
	}
}