package main

import (
  "fmt"
  "os"
  "log"
  "text/template"
  )
  
  func main(){
    
    tpl, err := template.ParseFiles("tpl.gohtml")
    if err != nil {
      log.Println(err)
    }
    err = tpl.Execute(os.Stdout, nil)
       if err != nil {
      log.Println(err)
    }
    
    fmt.Println("fim")
  }