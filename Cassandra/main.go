package Cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
	"time"
	"math/rand"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

func init() {
	time.Sleep(60 * time.Second)
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
		create table example.pokemon(id UUID, name text, level int, atack_name text, PRIMARY KEY(id));`).Exec(); err != nil {
		errs = append(errs, err.Error())
	} 

	if err := Session.Query(`
		create table example.trainer(id UUID, name text, badge_name text, PRIMARY KEY(id));`).Exec(); err != nil {
		errs = append(errs, err.Error())
	} 

	if created {
		fmt.Println("SUCESS")
	}else{
		fmt.Println("errors", errs)
	}
	fmt.Println("cassandra init done")

}

func InsertPokemon() {
	var errs []string

	var created bool = false
	if err := Session.Query(`
      INSERT INTO example.pokemon (id, name, level, atack_name) VALUES (?, ?, ?, ?)`,
	  gocql.TimeUUID(), "Pikachu", rand.Intn(100), "Thundershock").Exec(); err != nil {
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

func InsertTrainer() {
	var errs []string

	var created bool = false
	if err := Session.Query(`
      INSERT INTO example.trainer (id, name, badge_name) VALUES (?, ?, ?)`,
	  gocql.TimeUUID(), "Ash", "Hive Badge").Exec(); err != nil {
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