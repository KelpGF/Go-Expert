package main

import (
	"encoding/json"
	"fmt"
	"gofc-packages/services"
	"io"
	"net/http"
	"os"
)

func main() {
	// services.Run01()
	// services.Run02()
	// services.Run03()
	// searchByZipCode()

	// services.Run04()
	// services.Run05()
	// services.Run06()
	// services.Run07()
	// services.Run08()
	// services.Run09()
	// services.Run10()
	// services.Run11()
	// services.Run12()
	// services.Run13()
	services.Run14()

	// println("Choose a service to run.")
}

type SearchResult struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// compiles 01 - 03 lessons
func searchByZipCode() {
	zipCode := os.Args[1]
	if zipCode == "" {
		panic("Zip code is required")
	}

	url := "https://viacep.com.br/ws/" + zipCode + "/json/"

	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	var data SearchResult
	err = json.Unmarshal(res, &data)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("search-result.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("CEP: %s, Logradouro: %s, Complemento: %s, Bairro: %s, Localidade: %s, UF: %s, IBGE: %s, GIA: %s, DDD: %s, SIAFI: %s", data.Cep, data.Logradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf, data.Ibge, data.Gia, data.Ddd, data.Siafi))
	if err != nil {
		panic(err)
	}
}
