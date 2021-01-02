package main

import (
  "fmt"
  "text/template"
  "log"
  "os"
  )
  
  func main(){
    tpl, err := template.ParseFiles("mosca.gohtml","vespa.gohtml","uranos.gohtml")
    if err != nil{
      log.Println(err)
    }
     //executa o primeiro gohtml 
    err = tpl.Execute(os.Stdout, nil)
      if err != nil{
      log.Println(err)
    }
    
     //executa especifico gohtml que eu escolher
    err = tpl.ExecuteTemplate(os.Stdout,"mosca.gohtml", nil)
    if err != nil{
      log.Println(err)
    }
     //executa especifico gohtml que eu escolher
       err = tpl.ExecuteTemplate(os.Stdout,"vespa.gohtml", nil)
    if err != nil{
      log.Println(err)
    }
    
    //executa especifico gohtml que eu escolher
       err = tpl.ExecuteTemplate(os.Stdout,"uranos.gohtml", nil)
    if err != nil{
      log.Println(err)
    }
    
    fmt.Println("FIM")
  }