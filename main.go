package main

import (
	"./Cassandra"
)

func main()  {
	CassandraSession := Cassandra.Session
	
	defer CassandraSession.Close()

	Cassandra.InsertAttack()
	Cassandra.InsertTrainer()
	Cassandra.InsertPokemon()
	Cassandra.InsertAttack()
	Cassandra.InsertTrainer()
	Cassandra.InsertPokemon()
	Cassandra.InsertBattle()
}