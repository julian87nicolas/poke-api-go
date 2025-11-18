package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPokemonFromPokeApiSuccess(t *testing.T) {
	c := require.New(t)

	pokemon, err := GetPokemonFromPokeApi("bulbasaur")
	c.NoError(err)

	body, err := os.ReadFile("samples/poke_api_readed.json")
	c.NoError(err)

	var excepcted models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &excepcted)
	c.NoError(err)

	c.Equal(excepcted, pokemon)

}
