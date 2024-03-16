package main

// Daemon will keep track of the wallpapers and change them at the appropriate time.
// The daemon also keeps track of file changes to the wallpaper file.
type Daemon struct {
}

// NewDaemon returns a new daemon.
func NewDaemon() *Daemon {
	return &Daemon{}
}

// Start starts the daemon.
