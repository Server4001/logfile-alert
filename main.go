package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getConfig() []Watcher {
	// TODO : CHANGE THIS.
	configFilePath := "/Users/briceb/projects/go/src/github.com/server4001/logfile-alert/config.json"

	configRaw, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	var configParsed map[string]interface{}

	jsonErr := json.Unmarshal(configRaw, &configParsed)
	if jsonErr != nil {
		panic(jsonErr)
	}

	watchersParsed := configParsed["watchers"].([]interface{})
	watchers := make([]Watcher, 0)

	for idx := range watchersParsed {
		watcher := watchersParsed[idx].(map[string]interface{})
		logFiles := watcher["log_files"].([]interface{})
		regex := watcher["regex"].(string)

		watchers = append(watchers, Watcher{
			logFiles,
			regex,
		})
	}

	return watchers
}

func main() {
	watchers := getConfig()

	for _, watcher := range watchers {
		fmt.Println(watcher.regex)
		for _, logFile := range watcher.logFiles {
			fmt.Println(logFile)
		}
	}

	// sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)

}
