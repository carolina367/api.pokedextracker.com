package models

type GameFamilyDexNumber struct {
	GameFamilyID string `pg:",pk"`
	PokemonID    int    `pg:",pk"`
	DexNumber    int
}