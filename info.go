package main

import "time"

type WallpaperInfo struct {
	// ChangeTime specifies how often the wallpaper should change.
	// This feature only works if the wallpaper app is daemonized.
	ChangeTime ChangeTime `json:"changeTime"`
	// Shuffle randomizes the wallpaper chosen when changing the wallpaper.
	Shuffle    bool     `json:"shuffle,omitempty"`
	Wallpapers []string `json:"wallpapers"`
}

type ChangeTime struct {
	Min time.Duration `json:"min"`
	Max time.Duration `json:"max"`
	// If random is true then the wallpaper will change at a random time between
	// the min and max.
	//
	// If false, then the wallpaper will change at the min time.
	Random bool `json:"random"`
}
