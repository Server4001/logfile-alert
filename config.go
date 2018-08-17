package main

import "sync"

// Config - Represents a configuration file for our application.
type Config struct {
	sync.RWMutex
	watchers []Watcher
}

func (c *Config) setWatchers(watchers []Watcher) {
	c.Lock()
	defer c.Unlock()
	c.watchers = watchers
}

func (c *Config) getWatchers() []Watcher {
	c.RLock()
	defer c.RUnlock()
	return c.watchers
}
