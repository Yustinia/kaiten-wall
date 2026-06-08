# kaiten-wall (WIP)

Inspired by kaiten sushi (conveyor belt sushi) — fetches a random wallpaper from Wallhaven and applies it, generating a color scheme with either [matugen](https://github.com/InioX/matugen) or [wallust](https://codeberg.org/explosion-mental/wallust).

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

Build first (see above), then:

```bash
cp kaiten ~/.local/bin/
```

### Makepkg

Ensure you have `git` and `base-devel` installed:

```bash
pacman -S --needed git base-devel
git clone https://aur.archlinux.org/kaiten-wall-bin.git
cd kaiten-wall-bin
makepkg -si
```

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
