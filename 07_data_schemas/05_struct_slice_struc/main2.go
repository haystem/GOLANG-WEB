package main

import(
  "fmt"
  "text/template"
  "os"
  "log"
  )

var tpl *template.Template

type pessoas struct{
  Nome string
  Sobrenome string
  Idade int
}

type carro struct {
  Modelo string
  Ano int
}

type item struct{
  ps pessoas
  cr carro
}

func init(){
  tpl = template.Must(template.ParseGlob("*.gohtml"))
}
func main (){
  
  pessoa1:= pessoas {Nome:"Alberto",Sobrenome:"Rocha", Idade:29}
  pessoa2:= pessoas {Nome:"Ana", Sobrenome:"Carla",Idade: 19}
  pessoa3:= pessoas {Nome:"Adoberto",Sobrenome: "Lilia",Idade: 22}
  pessoa4:= pessoas {Nome:"Anaberto", Sobrenome:"Santos",Idade: 18}
  pessoa5:= pessoas {Nome:"Anb", Sobrenome:"Fabio",Idade: 99}
  
  carro1 := carro{Modelo:"Wolskvagen",Ano:2012}
  carro2 := carro{Modelo:"Fox",Ano:2017}
  carro3 := carro{Modelo:"Toyota",Ano:2020} 
  
  conj := []pessoas{pessoa1,pessoa2,pessoa3,pessoa4,pessoa5}
  conj2 :=[]carro{carro1,carro2,carro3}
  
  
  data :={conj,conj2}

  err := tpl.Execute(os.Stdout, data)
  if err != nil{
    log.Fatalln(err)
  }
 defer fmt.Println("Finalizado")
}