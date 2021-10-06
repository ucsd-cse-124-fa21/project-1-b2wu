package main

import (
        "fmt"
        "log"
        "net/http"
        "time"
        "math/rand"
        "strconv"
)

func main() {

        http.HandleFunc("/", handler)
        http.HandleFunc("/gendata", gendata)
        log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func gendata(w http.ResponseWriter, r *http.Request){
        var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOP~!@#$%^&* `;',./><QRSTUVWXYZ")
        rand.Seed(time.Now().UnixNano())
        inputArg, err := r.URL.Query()["numBytes"]
        intVar, er := strconv.Atoi(inputArg[0])
        if !err {
               log.Print("null input")
        }

        if er != nil {
                log.Print("error")
        }
  
        b := make([]rune, intVar)
        for i := range b {
                b[i] = letters[rand.Intn(len(letters))]
        }
        fmt.Fprintf(w, "%s\n",string(b))
}

func handler(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
        for k, v := range r.Header {
                fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
        }
        fmt.Fprintf(w, "Host = %q\n", r.Host)
        fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
        if err := r.ParseForm(); err != nil {
                log.Print(err)
        }
        for k, v := range r.Form {
                fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
        }
}
  
  
  
