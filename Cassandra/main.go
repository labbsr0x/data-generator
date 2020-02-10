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
	
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	var errs []string
	var created bool = false
	if err := Session.Query(`
		CREATE KEYSPACE example WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};`).Exec(); err != nil {
		errs = append(errs, err.Error())
	} 
	cluster.Keyspace = "name"
	cluster.ProtoVersion = 3
	cluster.Consistency = gocql.One
	
	if err := Session.Query(`
		create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));`).Exec(); err != nil {
		errs = append(errs, err.Error())
	} 
	
	if err := Session.Query(`
		create index on example.tweet(timeline);`).Exec(); err != nil {
		errs = append(errs, err.Error())
	} else {
		created = true
	}

	if created {
		fmt.Println("SUCESS")
	}else{
		fmt.Println("errors", errs)
	}
	fmt.Println("cassandra init done")

}

func InsertUser() {
	var gocqlUuid gocql.UUID
	var errs []string

	var created bool = false
	if err := Session.Query(`
      INSERT INTO example.tweet (timeline, id, text) VALUES (?, ?, ?)`,
	  "me", gocqlUuid, "hello world").Exec(); err != nil {
      errs = append(errs, err.Error())
    } else {
      created = true
	}
	
	if created {
		fmt.Println("SUCESS")
	}else{
		fmt.Println("errors", errs)
	}
}