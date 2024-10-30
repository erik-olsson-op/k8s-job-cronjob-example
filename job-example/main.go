package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "https://pokeapi.co/api/v2/pokemon/lugia", nil)
	r, _ := http.DefaultClient.Do(req)

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}
