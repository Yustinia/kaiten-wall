# kaiten-wall

Ever get tired of staring at the same wallpaper? kaiten-wall keeps things fresh — like a conveyor belt of wallpapers pulled straight from Wallhaven, applied automatically with a generated color scheme.

Written in Go.

> Items marked with `*` are planned but not yet implemented.

---

## Dependencies

**Wallpaper Daemons**

- [awww](https://codeberg.org/LGFae/awww)
- \*swaybg
- \*hyprpaper

**Wayland Compositors**

- niri
- hyprland
- mangowc

**Color Scheme Generators**

- [matugen](https://github.com/InioX/matugen) (default)
- [wallust](https://codeberg.org/explosion-mental/wallust)

---

## Building

### Manual

```bash
go build -o kaiten ./cmd
```

### Make

To build with `make`, just run `make`.

---

## Installing

### Manual

Build first, then:

```bash
cp kaiten ~/.local/bin/
```

### Make

```bash
make install
```

This automatically builds and copies `kaiten` to `~/.local/bin/kaiten`

---

## Usage

Run `kaiten` once after installing. It will generate a base configuration at:

```bash
~/.config/kaiten-wall/config.toml
```

Otherwise, run `kaiten --help` to show usable flags.

From there, run it again to fetch and apply a random wallpaper. The config is self-documenting — each field includes a comment explaining what it does and what values are accepted.

---

## Features

- Fetch and apply a random wallpaper from Wallhaven
- Filter wallpapers via config
- Configure wallpaper transition behavior
- Set a custom wallpaper output path
- Choose your wallpaper daemon
- Choose between matugen and wallust for color generation

---

## To Do

- [ ] swaybg support
- [ ] hyprpaper support
- [ ] Additional wallpaper sources
