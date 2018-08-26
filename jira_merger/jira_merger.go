package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "encoding/json"
    "os"

    "github.com/gorilla/mux"
)

type Configuration struct {
    servers,
    aggregated_keys [string]
}

file, _ := os.Open("config.json")
defer file.Close()
decoder := json.NewDecoder(file)
configuration := Configuration{}
err := decoder.Decode(&configuration)
if err != nil {
  fmt.Println("error:", err)
}

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/allData", GetAllData)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func GetAllData(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    data := [len(configuration.servers)]
    for i, server := configuration.servers {
      data[i] = dataFetchServerData(server, r)
    }
    merged = MergeData(data)
    json.NewEncoder(w).Encode(todos)
}

func FetchServerData(server string, initialRequest *http.Request) {
}

func MergeData(data) {
}
