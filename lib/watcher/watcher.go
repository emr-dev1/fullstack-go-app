package watcher

// State represents the last-known state of a URL.
type State struct {
	path   string
	status string
}
