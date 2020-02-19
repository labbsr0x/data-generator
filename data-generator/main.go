package main

import (
	"./Cassandra"
)

func main()  {
	CassandraSession := Cassandra.Session
	Cassandra.CreateSchema()
	
	defer CassandraSession.Close()

	for { 
		Cassandra.InsertData()
		go Cassandra.ReadData()
	}
	
}