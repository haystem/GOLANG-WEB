package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Println(err)
	}

// criar  o arquivo para jogar dados do tpl.gohtml para index.html
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("o erro: ", err)
	}
	defer nf.Close()
//executa o tpl direcionado para nf
	err = tpl.Execute(nf, nil)
	fmt.Println("fim")
}
