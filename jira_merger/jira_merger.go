package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "os"
    "io/ioutil"

    "github.com/gorilla/mux"
)

type Configuration struct {
    Servers []ServerConf `json:"servers"`
    Aggregated_keys []string `json:"aggregated_keys"`
    Status_mapping []int `json:"status_mapping"`
}

type ServerConf struct {
    Host string `json:"host"`
}

type AllData struct {
	  RapidViewId int `json:"rapidViewId"`
		Statistics Statistics `json:"statistics"`
		ColumnsData ColumnsData `json:"columnsData"`
		SwimlanesData SwimlanesData `json:"swimlanesData"`
		IssuesData IssuesData `json:"issuesData"`
		OrderData OrderData `json:"orderData"`
		SprintsData SprintData `json:"sprintsData"`
}

type Statistics struct {
  FieldConfigured bool `json:"fieldConfigured"`
  TypeId string `json:"typeId"`
  Id string `json:"id"`
  Name string `json:"name"`
}

type ColumnsData struct {
  RapidViewId int `json:"rapidViewId"`
  Columns []Columns `json:"columns"`
}

type Columns struct {
  Id int `json:"id"`
  Name string `json:"name"`
  StatusIds []string `json:"statusIds"`
  IsKanPlanColumn bool `json:"isKanPlanColumn"`
}

type SwimlanesData struct {
  RapidViewId int `json:"rapidViewId"`
  SwimlaneStrategy string `json:"swimlaneStrategy"`
}

type IssuesData struct {
	RapidViewId int `json:"rapidViewId"`
	ActiveFilters []ActiveFilters `json:"activeFilters"`
  Issues []Issue `json:"issues"`
}

type OrderData struct {
	RapidViewId int `json:"rapidViewId"`
	Rankable bool `json:"rankable"`
	RankCustomFieldId int `json:"rankCustomFieldId"`
}

type SprintData struct {
  RapidViewId int `json:"rapidViewId"`
  CanManageSprints bool `json:"canManageSprints"`
  EtageData EtageData `json:"etageData"`
}

type ActiveFilters struct {
	Id int `json:"id"`
}

type Issue struct {
	Id int `json:"id"`
	Key string `json:"key"`
	Hidden bool `json:"hidden"`
	TypeName string `json:"typeName"`
	TypeId string `json:"typeId"`
	Summary string `json:"summary"`
	TypeUrl string `json:"typeUrl"`
	Done bool `json:"done"`
	Assignee string `json:"assignee"`
	AssigneeName string `json:"assigneeName"`
	AvatarUrl string `json:"avatarUrl"`
	HasCustomUserAvatar bool `json:"hasCustomUserAvatar"`
	Color string `json:"color"`
	Epic string `json:"epic"`
	EpicField Epic `json:"epicField"`
	EstimateStatistic Statistic `json:"estimateStatistic"`
	TrackingStatistic Statistic `json:"trackingStatistic"`
	StatusId string `json:"statusId"`
	StatusName string `json:"statusName"`
	StatusUrl string `json:"statusUrl"`
	Status Status `json:"status"`
	FixVersions []FixVersion `json:"fixVersions"`
	ProjectId int `json:"projectId"`
	LinkedPagesCount int `json:"linkedPagesCount"`
	ExtraFields []ExtraField `json:"extraFields"`
}

type ExtraField struct {
	Id string `json:"id"`
	Label string `json:"label"`
	Editable bool `json:"editable"`
	Renderer string `json:"renderer"`
	Html string `json:"html"`
}

type FixVersion struct {
}

type Status struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	IconUrl string `json:"iconUrl"`
	StatusCategory StatusCategory `json:"statusCategory"`
}

type StatusCategory struct {
	Id string `json:"id"`
	Key string `json:"key"`
}

type Epic struct {
	Id string `json:"id"`
	Label string `json:"label"`
}


type Statistic struct {

}

type Sprints struct {
	Id string `json:"id"`
	Name string `json:"name"`
  Sequence int `json:"sequence"`
  State string `json:"state"`
  LinkedPagesCount int `json:"linkedPagesCount"`
  StartDate string `json:"startDate"`
  EndDate string `json:"endDate"`
  CompleteDate string `json:"completeDate"`
  CanUpdateSprint bool `json:"canUpdateSprint"`
  DaysRemaining int `json:"daysRemaining"`
}

type EtageData struct {
  RapidViewId int `json:"rapidViewId"`
  IssueCount int `json:"issueCount"`
  LastUpdated int `json:"lastUpdated"`
  QuickFilters string `json:"quickFilters"`
  Sprints int `json:"sprints"`
  Etag string `json:"etag"`
}

var configuration Configuration
var httpClient = &http.Client{}

func main() {
		configuration = loadConfig("/go/src/jira_merger/config.json")
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/allData", GetAllData)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func loadConfig(file string) Configuration {
  var config Configuration
  configFile, err := os.Open(file)
  defer configFile.Close()
  if err != nil {
      fmt.Println(err.Error())
  }
  jsonParser := json.NewDecoder(configFile)
  jsonParser.Decode(&config)
  return config
}

func DecodeData(jsonData []byte) AllData {
  allData := AllData{}
  jsonErr := json.Unmarshal(jsonData, &allData)
  if jsonErr != nil {
    log.Fatal(jsonErr)
  }
  return allData
}

func GetAllData(w http.ResponseWriter, r *http.Request) {
    var mergedData AllData
    data := make([]AllData, len(configuration.Servers))
    for i, server := range configuration.Servers {
      data[i] = DecodeData(FetchServerData(server, r))
    }
    mergedData = MergeData(data...)
    json.NewEncoder(w).Encode(mergedData)
}

func FetchServerData(server ServerConf, initialRequest *http.Request) []byte {
	initialRequest.URL.Query().Set("hostname", server.Host)
	req, err := http.NewRequest("GET", server.Host + initialRequest.URL.String(), nil)
	if err != nil {
      fmt.Printf("Error fetching server data : %s\n", err)
  }
	resp, err := httpClient.Do(req)
	if err != nil {
      fmt.Printf("Error performing server request : %s\n", err)
  }
  defer resp.Body.Close()

  body, readErr := ioutil.ReadAll(resp.Body)
  if readErr != nil {
		log.Fatal("Error reading server response : %s\n", readErr)
	}

	return body
}

func MergeData(dataArray ...AllData) AllData {
  for _, data := range dataArray[1:len(dataArray)] {
    dataArray[0] = merge(dataArray[0], data)
  }
	return dataArray[0]
}

func merge(base AllData, add AllData) AllData {
  base.IssuesData.Issues = append(base.IssuesData.Issues, add.IssuesData.Issues...)
  return base
}

func MapStatus(initial int) int {
  return initial
}
