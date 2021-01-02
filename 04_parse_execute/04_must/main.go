package main

import(
    "fmt"
    "text/template"
    "log"
    "os"
  )
  
  var tpl *template.Template
  
  func init(){
    // Must faz a verificação do erro
    tpl = template.Must(template.ParseGlob("pag-teste/*.html"))
  }
  
  func main(){
    err:= tpl.Execute(os.Stdout,nil)
    if err!= nil {
      log.Fatalln(err)
    }
    fmt.Println("Finalizado")
  }