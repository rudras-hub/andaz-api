package entity

import "fmt"

// NilDialoguePointerError for pointer to a Dialogue object is nil.
var NilDialoguePointerError = fmt.Errorf("nil pointer reference to dialogue struct")

// DialogueWithIDNotFoundError when a Dialogue with required Id is not found in data repository.
type DialogueWithIDNotFoundError struct{
	id int64
}

// CharacterWithIDNotFoundError when a Character with required Id is not found in data repository.
type CharacterWithIDNotFoundError struct{
	id int64
}

// NewDialogueWithIDNotFoundError creates a new instance of DialogueWithIDNotFoundError.
func NewDialogueWithIDNotFoundError(id int64)*DialogueWithIDNotFoundError {
	return &DialogueWithIDNotFoundError{id: id}
}

// NewCharacterWithIDNotFoundError creates a new instance of CharacterWithIDNotFoundError.
func NewCharacterWithIDNotFoundError(id int64)*CharacterWithIDNotFoundError {
	return &CharacterWithIDNotFoundError{id: id}
}

// Error implements builtin.error interface on DialogueWithIDNotFoundError.
func (d *DialogueWithIDNotFoundError) Error() string{
	return fmt.Sprintf("dialogue with ID = %v not found", d.id)
}

// Error implements builtin.error interface on CharacterWithIDNotFoundError.
func (c *CharacterWithIDNotFoundError) Error() string{
	return fmt.Sprintf("character with ID = %v not found", c.id)
}