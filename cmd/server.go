package main

import (
	"net/http"
	"log"
	"ferp-api/pkg/api"
	"ferp-api/pkg/model"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	port, databaseName, databaseURL, _ := readFileConfig()
	model.InitDatabase(databaseName, databaseURL)
	// model.CreateDatabase()
	api := api.NewAPI(api.NewRouter())
	fmt.Println("server start... => http://localhost" + port)
	log.Fatal(http.ListenAndServe(":" + port, api.MakeHandler()))
}

func readFileConfig() (portOut, databaseNameOut, databaseURLOut, setUpDatabaseOut string) {
	config, err := ioutil.ReadFile("./config.json")
	checkErr(err)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(config, &objmap)
	checkErr(err)
	var port, databaseName, databaseURL, setUpDatabase string
	err = json.Unmarshal(*objmap["port"], &port)
	checkErr(err)
	err = json.Unmarshal(*objmap["databaseName"], &databaseName)
	checkErr(err)
	err = json.Unmarshal(*objmap["databaseURL"], &databaseURL)
	checkErr(err)
	err = json.Unmarshal(*objmap["setUpDatabase"], &setUpDatabase)
	checkErr(err)
	return port, databaseName, databaseURL, setUpDatabase
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}