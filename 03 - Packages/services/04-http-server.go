package services

import (
	"encoding/json"
	"io"
	"net/http"
)

func Run04() {
	http.HandleFunc("/", searchZipCodeHandler)
	// default mux of http package
	http.ListenAndServe(":8080", nil)
}

func searchZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	zipCodeParam := r.URL.Query().Get("zip-code")
	if zipCodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorBody := map[string]string{"error": "cep is required"}
		json.NewEncoder(w).Encode(errorBody)
		return
	}

	result, err := searchZipCode(zipCodeParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

	// this is the same as the commented code below
	json.NewEncoder(w).Encode(result)

	// encoded, err := json.Marshal(result)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }
	// w.Write(encoded)

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

func searchZipCode(zipCode string) (*SearchResult, error) {
	url := "https://viacep.com.br/ws/" + zipCode + "/json/"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var output SearchResult
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
