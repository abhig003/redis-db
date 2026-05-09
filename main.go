package main

import (
	"flag"
	"log"
	"github.com/abhig003/redis/config"
	"github.com/abhig003/redis/server"
)




func setup(){
	flag.StringVar(&config.Host,"host",config.Host,"Host to listen on")
	flag.StringVar(&config.Port,"port",config.Port,"Port to listen on")
	flag.Parse()
}


func main(){
	setup()
	log.Printf("Starting server on %s:%s",config.Host,config.Port)
	log.Printf("rolling the redis db")
	 server.RunSyncTcp()
}