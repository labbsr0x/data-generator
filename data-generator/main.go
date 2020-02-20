package main

import (
	"./Cassandra"
	"os"
	"os/signal"
	"sync"
	"log"
)

func main()  {
	CassandraSession := Cassandra.Session
	Cassandra.CreateSchema()
	
	defer CassandraSession.Close()

	doneCh := make(chan struct{})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go Cassandra.ReadData(doneCh, &wg)
	wg.Add(1)
	go Cassandra.InsertData(doneCh, &wg)

    signalCh := make(chan os.Signal, 1)
    signal.Notify(signalCh, os.Interrupt)
	<-signalCh
	
	close(doneCh)
	wg.Wait()
	log.Print("Finished data generator")
	
}