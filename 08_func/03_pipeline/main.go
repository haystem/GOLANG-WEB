package main

import (
  "fmt"
  "time"
  "text/template"
  "os"
  "log"
  "math"
  )

var tpl  *template.Template

func formatarData(t time.Time) string{
 //  data := t.Format(time.Kitchen) 
  data := t.Format("2006-01-02")  //final 06 é referente ao ano / 01 - dia / 02 - mês
   return data
}

func soma(x int) int{
   return x+x
}

func quadrado(x int) float64{
   return math.Sqrt(float64(x))
}

func potencia(x float64) float64{
   return math.Pow(float64(x),3)
}


var fm = template.FuncMap{
    "cvtDT" : formatarData,
    "sqrt": quadrado,
    "soma":soma,
    "pot": potencia,
}

func init(){
  
tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

func main(){
  err := tpl.ExecuteTemplate(os.Stdout,"index.gohtml", 4)
  if err != nil{
    log.Fatalln(err)
  }
  fmt.Println("finalizado")
}