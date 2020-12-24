package entity

// CharacterRepositoryReader defines an interface to retrieve
//Character or Characters from repository
type CharacterRepositoryReader interface{
	RetrieveAll()(Characters, error)
	RetrieveByID(characterID int64)(*Character, error)
}

// DialogueRepositoryReader defines interface to retrieve
// Dialogue / Dialogues form repository
type DialogueRepositoryReader interface {
	RetrieveAllForCharacter(characterID int64)(Dialogues, error)
	RetrieveRandomForCharacter(characterID int64)(*Dialogue, error)
	RetrieveRandom()(*Dialogue, error)
}
