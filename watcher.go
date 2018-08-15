package main

// Watcher - Represents a watcher node in the config.json file.
type Watcher struct {
	logFiles []interface{}
	regex    string
}
