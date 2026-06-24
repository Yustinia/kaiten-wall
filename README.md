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

Make sure dependencies are available:

```bash
go mod tidy
```

Or manually:

```bash
go get github.com/Yustinia/gopaper  # Wallhaven API wrapper
go get github.com/BurntSushi/toml   # TOML parsing
```

### With Just

```bash
just build
```

### Manual

```bash
go build -o kaiten ./cmd
```

---

## Installing

### With Just

Builds and copies the binary to `~/.local/bin`:

```bash
just install
```

### Manual

Build first, then:

```bash
cp kaiten ~/.local/bin/
```

---

## Usage

Run `kaiten` once after installing. It will generate a base configuration at:

```bash
~/.config/kaiten-wall/config.toml
```

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
