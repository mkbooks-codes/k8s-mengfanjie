package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	name := "testing"
	fmt.Printf("%d\n", name)
	fmt.Printf("%s\n", name, name)

	i := 0
	fmt.Println(i != 0 || i != 1)

	words := []string{"foo", "bar", "baz"}
	for _, word := range words {
		go func() {
			fmt.Println(word)
		}()
	}

	res, err := http.Get("https://www.spreadsheetdb.io/")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
