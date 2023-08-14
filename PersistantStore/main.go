package main

import (
	"PersistantStore/model"
)

func main() {
	persistantStore := model.NewDBConnection("localdost:000")
	persistantStore.StartRouter()
}
