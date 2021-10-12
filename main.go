package main

import(
    "net/http"
    "html/template"
    "log"
    "fmt"
)

var cnt int

func main(){
    cnt=0
    http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs/"))))
    http.HandleFunc("/",index)
    http.HandleFunc("/count",Count)
    log.Fatal(http.ListenAndServe(":80",nil))
}
func index(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("./docs/index.html")
    if err != nil {
        panic(err.Error())
    }
    cnt++
    if err := t.Execute(w, nil); err != nil {
        panic(err.Error())
    }
}
func Count(w http.ResponseWriter, r *http.Request){
    if r.Method!=http.MethodGet{
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("Getのみです"))
        return
    }
    fmt.Fprint(w,cnt)
}
