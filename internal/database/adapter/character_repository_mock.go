package adapter

import (
	"andaz-api/internal/domain/entity"
	"sync"
)

// CharacterRepositoryMap defines the type for mock character repository.
type characterRepositoryMap map[string]*entity.Character

type CharacterRepositoryMock struct{
	characterMap characterRepositoryMap
}

// characterRepositoryInstance is the singleton instance of mock character repository.
var characterRepositoryInstance *CharacterRepositoryMock

// once for singleton creation.
var once sync.Once

// mutex for locking and unlocking on repository operations.
var mutex = &sync.Mutex{}

// GetCharacterRepositoryInstance initializes the singleton instance the first time,
// else returns the existing instance.
func GetCharacterRepositoryInstance() *CharacterRepositoryMock {
	once.Do(func() {
		characterRepositoryInstance = initializeCharacterRepoMock()
	})
	return characterRepositoryInstance
}

// initializeCharacterRepoMock initializes the mock character repository.
// Called once in get instance method.
func initializeCharacterRepoMock() *CharacterRepositoryMock {
	mockRepo := &CharacterRepositoryMock{}
	charNames := GetAllNames()
	charAliases := GetNameVsAliasMap()
	mockRepo.characterMap = characterRepositoryMap{}

	for _, name := range charNames{
		characterObject := &entity.Character{Name: name}
		if _, aliasExists := charAliases[name]; aliasExists{
			characterObject.Alias = charAliases[name]
		}
		mockRepo.characterMap[name] = characterObject
	}

	return mockRepo
}

// RetrieveAll retrieves all entity.Characters from mock repository.
// Implement entity.CharacterRepositoryReader interface on CharacterRepositoryMock
func (c *CharacterRepositoryMock)RetrieveAll()(entity.Characters, error){
	var chars []*entity.Character
	mutex.Lock()
	defer mutex.Unlock()
	for _, character := range c.characterMap{
		chars = append(chars, character)
	}
	return chars, nil
}

// RetrieveByName retrieves the entity.Character object for the given name.
//Implement entity.CharacterRepositoryReader interface on CharacterRepositoryMock
func (c *CharacterRepositoryMock)RetrieveByName(characterName string)(*entity.Character, error){
	mutex.Lock()
	defer mutex.Unlock()

	if _, nameExists := c.characterMap[characterName]; !nameExists{
		return nil, entity.NewCharacterWithNameNotFoundError(characterName)
	}

	return c.characterMap[characterName], nil
}