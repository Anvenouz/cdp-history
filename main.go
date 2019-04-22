package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

	//EqualizerCDP Structure like the json from cdp-equalizer
	EqualizerCDP struct {
		Current  *CDP  `json:"current"`
		KeyPrices []CDP `json:"KeyPrices"`
		//Up       [50]CDP `json:"Up"`
		//Down     [50]CDP `json:"Down"`
	}

	//HTMLPage page html returned
	HTMLPage struct {
		PageTitle string
		History   []History
		LastDate  string
		Equalizer EqualizerCDP
	}
)

var (
	//CDPData for page history
	CDPData []History
	//EqualizerCDPData data json from the cdp equalizer
	EqualizerCDPData EqualizerCDP
	jsonFile         *string
	jsonCDP          *string
)

func main() {
	jsonFile = flag.String("file", "", "History file from your cdp.")
	host := flag.String("host", "localhost:8888", "Hostname or IP and port to use for the service")
	jsonCDP = flag.String("cdp", "https://anvenouz.be/", "Json CDP from your cdp-equalizer service (see : https://github.com/efemero/cdp-equalizer)")

	flag.Parse()

	if "" == *jsonFile {
		log.Fatal("You need to add the history file ! see --help.")
	}

	if "" != *jsonCDP {

		var myClient = &http.Client{Timeout: 10 * time.Second}

		r, err := myClient.Get(*jsonCDP)
		if err != nil {
			fmt.Println(err)
		}
		defer r.Body.Close()

		err = json.NewDecoder(r.Body).Decode(&EqualizerCDPData)

		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(EqualizerCDPData.Current.EthNet)
		//return
	}

	data, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &CDPData)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", index)           // set router
	err = http.ListenAndServe(*host, nil) // set listen host
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	readJSONFile(jsonFile)
	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	lastDate := CDPData[len(CDPData)-1].Date
	data := HTMLPage{
		PageTitle: "CDP - HISTORY",
		History:   CDPData,
		LastDate:  lastDate,
		Equalizer: EqualizerCDPData,
	}
	err := tmpl.Execute(w, data)
	if nil != err {
		log.Println(err)
	}
}

func readJSONFile(file *string) {
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &CDPData)
	if err != nil {
		log.Fatal(err)
	}
}
