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
  
  slc := []string{"Ana","Carolina","Resende","Fagundes"}
  
  err := tpl.Execute(os.Stdout, slc)
  if err != nil{
    log.Fatalln(err)
  }
 defer fmt.Println("Finalizado")
}