package main

import (
	"./Cassandra"
)

func main()  {
	

	/*var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 3
    cluster.Keyspace = "name"
    cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	
	if err != nil {
		panic(err)
	}
	defer session.Close()
	fmt.Println("cassandra init done")
	fmt.Printf("hello world")*/

	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()
	
}
