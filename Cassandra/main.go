package Cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
	"time"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

func init() {
	time.Sleep(10 * time.Second)
	var err error

	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "name"
	cluster.ProtoVersion = 3
	cluster.Consistency = gocql.One
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}