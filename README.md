# xwallpaper

## Usage

```
Usage: wallpaper <action> [path to image]
Actions:
  set <path to image> - save the wallpaper from the given path (doesn't apply it)
  apply - apply the wallpaper from the saved path
  get - print the saved wallpaper json
```

Also optional flag when using `set`, `-o` which sets the path to save the wallpapers json.

## Building

```sh
$ go build -o bin/wallpaper .
```

## Plans

- Add daemon support
- Add `append` action
- Allow randomly picking a wallpaper
- Add timed wallpaper changes
