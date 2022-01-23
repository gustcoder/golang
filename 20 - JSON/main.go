package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Gato struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func json_encode(dados []byte) {
	fmt.Println(bytes.NewBuffer(dados))
}

func main() {
	gato := Gato{"Lisa", "SRD", 1}
	//fmt.Println(gato)

	gatoJson, err := json.Marshal(gato)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(gatoJson)  # retorno em []bytes uint8 - usa-se bytes.NewBuffer para converter os dados
	fmt.Println(bytes.NewBuffer(gatoJson))

	gato2 := map[string]string{
		"Nome":  "Bart",
		"Raca":  "SRD",
		"Idade": "1",
	}

	gato2Json, err := json.Marshal(gato2)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(gato2Json) # retorno em []bytes uint8 - usa-se bytes.NewBuffer para converter os dados
	json_encode(gato2Json)

	gatoEmJson := `{"Idade":1,"Nome":"Ori","Raca":"SRD"}`

	var g Gato

	// precisa fazer o cast para byte, o inverso do que acontece com o Marshal
	if err := json.Unmarshal([]byte(gatoEmJson), &g); err != nil {
		log.Fatal(err)
	}

	fmt.Println(g)
}
