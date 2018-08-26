package main

import (
    "bytes"
    "fmt"
    "html"
    "log"
    "net/http"
    "encoding/json"
    "os"

    "github.com/gorilla/mux"
)

type Configuration struct {
    servers         []ServerConf `json:"array"`
    aggregated_keys []string
}

type ServerConf struct {
    host     string  `json:"host"`
}

var configuration Configuration
var httpClient = &http.Client{}

func main() {
		loadConfig("config.json")
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/allData", GetAllData)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func loadConfig(configName string) {
	file, _ := os.Open(configName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
	  fmt.Println("error:", err)
	}
}

func GetAllData(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    data := make([]string, len(configuration.servers))
    for i, server := range configuration.servers {
      data[i] = FetchServerData(server, r)
    }
    merged := MergeData(data...)
    json.NewEncoder(w).Encode(merged)
}

func FetchServerData(server ServerConf, initialRequest *http.Request) string {
	initialRequest.URL.Query().Set("hostname", server.host)
	req, err := http.NewRequest("GET", initialRequest.URL.String(), nil)
	if err != nil {
      fmt.Printf("Error : %s", err)
  }
	resp, err := httpClient.Do(req)
	if err != nil {
      fmt.Printf("Error : %s", err)
  }
  fmt.Println("resp")

  var b bytes.Buffer
  _, err = b.ReadFrom(resp.Body)
  if err != nil {
      log.Fatal("Error : %s", err)
  }
	return b.String()
}

func MergeData(data ...string) string {
	return "ok"
}
