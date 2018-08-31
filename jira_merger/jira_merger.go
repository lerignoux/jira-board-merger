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

type AllData struct {
	  rapidViewId int `json:"rapidViewId"`
		statistics
		columnsData ColumnsData
		orderData OrderData
		sprintsData SprintData
}

type ColumnData struct {

}

type IssuesData struct {
	rapidViewId string
	activeFilters []Filters
}

type OrderData struct {
	rapidViewId int
	rankable bool
	rankCustomFieldId int
}

type SprintData struct {}

type Filters struct {
	id: int
}

type Issue struct {
	id int
	key string
	hidden bool
	typeName string
	typeId string
	summary string
	typeUrl string
	done bool
	assignee string
	assigneeName string
	avatarUrl string
	hasCustomUserAvatar bool
	color bool
	epic string
	epicField Epic
	estimateStatistic Statistic
	trackingStatistic Statistic
	statusId string
	statusName string
	statusUrl string
	status Status
	fixVersions []FixVersion
	projectId int
	linkedPagesCount string
	extraFields []ExtraField
}

type ExtraField struct {
	id: string
	labe: string
	editable: string
	renderer: string
	html: string
}

type FixVersion struct {

}

type Status struct {
	id string
	name string
	description string
	iconUrl string
	statusCategory StatusCategory
}

type StatusCategory struct {
	id string
	key string
}


type Epic struct {
	id string
	label string
}

type Statistic struct {

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

func mergeIssues(base IssueData, add IssueData) IssueData {}
