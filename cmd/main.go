package main

import (
	"andaz-api/internal/database/adapter"
	"andaz-api/internal/logging"
)

func main(){
	l := logging.NewStdLoggerExtension()
	c := adapter.GetCharacterRepositoryInstance()
	characterEntities, err := c.RetrieveAll()
	if err != nil{
		l.Errorf("error = %v", err)
		return
	}

	for _, char := range  characterEntities{
		l.Infof("Name = %v, Alias = %v", char.Name, char.Alias)
	}
}