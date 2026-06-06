# kaiten-wall (WIP)

Inspired from the kaiten sushi (conveyer belt sushi); selects a random wallpaper from Wallhaven and applies it as the wallpaper with \*matugen scheme generation

Written in Golang.

> Note: Terms tagged as '\*' are planned to be implemented

## Dependencies

- Wallpaper daemon (awww, \*swaybg, \*hyprpaper)
- Wayland compositor (niri, hyprland, mangowc)

## Build

Ensure that you have the following:

```bash
# wallhaven API wrapper
go get github.com/Yustinia/gopaper

# toml parsing
go get github.com/BurntSushi/toml
```

Otherwise, simply do `go mod tidy` to automatically handle build dependencies

### Just

Ensure that you have `just` installed through your package manager

```bash
just build
```

### Manual

```bash
go build -o kaiten ./cmd
```

## Install

> TODO

## Features

- Fetch a random wallpaper and apply it
- Filter wallpapers through the config toml file
- Change the default output path for wallpapers
- Configure how wallpapers transition
- Change which wallpaper daemon to use

## To Do

- Expand wallpaper daemon selection
  - swaybg
  - hyprpaper
- Expand wallpaper sources
- Implement matugen color generation
