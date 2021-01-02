package main

import(
  "fmt"
  "text/template"
  "os"
  "log"
  )

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("*.gohtml"))
}
func main (){
  
  slc := map[string]int{
          "Ana":33,
          "Carlos": 44,
          "Lylia": 34,
          "Joana": 22,
          "Alberto": 20,
  }
  
  err := tpl.Execute(os.Stdout, slc)
  if err != nil{
    log.Fatalln(err)
  }
 defer fmt.Println("Finalizado")
}