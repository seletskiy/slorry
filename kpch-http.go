package main

import "fmt"
//import "log"
import "os"
//import "image"
import "image/png"
import "net/http"

import "slorry/kpch/crack"

func handler(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query()["q"][0]

	w.Header().Set("Content-Type", "text/html")
    
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(w, "%v\n", err)
        return
    }

    img, err := png.Decode(resp.Body)
    if err != nil {
        fmt.Fprintf(w, "%v\n", err)
        return
    }

    for _, s := range crack.Crack(img) {
        fmt.Fprint(w, s.Sym.Char)
    }

    fmt.Fprintf(w, `<img src="%s"/>`, url)
}

func main() {
    listenTo := os.Args[1]
    http.HandleFunc("/", handler)
    http.ListenAndServe(listenTo, nil)
}
