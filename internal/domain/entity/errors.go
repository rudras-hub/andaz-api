package entity

import "fmt"

// NilDialoguePointerError for pointer to a Dialogue object is nil.
var NilDialoguePointerError = fmt.Errorf("nil pointer reference to dialogue struct")

// DialogueWithIDNotFoundError when a Dialogue with required Id is not found in data repository.
type DialogueWithIDNotFoundError struct{
	id int64
}

// CharacterWithNameNotFoundError when a Character with required Id is not found in data repository.
type CharacterWithNameNotFoundError struct{
	name string
}

// NewDialogueWithIDNotFoundError creates a new instance of DialogueWithIDNotFoundError.
func NewDialogueWithIDNotFoundError(id int64)*DialogueWithIDNotFoundError {
	return &DialogueWithIDNotFoundError{id: id}
}

// NewCharacterWithNameNotFoundError creates a new instance of CharacterWithNameNotFoundError.
func NewCharacterWithNameNotFoundError(name string)*CharacterWithNameNotFoundError {
	return &CharacterWithNameNotFoundError{name: name}
}

// Error implements builtin.error interface on DialogueWithIDNotFoundError.
func (d *DialogueWithIDNotFoundError) Error() string{
	return fmt.Sprintf("dialogue with ID = %v not found", d.id)
}

// Error implements builtin.error interface on CharacterWithNameNotFoundError.
func (c *CharacterWithNameNotFoundError) Error() string{
	return fmt.Sprintf("character with name = %v not found", c.name)
}