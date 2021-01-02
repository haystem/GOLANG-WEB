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

func init(){
  tpl = template.Must(template.ParseGlob("*.gohtml"))
}
func main (){
  
  pessoa1:= pessoas {Nome:"Alberto",Sobrenome:"Rocha", Idade:29}
  
  err := tpl.Execute(os.Stdout, pessoa1)
  if err != nil{
    log.Fatalln(err)
  }
 defer fmt.Println("Finalizado")
}