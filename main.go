package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	savePath := flag.String("o", defaultWallpaperPath, "path to save wallpaper")
	flag.Parse()

	if len(os.Args) < 2 || os.Args[1] == "help" {
		fmt.Println("Usage: wallpaper <action> [path to image]")
		fmt.Println("Actions:")
		fmt.Println("  set <path to image> - save the wallpaper from the given path (doesn't apply it)")
		fmt.Println("  apply - apply the wallpaper from the saved path")
		fmt.Println("  get - print the saved wallpaper json")
		os.Exit(1)
	}

	actionStr := os.Args[1]
	action, err := ActionFromString(actionStr)
	if err != nil {
		fmt.Printf("invalid action %s\n", actionStr)
		os.Exit(1)
	}

	// validate the given paths.
	if err := validatePath(*savePath); err != nil {
		panic(err)
	}

	// Based on the provided action, perform the appropriate action.
	switch action {
	case Append:
		fmt.Println("append not implemented yet, setting instead...")
		fallthrough
	case Set:
		if len(os.Args) < 3 {
			fmt.Print("wallpaper set <path to image>\n")
			flag.Usage()
		}
		wallpaperPath := os.Args[2]
		if err := setWallpaper(wallpaperPath, *savePath); err != nil {
			panic(err)
		}
	case Apply:
		if err := applyWallpaper(filepath.Join(*savePath, _fileName)); err != nil {
			panic(err)
		}
	case Get:
		if err := getWallpaper(filepath.Join(*savePath, _fileName)); err != nil {
			panic(err)
		}
	}
}

// validatePath returns an error if the path is malformed
func validatePath(p string) error {
	_, err := filepath.Abs(p)
	return err
}
