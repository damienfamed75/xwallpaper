package main

import (
	"io"
	"os"
)

func getWallpaper(savePath string) error {
	f, err := os.Open(savePath)
	if err != nil {
		return err
	}
	defer f.Close()

	bb, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	os.Stdout.Write(bb)

	return nil
}
