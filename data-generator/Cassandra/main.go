package Cassandra

import (
	"github.com/gocql/gocql"
	"time"
	"math/rand"
	"log"
	"../Data"
	"sync"
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
		CREATE KEYSPACE cortex WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};`).Exec(); err != nil {
			log.Print(err)
	} 
	cluster.Keyspace = "cortex"
	cluster.ProtoVersion = 3
	cluster.Consistency = gocql.One
	
	log.Print("Cassandra init done")

}

func CreateSchema() {
	if err := Session.Query(`
		create table cortex.chunk(hash text, range blob, value blob, PRIMARY KEY(hash, range));`).Exec(); err != nil {
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
	time.Sleep(4 * time.Second)
	var counter int
	var name string
	if err := Session.Query(`SELECT COUNT(*), name FROM example.pokemon`).Scan(&counter, &name); err != nil {
		log.Fatal(err)
	}
	log.Printf("Read %d data", counter)
}

func InsertData(doneCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select{
		case <-doneCh:
			log.Print("Stopping data inserting")
			return
		default:
			if err := Session.Query(`
      			INSERT INTO cortex.chunk (hash, range, value) VALUES (?, ?, ?)`,
	  			randomString(100), randomString(4), randomString(1000)).Exec(); err != nil {
				log.Fatal(err)
			}
			log.Print("Data created!")
		}
	}
}

func ReadData(doneCh chan struct{}, wg *sync.WaitGroup) {
	log.Print("Reading data")
	defer wg.Done()
	for {
		select{
		case <-doneCh:
			log.Print("Stopping data reading")
			return
		default:
			var counter int
			var hash string
			if err := Session.Query(`SELECT COUNT(*), hash FROM cortex.chunk`).Scan(&counter, &hash); err != nil {
				log.Fatal(err)
			}
			log.Printf("Readed %d rows!", counter)
		}
	}
}

func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(33, 122))
    }
    return string(bytes)
}