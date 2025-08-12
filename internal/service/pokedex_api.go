package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/fanchann/charm-pokedex/internal/globals"
	"github.com/fanchann/charm-pokedex/internal/models"
)

func FetchPokemonList(offset, limit int) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?offset=%d&limit=%d", offset, limit)

		resp, err := http.Get(url)
		if err != nil {
			return globals.ErrorMsg{Err: err}
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return globals.ErrorMsg{Err: err}
		}

		var pokemonList models.PokemonList
		if err := json.Unmarshal(body, &pokemonList); err != nil {
			return globals.ErrorMsg{Err: err}
		}

		return globals.PokemonListMsg{Pokemons: pokemonList.Results}
	}
}

func FetchPokemonDetail(name string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", strings.ToLower(name))

		resp, err := http.Get(url)
		if err != nil {
			return globals.ErrorMsg{Err: err}
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return globals.ErrorMsg{Err: err}
		}

		var pokemon models.Pokemon
		if err := json.Unmarshal(body, &pokemon); err != nil {
			return globals.ErrorMsg{Err: err}
		}

		return globals.PokemonDetailMsg{Pokemon: pokemon}
	}
}

func SearchPokemon(query string) tea.Cmd {
	return func() tea.Msg {
		return FetchPokemonDetail(query)()
	}
}
