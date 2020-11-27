package main

import (
	"fmt"

	"github.com/minkj1992/go_nomad/dict/mydict"
)

func handleError(successMsg string, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(successMsg)
	}
}

func main() {
	dictionary := mydict.Dictionary{}
	// Add
	handleError("Add sucessed", dictionary.Add("leoo.j", "Nice guy"))
	// Search
	handleError(dictionary.Search("leoo.j"))
	// Update
	handleError("Update sucessed", dictionary.Update("leoo.j", "Super Nice guy"))
	// Search
	handleError(dictionary.Search("leoo.j"))
	// Delete
	handleError("Delete sucessed", dictionary.Delete("leoo.j"))
	// Search
	handleError(dictionary.Search("leoo.j"))
}
