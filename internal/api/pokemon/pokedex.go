package pokemon

import "github.com/BrownieBrown/pokedex/internal/models"

type Pokedex struct {
	Index map[string]models.Pokemon
}

func NewPokedex() *Pokedex {
	p := Pokedex{
		Index: make(map[string]models.Pokemon),
	}
	return &p
}

func (p *Pokedex) Add(pokemon models.Pokemon) {
	p.Index[pokemon.Name] = pokemon
}

func (p *Pokedex) Get(name string) (models.Pokemon, bool) {
	pokemon, ok := p.Index[name]
	return pokemon, ok
}

func (p *Pokedex) Remove(name string) {
	delete(p.Index, name)
}

func (p *Pokedex) GetAll() map[string]models.Pokemon {
	return p.Index
}
