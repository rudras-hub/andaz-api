package entity

// CharacterRepositoryReader defines an interface to retrieve
//Character or Characters from repository
type CharacterRepositoryReader interface{
	RetrieveAll()(Characters, error)
	RetrieveByName(characterName string)(*Character, error)
}

// DialogueRepositoryReader defines interface to retrieve
// Dialogue / Dialogues form repository
type DialogueRepositoryReader interface {
	RetrieveAllForCharacter(characterName string)(Dialogues, error)
	RetrieveRandomForCharacter(characterName string)(*Dialogue, error)
	RetrieveRandom()(*Dialogue, error)
}
