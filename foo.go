package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getConfig() {
	// TODO : CHANGE THIS.
	configFilePath := "/Users/briceb/projects/go/src/github.com/server4001/logfile-alert/config.json"

	configRaw, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	type ConfigJson map[string]interface{}

	var configParsed ConfigJson

	jsonErr := json.Unmarshal(configRaw, &configParsed)
	if jsonErr != nil {
		panic(jsonErr)
	}

	watchers := configParsed["watchers"].([]interface{})
	watcher := watchers[0].(map[string]interface{})
	logFiles := watcher["log_files"].([]interface{})
	regex := watcher["regex"]
	fmt.Println(logFiles[0], logFiles[1])
	fmt.Println(regex)
}

func main() {
	getConfig()
}
