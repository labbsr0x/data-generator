package Data

import (
	"encoding/json"
	"math/rand"
)

var pokemonList[] string

func GetRandomPokemonName(list string) string {
	_ = json.Unmarshal([]byte(list), &pokemonList)
	randomPokemonNumber := rand.Intn(len(pokemonList))

	return pokemonList[randomPokemonNumber]
}