package main

import (
    "fmt"
//    "html/template"
    "log"
    "net/http"
    "path/filepath"
    "strings"
    "os"
    "encoding/json"
    "io/ioutil"
    "math/rand"

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

type buttstickerHandler struct {
    h           http.Handler
    contentDir  string
}

func (bh buttstickerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Vary", "Accept-Encoding")

    tickerJson, err := os.Open("/usr/share/tickerdata/tickers.json")
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    byteVal, _ := ioutil.ReadAll(tickerJson)
    var tickers []string
    json.Unmarshal(byteVal, &tickers)

    fmt.Fprintf(w, tickers[rand.Intn(4)])

    defer tickerJson.Close()
}

func main() {
    router := mux.NewRouter()
    router.Handle("/", handlers.CompressHandler(buttstickerHandler{nil, filepath.Clean(".")})).Methods("GET")
    http.Handle("/", router)

    log.Fatal(http.ListenAndServe(strings.Join([]string{":", "80"}, ""), router));
}
