package main

import (
	"./Cassandra"
)



func main()  {
	CassandraSession := Cassandra.Session
	Cassandra.CreateSchema()
	
	defer CassandraSession.Close()

	for { 
		Cassandra.InsertAttack()
		Cassandra.InsertTrainer()
		Cassandra.InsertPokemon()
		Cassandra.InsertAttack()
		Cassandra.InsertTrainer()
		Cassandra.InsertPokemon()
		Cassandra.InsertBattle()
	}
}