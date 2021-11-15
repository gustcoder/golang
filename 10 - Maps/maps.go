package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"aluno": {
			"nome": "Gustavo",
		},
		"curso": {
			"nome": "Engenharia",
		},
		"status": {
			"nome": "ativo",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["aluno"]["nome"])
	fmt.Println(usuario2["curso"]["nome"])

	// excluindo chave
	delete(usuario2, "status")
	fmt.Println(usuario2)

	// adicionando chave de volta
	usuario2["status"] = map[string]string{
		"nome": "inativo",
	}
	fmt.Println(usuario2)
}
