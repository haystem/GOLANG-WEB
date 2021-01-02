package main

import (
  "fmt"
  "io"
  "os"
  "log"
  "strings"
  )

func main(){
  name:= "Alberto Rocha"
  str := fmt.Sprint(`
        <!DOCTYPE html>
      <html>
        <head>
          <meta http-equiv="content-type" content="text/html; charset=utf-8" />
          <meta name="" content="">
          <title> Hello Dear</title>
        </head>
        <body>
         <h1>` + name+ `</h1>
        </body>
      </html>
      
  `)
 page, err := os.Create("main1.html")
 if err != nil {
   log.Println(err)
 }
 defer page.Close()
 io.Copy(page, strings.NewReader(str))
 fmt.Println("tudo certo")
}