package main 

import(
  "fmt"
  "os"
  "text/template"
  "log"
  )
  
  //ponteiro para um tipo em template
  var tpl *template.Template
  
  func init(){
    tpl = template.Must(template.ParseFiles("index.html"))
  }
  
  func main(){
    er:= "Documento Provisorios"
    err := tpl.Execute(os.Stdout,er)
    if err != nil {
      log.Fatalln(err)
    }
    fmt.Println("realizado")
  }