package main

import (
  "fmt"
  "time"
  "text/template"
  "os"
  "log"
  )

var tpl  *template.Template

func formatarData(t time.Time) string{
 //  data := t.Format(time.Kitchen) 
  data := t.Format("2006-01-02")  //final 06 é referente ao ano / 01 - dia / 02 - mês
   return data
}

var fm = template.FuncMap{
    "cvtDT" : formatarData,
}

func init(){
  
tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

func main(){
  err := tpl.ExecuteTemplate(os.Stdout,"index.gohtml", time.Now())
  if err != nil{
    log.Fatalln(err)
  }
  fmt.Println("finalizado")
}