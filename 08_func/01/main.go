package main

import(
  "fmt"
  "text/template"
  "os"
  "log"
  "strings"
  )

var tpl *template.Template

// variavel que permite função no html
var fm = template.FuncMap{
   "uc" : strings.ToUpper ,
   "st" : str,
}

type pessoas struct{
  Nome string
  Sobrenome string
  Idade int
}

type carro struct {
  Modelo string
  Ano int
}

// criei uma função para chamar dentro da variavel fm
func str(s string) string{
  s2 := strings.TrimSpace(s)
  s2 = s[:3]
  return s2
}

//permite chamar a função e tratar erro

func init() {	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))}

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
  
  data := struct{
    Colab []pessoas
    Transp []carro
  }{
    conj,
    conj2,
  }
  
  err := tpl.Execute(os.Stdout, data)
  if err != nil{
    log.Fatalln(err)
  }
 defer fmt.Println("Finalizado")
}