package main

import (
	"./Cassandra"
)

func main()  {
	CassandraSession := Cassandra.Session
	
	defer CassandraSession.Close()

	Cassandra.InsertTwitter()
}