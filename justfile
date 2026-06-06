entryPoint := "./cmd"
outName := "kaiten"
localBin := "~/.local/bin/"

default:
    just --list

run:
    @go run {{ entryPoint }}

build:
    @go build -o {{ outName }} {{ entryPoint }}

install: build
    @mv -v {{ outName }} {{ localBin }}
    @echo installed {{ outName }} to {{ localBin }}
