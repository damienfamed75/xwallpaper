package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	defaultWallpaperPath = os.Getenv("HOME") + "/.config/wallpaper"
)

const (
	_pathPermission = 0755
	_fileName       = "wallpapers.json"
)

func checkWallpaperDir(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("creating %s\n", path)
		if err := os.MkdirAll(path, _pathPermission); err != nil {
			return err
		}
	}

	return nil
}

// validatePath returns an error if the path is malformed
// func validatePath(p string) error {
// 	_, err := filepath.Abs(p)
// 	return err
// }

func setWallpaper(wallpaperPath, savePath string) error {
	// Ensure that the given path is an absolute path.
	if err := validatePath(wallpaperPath); err != nil {
		return err
	}

	// Check that the save path exists, and create it if it doesn't.
	if err := checkWallpaperDir(savePath); err != nil {
		return err
	}

	fmt.Printf("wallpaperPath: %s\nsavePath: %s\n", wallpaperPath, savePath)

	f, err := os.OpenFile(filepath.Join(savePath, _fileName), os.O_CREATE|os.O_RDWR, _pathPermission)
	if err != nil {
		return err
	}
	defer f.Close()

	// todo: check if append or set
	wallpaperInfo := WallpaperInfo{
		Wallpapers: []string{wallpaperPath},
	}

	bb, err := json.Marshal(wallpaperInfo)
	if err != nil {
		return err
	}
	// Truncate the file to 0 bytes before writing.
	if err := f.Truncate(0); err != nil {
		return err
	}

	if _, err := f.Write(bb); err != nil {
		return err
	}

	fmt.Printf("set wallpaper to %s\n", wallpaperPath)

	return nil
}

func overwriteWallpapers() error {
	return nil

}
