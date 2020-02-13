package Data

import (
	"encoding/json"
	"math/rand"
)

func GetRandomPokemonName(list string) string {
	var pokemonList[] string
	_ = json.Unmarshal([]byte(list), &pokemonList)
	randomPokemonNumber := rand.Intn(len(pokemonList))

	return pokemonList[randomPokemonNumber]
}

func GetRandomTrainerName(list string) string {
	var trainerlist[] string
	_ = json.Unmarshal([]byte(list), &trainerlist)
	randomTrainerNumber := rand.Intn(len(trainerlist))

	return trainerlist[randomTrainerNumber]
}

func GetRandomAttackName(list string) string {
	var attackList[] string
	_ = json.Unmarshal([]byte(list), &attackList)
	randomAttackNumber := rand.Intn(len(attackList))

	return attackList[randomAttackNumber]
}