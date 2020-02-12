package Cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
	"time"
	"math/rand"
	"log"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

func init() {
	time.Sleep(40 * time.Second)
	var err error

	cluster := gocql.NewCluster("cassandra")
	
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	if err := Session.Query(`
		CREATE KEYSPACE example WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};`).Exec(); err != nil {
			log.Fatal(err)
	} 
	cluster.Keyspace = "name"
	cluster.ProtoVersion = 3
	cluster.Consistency = gocql.One
	
	log.Print("Cassandra init done")

}

func CreateSchema() {
	if err := Session.Query(`
		create table example.pokemon(id UUID, name text, level int, attack text, owner_name text, PRIMARY KEY(id));`).Exec(); err != nil {
			log.Fatal(err)
	} 

	if err := Session.Query(`
		create table example.trainer(id UUID, name text, badge_name text, PRIMARY KEY(id));`).Exec(); err != nil {
			log.Fatal(err)
	} 

	if err := Session.Query(`
		create table example.battle(id UUID, first_pokemon UUID, second_pokemon UUID, PRIMARY KEY(id, first_pokemon, second_pokemon));`).Exec(); err != nil {
			log.Fatal(err)
	} 

	if err := Session.Query(`
		create table example.attack(id UUID, attack_name text, damage int, PRIMARY KEY(id));`).Exec(); err != nil {
			log.Fatal(err)
	} 
	log.Print("Schema Created")
}

func InsertPokemon() {
	if err := Session.Query(`
      INSERT INTO example.pokemon (id, name, level, attack, owner_name) VALUES (?, ?, ?, ?, ?)`,
	  gocql.TimeUUID(), "Pikachu", rand.Intn(100), FindAttack(), FindTrainer()).Exec(); err != nil {
      	log.Fatal(err)
	} 
	log.Print("Pokemon Created")
}

func FindAttack() string{
	var attackName string
	if err := Session.Query(`
		SELECT attack_name FROM example.attack LIMIT 1
		`).Scan(&attackName); err != nil {
			log.Fatal(err)
	}
	return attackName
}
func FindTrainer() string{
	var trainerName string
	if err := Session.Query(`
		SELECT name FROM example.trainer LIMIT 1
	`).Scan(&trainerName); err != nil {
		log.Fatal(err)
	}
	return trainerName
}

func InsertTrainer() {
	if err := Session.Query(`
      INSERT INTO example.trainer (id, name, badge_name) VALUES (?, ?, ?)`,
	  gocql.TimeUUID(), "Ash", "Hive Badge").Exec(); err != nil {
		log.Fatal(err)
	}
	log.Print("Trainer Created")
}

func InsertBattle() {
	var id gocql.UUID
	var battleParticipants[] gocql.UUID

	//select last pokemons registered to add to battle
	iter := Session.Query(`SELECT id FROM example.pokemon LIMIT 2`).Iter()
	for iter.Scan(&id) {
		fmt.Printf("pokemon id: %v", id)
		battleParticipants = append(battleParticipants, id)
    }

	fmt.Printf("%v", battleParticipants[0])
	if err := Session.Query(`INSERT INTO example.battle (id, first_pokemon, second_pokemon) VALUES (?, ?, ?)`, 
		id, battleParticipants[0], battleParticipants[1]).Exec(); err != nil {
			log.Fatal(err)
	}

	log.Print("Battle Created")
}

func InsertAttack() {
	if err := Session.Query(`
		INSERT INTO example.attack (id, attack_name, damage) VALUES (?, ?, ?)`,
		gocql.TimeUUID(), "Thundershock", rand.Intn(100)).Exec(); err != nil {
      		log.Fatal(err)
	} 
	log.Print("Attack Created")
}