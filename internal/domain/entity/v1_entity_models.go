package entity

import "fmt"

// Character struct defines a character of the movie.
type Character struct{
	Id int64
	Name string
	Alias string
}

// Dialogue struct defines a dialogue of the movie.
type Dialogue struct{
	Id        int64
	Character Character
	Quote     string
}

// Characters is a slice of Character types.
type Characters []*Character

// Dialogues is a slice of Dialogue types.
type Dialogues []*Dialogue

// String implements fmt.Stringer interface on Dialogue pointer.
func(d *Dialogue)String() string{
	return fmt.Sprintf("%v: %v", d.Character.Name, d.Quote)
}