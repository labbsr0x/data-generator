package Cassandra

import (
	"github.com/gocql/gocql"
	"time"
	"math/rand"
	"log"
	"../Data"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

func init() {
	time.Sleep(30 * time.Second)
	var err error

	cluster := gocql.NewCluster("cassandra")
	
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	if err := Session.Query(`
		CREATE KEYSPACE example WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};`).Exec(); err != nil {
			log.Print(err)
	} 
	cluster.Keyspace = "name"
	cluster.ProtoVersion = 3
	cluster.Consistency = gocql.One
	
	log.Print("Cassandra init done")

}

func CreateSchema() {
	if err := Session.Query(`
		create table example.pokemon(id UUID, name text, level int, attack text, owner_name text, PRIMARY KEY(id));`).Exec(); err != nil {
			log.Print(err)
	} 

	if err := Session.Query(`
		create table example.trainer(id UUID, name text, badge_name text, PRIMARY KEY(id));`).Exec(); err != nil {
			log.Print(err)
	} 

	if err := Session.Query(`
		create table example.battle(id UUID, first_pokemon UUID, second_pokemon UUID, PRIMARY KEY(id, first_pokemon, second_pokemon));`).Exec(); err != nil {
			log.Print(err)
	} 

	if err := Session.Query(`
		create table example.attack(id UUID, attack_name text, damage int, PRIMARY KEY(id));`).Exec(); err != nil {
			log.Print(err)
	} 
	log.Print("Schema Created")
}

func InsertPokemon() {
	pokemonName := Data.GetRandomPokemonName(Data.PokemonList)
	uuid := gocql.TimeUUID()
	if err := Session.Query(`
      INSERT INTO example.pokemon (id, name, level, attack, owner_name) VALUES (?, ?, ?, ?, ?)`,
	  uuid, pokemonName, rand.Intn(100), FindAttack(), FindTrainer()).Exec(); err != nil {
      	log.Fatal(err)
	} 
	log.Print("Pokemon Created. uuid: ", uuid)
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
	trainerName := Data.GetRandomTrainerName(Data.TrainerList)
	if err := Session.Query(`
      INSERT INTO example.trainer (id, name, badge_name) VALUES (?, ?, ?)`,
	  gocql.TimeUUID(), trainerName, "Hive Badge").Exec(); err != nil {
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
		log.Print("Pokemon id: ", id)
		battleParticipants = append(battleParticipants, id)
	}
	
	if err := Session.Query(`INSERT INTO example.battle (id, first_pokemon, second_pokemon) VALUES (?, ?, ?)`, 
		id, battleParticipants[0], battleParticipants[1]).Exec(); err != nil {
			log.Fatal(err)
	}

	log.Print("Battle Created")
}

func InsertAttack() {
	attackName := Data.GetRandomAttackName(Data.AttackList)
	if err := Session.Query(`
		INSERT INTO example.attack (id, attack_name, damage) VALUES (?, ?, ?)`,
		gocql.TimeUUID(), attackName, rand.Intn(100)).Exec(); err != nil {
      		log.Fatal(err)
	} 
	log.Print("Attack Created")
}

func GetPokemon() {
	var counter int
	var name string
	time.Sleep(10 * time.Second)
	if err := Session.Query(`SELECT COUNT(*), name FROM example.pokemon`).Scan(&counter, &name); err != nil {
		log.Fatal(err)
	}
	log.Printf("Read %d data", counter)
}