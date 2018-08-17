package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

var (
	config = &Config{}
)

func getConfig() (watchers []Watcher, err error) {
	// TODO : CHANGE THIS.
	configFilePath := "/Users/briceb/projects/go/src/github.com/server4001/logfile-alert/config.json"

	configRaw, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return
	}

	var configParsed map[string]interface{}

	err = json.Unmarshal(configRaw, &configParsed)
	if err != nil {
		return
	}

	watchersParsed, casted := configParsed["watchers"].([]interface{})
	if !casted {
		err = errors.New("invalid watchers node in config file")
		return
	}

	for idx := range watchersParsed {
		watcher := watchersParsed[idx].(map[string]interface{})

		logFiles, casted := watcher["log_files"].([]interface{})
		if !casted {
			err = errors.New(fmt.Sprint("invalid log_files node in config file for watcher index ", idx))
			return
		}

		regex, casted := watcher["regex"].(string)
		if !casted {
			err = errors.New(fmt.Sprint("invalid regex node in config file for watcher index ", idx))
			return
		}

		watchers = append(watchers, Watcher{
			logFiles,
			regex,
		})
	}

	return
}

func main() {
	watchers, err := getConfig()

	if err != nil {
		// TODO : Change to log.Println("open config: ", err); os.Exit(1)
		panic(err)
	}

	config.setWatchers(watchers)

	// TODO : REMOVE.
	for _, watcher := range config.getWatchers() {
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
	for range configReload {
		watchers, err := getConfig()

		if err != nil {
			// TODO : Change to STDERR.
			fmt.Println("Failed to reload config:", err)
		} else {
			config.setWatchers(watchers)
			fmt.Println("Reloaded config.")
		}
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
