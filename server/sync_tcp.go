package server

import (
	"io"
	"log"
	"net"


	"github.com/abhig003/redis/config"
)



func  Readcommand(c  net.Conn) (string,error){
	var buf []byte =make([]byte,1024)//declaring the buffer slice to read the command
	n , err := c.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Println("Error reading from connection: ", err)
		}
		return "", err
	}
	cmd:=string(buf[:n])
	return cmd,nil
}

func WriteResponse(c net.Conn, response string) error {
	_, err := c.Write([]byte(response))

	if err!=nil{
		return err
	}
	return err
}


func RunSyncTcp(){
  log.Println("starting a synchronous TCP server on ",config.Host,config.Port)
  var con_client int =0;
// listening to the configured host:port
  ln , err:=net.Listen("tcp",config.Host+":"+config.Port)

  if err!=nil{
	panic(err)
  }
    
  for{

	c, err := ln.Accept()
	if err != nil {
		log.Println("Error accepting connection: ", err)
		continue
	}
	con_client++//increament the client count
	log.Println("client connected with address",c.RemoteAddr().String(),"total clients connected ",con_client)
   // over the socket, continuously read the command and print it out
	for {
		cmd,err:=Readcommand(c)

		if err!=nil{
					c.Close()//connection closed by client
					con_client--//decreament the client count
					log.Println("client disconnected with address",c.RemoteAddr().String(),"total clients connected ",con_client)
					break
				
					
			}
		log.Println("Received command: ", cmd)
		err=WriteResponse(c,cmd)
		if err!=nil{
			log.Println("error occured during sending the response to client",c.RemoteAddr().String(),"error ",err)
			break
		}
	}

  }
}
