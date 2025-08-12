package globals

import "github.com/fanchann/charm-pokedex/internal/models"

type PokemonListMsg struct {
	Pokemons []models.PokemonListItem
}

type PokemonDetailMsg struct {
	Pokemon models.Pokemon
}

type ErrorMsg struct {
	Err error
}
