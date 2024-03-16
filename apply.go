package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
)

func applyWallpaper(savePath string) error {
	f, err := os.Open(savePath)
	if err != nil {
		return err
	}
	defer f.Close()

	var wallpaperInfo WallpaperInfo

	if err := json.NewDecoder(f).Decode(&wallpaperInfo); err != nil {
		return err
	}

	if len(wallpaperInfo.Wallpapers) == 0 {
		return errors.New("no wallpapers foun")
	}

	chosen := wallpaperInfo.Wallpapers[0]
	if wallpaperInfo.Shuffle {
		chosen = wallpaperInfo.Wallpapers[rand.IntN(len(wallpaperInfo.Wallpapers)-1)]
	}

	// nitrogen --set-zoom-fill $wallpaper --head=$i
	for i := 0; i < 5; i++ {
		cmd := exec.CommandContext(context.Background(), "nitrogen", "--set-zoom-fill", chosen, fmt.Sprintf("--head=%d", i))
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to run [%s]: %w", cmd.String(), err)
		}
	}

	return nil
}
