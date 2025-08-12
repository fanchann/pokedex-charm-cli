package service

import (
	"testing"

	"github.com/fanchann/charm-pokedex/internal/globals"
)

func TestFetchPokemonList(t *testing.T) {
	// Test case for fetching Pokemon list
	cmd := FetchPokemonList(0, 20)
	msg := cmd()

	switch msg := msg.(type) {
	case globals.PokemonListMsg:
		if len(msg.Pokemons) == 0 {
			t.Error("Expected Pokemon list to have items")
		}
	case globals.ErrorMsg:
		t.Errorf("Unexpected error: %v", msg.Err)
	default:
		t.Error("Unexpected message type")
	}
}

func TestFetchPokemonDetail(t *testing.T) {
	// Test case for fetching Pokemon detail
	cmd := FetchPokemonDetail("pikachu")
	msg := cmd()

	switch msg := msg.(type) {
	case globals.PokemonDetailMsg:
		if msg.Pokemon.Name != "pikachu" {
			t.Errorf("Expected pokemon name to be 'pikachu', got '%s'", msg.Pokemon.Name)
		}
	case globals.ErrorMsg:
		t.Errorf("Unexpected error: %v", msg.Err)
	default:
		t.Error("Unexpected message type")
	}
}

func TestSearchPokemon(t *testing.T) {
	// Test search by ID
	cmd := SearchPokemon("25")
	msg := cmd()

	switch msg := msg.(type) {
	case globals.PokemonDetailMsg:
		if msg.Pokemon.ID != 25 {
			t.Errorf("Expected pokemon ID to be 25, got %d", msg.Pokemon.ID)
		}
	case globals.ErrorMsg:
		t.Errorf("Unexpected error: %v", msg.Err)
	default:
		t.Error("Unexpected message type")
	}
}
