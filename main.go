package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
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

	configReload := make(chan os.Signal, 1)
	shutdown := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(configReload, syscall.SIGHUP)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go reloadHandler(configReload)
	go shutdownHandler(shutdown, configReload, done)

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

func reloadHandler(configReload chan os.Signal) {
	for sig := range configReload {
		fmt.Println("Reloading config due to:", sig)
	}
}

func shutdownHandler(
	shutdown chan os.Signal,
	configReload chan os.Signal,
	done chan bool) {

	sig := <-shutdown
	fmt.Println("Quitting due to signal:", sig)

	// Shut down config reloads.
	signal.Stop(configReload)
	close(configReload)
	done <- true
}
