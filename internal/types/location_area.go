package types

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationArea struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type LocAreaUrls struct {
	Next     string
	Previous string
}

func DefaultLocAreaUrls() LocAreaUrls {
	return LocAreaUrls{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "https://pokeapi.co/api/v2/location-area/",
	}
}

type LocationAreaPokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
