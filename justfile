entryPoint := "./cmd"
outName := "kaiten"

default:
    just --list

run:
    @go run {{ entryPoint }}

build:
    @go build -o {{ outName }} {{ entryPoint }}
