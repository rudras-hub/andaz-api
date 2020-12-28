package main

import (
	"andaz-api/internal/database/adapter"
	"fmt"
)

func main(){
	c := adapter.GetCharacterRepositoryInstance()
	characterEntities, err := c.RetrieveAll()
	if err != nil{
		fmt.Println("Error = ", err)
		return
	}

	for _, char := range  characterEntities{
		fmt.Println("Name = ", char.Name, " Alias = ", char.Alias)
	}
}