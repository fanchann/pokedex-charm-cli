package models

// Pokemon represents a Pokemon from the API
type Pokemon struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	Height    int              `json:"height"`
	Weight    int              `json:"weight"`
	BaseExp   int              `json:"base_experience"`
	Types     []PokemonType    `json:"types"`
	Stats     []PokemonStat    `json:"stats"`
	Species   PokemonSpecies   `json:"species"`
	Sprites   PokemonSprites   `json:"sprites"`
	Abilities []PokemonAbility `json:"abilities"`
}

// PokemonType represents a Pokemon type
type PokemonType struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

// PokemonStat represents a Pokemon stat
type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}

// PokemonSpecies represents a Pokemon species
type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonSprites represents Pokemon sprites
type PokemonSprites struct {
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
}

// PokemonAbility represents a Pokemon ability
type PokemonAbility struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
	Ability  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"ability"`
}

// PokemonListItem represents a Pokemon in the list
type PokemonListItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonList represents the API response for Pokemon list
type PokemonList struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []PokemonListItem `json:"results"`
}

func (p PokemonListItem) FilterValue() string { return p.Name }

func (p PokemonListItem) Title() string { return p.Name }

func (p PokemonListItem) Description() string { return "Pokemon" }
