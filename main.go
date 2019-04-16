package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
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
	}
)

var (
	CDPData []History
	//TemplateFile *string
)

func main() {
	jsonFile := flag.String("file", "", "History file from your cdp.")
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

	//fmt.Println(CDPData)

	//for _, v := range CDPData {
	//	fmt.Println(v.CDP)
	//}

	//for _, v := range vJSON {
	//	fmt.Printf("%+v \n", v.Date)
	//	fmt.Printf("%+v \n", v.CDP)
	//}

	//tmpl, err := template.ParseFiles("templates/template.html")

	http.HandleFunc("/", index)           // set router
	err = http.ListenAndServe(":80", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	//fmt.Println("The next one ought to fail.")
	//tErr := template.New("check parse error with Must")
	//template.Must(tErr.Parse(" some static text {{ .Name }"))

}

func index(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	//datas := make([]CDP, len(*CDPData))
	//for _, v := range CDPData {
	//	r := v.CDP
	//	//fmt.Println("test", r)
	//	//datas = append(datas, *r)
	//}
	//fmt.Println(CDPData)
	//fmt.Println(datas)
	data := HTMLPage{
		PageTitle: "Mon super site internet",
		History:   CDPData,
	}
	tmpl.Execute(w, data)

	//r.ParseForm()  // parse arguments, you have to call this by yourself
	//fmt.Println(r.Form)  // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//    fmt.Println("key:", k)
	//    fmt.Println("val:", strings.Join(v, ""))
	//}
	//fmt.Fprintf(w, data) // send data to client side
}
