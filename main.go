package main

import (
	"fmt"
	// Pacote que possibilita utilizar json sem ter que definir uma struct
	"github.com/valyala/fastjson"
)

func main() {
	var parser fastjson.Parser

	jsonData := `{"id": 1, "nome": "José", "remoto":true, "notas": [9,8,9,7]}`

	aluno, err := parser.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("id=%d\n", aluno.GetInt("id"))
	fmt.Printf("nome=%s\n", aluno.GetStringBytes("nome"))
	fmt.Printf("remoto=%v\n", aluno.GetBool("remoto"))

	notas := aluno.GetArray("notas")

	for i, nota := range notas {
		fmt.Printf("Indice %d: %s\n", i, nota)
	}
}
