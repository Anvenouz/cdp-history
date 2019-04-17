package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type (
	//CDP json info
	CDP struct {
		ID      int     `json:"ID"`
		DaiDebt float64 `json:"DaiDebt"`
		EthCol  float64 `json:"EthCol"`
		Price   float64 `json:"Price"`
		Ratio   float64 `json:"Ratio"`
		DaiNet  float64 `json:"DaiNet"`
		EthNet  float64 `json:"EthNet"`
	}
	//History export
	History struct {
		Date string `json:"date"`
		CDP  *CDP   `json:"cdp"`
	}

	//HTMLPage page html returned
	HTMLPage struct {
		PageTitle string
		History   []History
		LastDate  string
	}
)

var (
	//CDPData for page history
	CDPData []History
)

func main() {
	jsonFile := flag.String("file", "", "History file from your cdp.")
	port := flag.Int("port", 80, "Port to use for the service")
	flag.Parse()

	if "" == *jsonFile {
		log.Fatal("You need to add the history file ! see --help.")
	}

	data, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	CDPData = make([]History, 1000)
	err = json.Unmarshal(data, &CDPData)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", index)           // set router
	err = http.ListenAndServe(":"+ strconv.Itoa(*port), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	lastDate := CDPData[len(CDPData)-1].Date
	data := HTMLPage{
		PageTitle: "CDP - HISTORY",
		History:   CDPData,
		LastDate:  lastDate,
	}
	err := tmpl.Execute(w, data)
	if nil != err {
		log.Println(err)
	}
}
