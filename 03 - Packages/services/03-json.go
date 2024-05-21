package services

import (
	"encoding/json"
	"fmt"
	"os"
)

func Run03() {
	serialize()
	encoder()
	deserialize()
}

type AccountBank struct {
	Number  int `json:"n"` // tag to change the field Number in the JSON to "n"
	Balance int `json:"b"` // tag to change the field Balance in the JSON to "b"
	Agency  int `json:"-"` // tag to ignore the field Agency in the JSON
}

func serialize() []byte {
	account := AccountBank{Number: 123, Balance: 1000, Agency: 456}

	// serialize to JSON and return as bytes
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}

	fmt.Println("Account Struct", account)
	fmt.Println("Account JSON Bytes", res)
	fmt.Println("Account JSON Bytes to String", string(res))

	return res
}

func encoder() {
	account := AccountBank{Number: 123, Balance: 1000, Agency: 456}

	// serialize to JSON and write to somewhere
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(account)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}
}

// deserialize from JSON
func deserialize() {
	jsonStr := `{"n":123,"b":1000}`
	jsonBytes := []byte(jsonStr)
	// jsonBytes = serialize()

	account := AccountBank{}

	// deserialize from JSON bytes and fill the struct fields using the pointer
	err := json.Unmarshal(jsonBytes, &account)
	if err != nil {
		panic(err)
	}

	fmt.Println("Account Number ", account.Number)
	fmt.Println("Account Balance ", account.Balance)
	fmt.Println("Account Agency ", account.Agency)
}
